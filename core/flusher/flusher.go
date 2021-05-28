package flusher

import (
	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/internal/utils"
)

type Flusher interface {
	Flush(notes []note.Note) []note.Note
}

type flusher struct {
	chunkSize int
	storage   repo.Repo
}

func New(chunkSize int, storage repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		storage:   storage,
	}
}

func (f *flusher) Flush(notes []note.Note) []note.Note {

	chunks := utils.SplitNoteSlice(notes, f.chunkSize)
	var successPos = 0

	for _, val := range chunks {
		if err := f.storage.AddNotes(val); err != nil {
			return notes[successPos:]
		}

		successPos += len(val)
	}

	return nil
}
