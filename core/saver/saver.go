package saver

import (
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Saver interface {
	Save(entity note.Note)
	// Init()
	Close()
}
