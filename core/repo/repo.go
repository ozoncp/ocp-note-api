package repo

import "github.com/ozoncp/ocp-note-api/core/note"

type Repo interface {
	AddNotes(notes []note.Note) error
}
