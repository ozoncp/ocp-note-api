package saver

import (
	"context"
	"errors"
	"log"

	"github.com/ozoncp/ocp-note-api/core/alarmer"
	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Saver interface {
	Save(note note.Note)
	Init() error
	Close()
}

type saver struct {
	ctx         context.Context
	capacity    uint
	flusher     flusher.Flusher
	alarmer     alarmer.Alarmer
	notesChan   chan note.Note
	notes       []note.Note
	lossAllData bool
	end         chan struct{}
	initPassed  bool
}

func New(ctx context.Context, capacity uint, flusher flusher.Flusher, alarmer alarmer.Alarmer, lossAllData bool) Saver {

	if capacity <= 0 || flusher == nil || alarmer == nil {
		return nil
	}

	return &saver{
		ctx:         ctx,
		capacity:    capacity,
		flusher:     flusher,
		alarmer:     alarmer,
		notesChan:   make(chan note.Note),
		notes:       []note.Note{},
		lossAllData: lossAllData,
		end:         make(chan struct{}),
		initPassed:  false,
	}
}

func (s *saver) Init() error {

	if s.initPassed {
		return errors.New("the saver has already been initialized")
	}

	err := s.alarmer.Init()

	if err != nil {
		panic("alarm is faulty")
	}

	go func() {
		for {
			select {
			case noteTmp := <-s.notesChan:
				s.saveData(noteTmp)
			case _, ok := <-s.alarmer.Alarm():
				if ok {
					s.flushData()
				} else {
					log.Fatalln("signal reception error on the flush")
				}
			case <-s.end:
				return
			}
		}
	}()

	s.initPassed = true
	return nil
}

func (s *saver) Save(note note.Note) {

	if !s.initPassed {
		panic("the saver has not been initialized")
	}

	s.notesChan <- note
}

func (s *saver) saveData(note note.Note) {
	if len(s.notes) >= int(s.capacity) {
		if s.lossAllData {
			s.notes = s.notes[:0]
		} else {
			s.notes = s.notes[1:]
		}
	}

	s.notes = append(s.notes, note)
}

func (s *saver) flushData() {
	response := s.flusher.Flush(s.ctx, s.notes)

	if response != nil {
		log.Fatalln("failed to flush")
	}

	s.notes = s.notes[:copy(s.notes, response)]
}

func (s *saver) Close() {
	s.end <- struct{}{}
	s.flushData()
	s.alarmer.Close()
}
