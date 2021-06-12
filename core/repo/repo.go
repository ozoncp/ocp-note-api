package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Repo interface {
	AddNotes(ctx context.Context, notes []note.Note) error
	DescribeNote(ctx context.Context, id uint64) (*note.Note, error)
	ListNotes(ctx context.Context, count, offset uint64) ([]note.Note, error)
	RemoveNote(ctx context.Context, id uint64) error
}

const (
	tableName = "notes"
)

type repo struct {
	db sqlx.DB
}

func New(db sqlx.DB) Repo {
	return &repo{db: db}
}

func (r *repo) AddNotes(ctx context.Context, notes []note.Note) error {
	return nil
}

func (r *repo) DescribeNote(ctx context.Context, id uint64) (*note.Note, error) {
	return nil, nil
}

func (r *repo) ListNotes(ctx context.Context, count, offset uint64) ([]note.Note, error) {
	return nil, nil
}

func (r *repo) RemoveNote(ctx context.Context, id uint64) error {
	return nil
}
