package saver

import (
	"fmt"

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
	notes       []note.Note
	lossAllData bool
}

func New(alarmer alarmer.Alarmer) Saver {
	return &saver{
		alarmer: alarmer,
	}
}

func (s *saver) Init() {
	go func() {
		for {
			select {
			case _, ok := <-s.alarmer.Alarm():
				if ok {
					fmt.Println("check")
				} else {
					fmt.Println("non check")
				}
			}
		}
	}()
}

func (s *saver) Save(note note.Note) {

	if len(s.notes) >= int(s.capacity) {
		if s.lossAllData {

		} else {

		}
	}

	s.notes = append(s.notes, note)

	// if channal {
	// 	s.flusher.Flush(s.notes)
	// }
}

func (s *saver) Close() {
	s.flusher.Flush(s.notes)
}
