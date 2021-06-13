package repo

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-note-api/core/note"
)

type Repo interface {
	AddNote(ctx context.Context, note *note.Note) (uint64, error)
	AddNotes(ctx context.Context, notes []note.Note) (uint64, error)
	UpdateNote(ctx context.Context, notes *note.Note) error
	DescribeNote(ctx context.Context, id uint64) (*note.Note, error)
	ListNotes(ctx context.Context, limit, offset uint64) ([]note.Note, error)
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

func (r *repo) AddNote(ctx context.Context, note *note.Note) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("user_id", "classroom_id", "document_id").
		Values(note.UserId, note.ClassroomId, note.DocumentId).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&note.Id)

	if err != nil {
		return 0, err
	}

	return note.Id, nil
}

func (r *repo) AddNotes(ctx context.Context, notes []note.Note) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("user_id", "classroom_id", "document_id").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, note := range notes {
		query = query.Values(note.UserId, note.ClassroomId, note.DocumentId)
	}

	result, err := query.ExecContext(ctx)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return uint64(rowsAffected), err
}

func (r *repo) UpdateNote(ctx context.Context, note *note.Note) error {
	query := sq.Update(tableName).
		Set("user_id", note.UserId).
		Set("classroom_id", note.ClassroomId).
		Set("document_id", note.DocumentId).
		Where(sq.Eq{"id": note.Id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("not found note")
	}

	return nil
}

func (r *repo) DescribeNote(ctx context.Context, id uint64) (*note.Note, error) {
	query := sq.Select("id", "user_id", "classroom_id", "document_id").
		From(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var note note.Note

	if err := query.QueryRowContext(ctx).Scan(&note.Id, &note.UserId, &note.ClassroomId, &note.DocumentId); err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *repo) ListNotes(ctx context.Context, limit, offset uint64) ([]note.Note, error) {
	query := sq.Select("id", "user_id", "classroom_id", "document_id").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	var notes []note.Note

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note note.Note
		err = rows.Scan(&note.Id, &note.UserId, &note.ClassroomId, &note.DocumentId)

		if err != nil {
			continue
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (r *repo) RemoveNote(ctx context.Context, id uint64) error {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("not found note")
	}

	return nil
}
