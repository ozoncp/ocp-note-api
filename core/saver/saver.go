package saver

import (
	"fmt"
	"log"

	"github.com/ozoncp/ocp-note-api/core/alarmer"
	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Saver interface {
	Save(note note.Note)
	Init()
	Close()
}

type saver struct {
	capacity    uint
	flusher     flusher.Flusher
	alarmer     alarmer.Alarmer
	notesChan   chan note.Note
	notes       []note.Note
	lossAllData bool
	end         chan struct{}
}

func New(capacity uint, flusher flusher.Flusher, alarmer alarmer.Alarmer, lossAllData bool) Saver {
	return &saver{
		capacity:    capacity,
		flusher:     flusher,
		alarmer:     alarmer,
		notesChan:   make(chan note.Note),
		end:         make(chan struct{}),
		lossAllData: lossAllData,
	}
}

func (s *saver) Init() {
	go func() {
		for {
			select {
			case noteTmp := <-s.notesChan:
				s.saveData(noteTmp)
				fmt.Printf("size: %v\n", len(s.notes))

			case _, ok := <-s.alarmer.Alarm():
				if ok {
					fmt.Println("flush")
					s.flushData()
				} else {
					fmt.Println("non flush")
				}
			case <-s.end:
				fmt.Println("finish")
				return
			}
		}
	}()
}

func (s *saver) Save(note note.Note) {
	s.notesChan <- note
}

func (s *saver) saveData(note note.Note) {
	if len(s.notes) >= int(s.capacity) {
		if s.lossAllData {
			fmt.Println("0.1")
			s.notes = s.notes[:0]
		} else {
			fmt.Println("0.2")
			s.notes = s.notes[1:]
		}
	}

	s.notes = append(s.notes, note)
}

func (s *saver) flushData() {
	response := s.flusher.Flush(s.notes)

	if response != nil {
		log.Fatal("failed to flush")
	}

	s.notes = s.notes[:copy(s.notes, response)]
}

func (s *saver) Close() {
	s.end <- struct{}{}
	s.flushData()
	s.alarmer.Close()
}
