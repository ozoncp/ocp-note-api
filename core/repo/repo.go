package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Repo interface {
	AddNote(ctx context.Context, note note.Note) (uint64, error)
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

func (r *repo) AddNote(ctx context.Context, note note.Note) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("user_id", "classroom_id", "document_id").
		Values(note.UserId, note.ClassroomId, note.DocumentId).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	query.QueryRowContext(ctx).Scan(&note.Id)

	return note.Id, nil
}

func (r *repo) AddNotes(ctx context.Context, notes []note.Note) error {
	query := sq.Insert(tableName).
		Columns("user_id", "classroom_id", "document_id").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, note := range notes {
		query = query.Values(note.UserId, note.ClassroomId, note.DocumentId)
	}

	_, err := query.ExecContext(ctx)

	return err
}

func (r *repo) DescribeNote(ctx context.Context, id uint64) (*note.Note, error) {
	query := sq.Select("id", "user_id", "classroom_id", "document_id").
		From(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var note note.Note

	if err := query.QueryRowContext(ctx).Scan(&note); err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *repo) ListNotes(ctx context.Context, count, offset uint64) ([]note.Note, error) {
	return nil, nil
}

func (r *repo) RemoveNote(ctx context.Context, id uint64) error {
	return nil
}
