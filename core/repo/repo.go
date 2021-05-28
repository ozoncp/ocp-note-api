package repo

import "github.com/ozoncp/ocp-note-api/core/note"

type Repo interface {
	AddTasks(notes []note.Note) error
}
