package repo

import "github.com/ozoncp/ocp-note-api/core/note"

type Repo interface {
	AddNotes(notes []note.Note) error
	DescribeNote(id uint64) (*note.Note, error)
	ListNotes(count, offset uint64) ([]note.Note, error)
	RemoveNote(id uint64) error
}
