package saver

import (
	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Saver interface {
	Save(note note.Note)
	// Init()
	Close()
}

type saver struct {
	capacity uint
	flusher  flusher.Flusher
}

func (s *saver) Save(note note.Note) {

}

func (s *saver) Close() {

}

func New(capacity uint, flusher flusher.Flusher) Saver {
	return &saver{
		capacity: capacity,
		flusher:  flusher,
	}
}
