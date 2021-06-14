package flusher

import (
	"context"

	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, notes []note.Note) []note.Note
}

type flusher struct {
	chunkSize uint32
	storage   repo.Repo
}

func New(storage repo.Repo, chunkSize uint32) Flusher {
	return &flusher{
		storage:   storage,
		chunkSize: chunkSize,
	}
}

func (f *flusher) Flush(ctx context.Context, notes []note.Note) []note.Note {

	chunks := utils.SplitNoteSlice(notes, f.chunkSize)
	var successPos = 0

	for _, val := range chunks {
		if _, err := f.storage.MultiAddNotes(ctx, val); err != nil {
			return notes[successPos:]
		}

		successPos += len(val)
	}

	return nil
}
