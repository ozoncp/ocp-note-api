package repo

import (
	"context"
	"errors"
	"fmt"
	"unsafe"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/internal/utils"
)

type Repo interface {
	AddNote(ctx context.Context, note *note.Note) (uint64, error)
	MultiAddNotes(ctx context.Context, notes []note.Note) (uint64, error)
	UpdateNote(ctx context.Context, notes *note.Note) error
	DescribeNote(ctx context.Context, id uint64) (*note.Note, error)
	ListNotes(ctx context.Context, limit, offset uint64) ([]note.Note, error)
	RemoveNote(ctx context.Context, id uint64) (error, bool)
}

const (
	tableName = "notes"
)

type repo struct {
	db        sqlx.DB
	chunkSize uint32
}

func New(db sqlx.DB, chunkSize uint32) Repo {
	return &repo{
		db:        db,
		chunkSize: chunkSize,
	}
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

func (r *repo) MultiAddNotes(ctx context.Context, notes []note.Note) (uint64, error) {

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiAddNotes global")
	defer span.Finish()

	chunks := utils.SplitNoteSlice(notes, r.chunkSize)

	var numberOfNotesCreated int64 = 0

	for index, val := range chunks {
		err := func() error {
			// Create a Child Span. Note that we're using the ChildOf option.
			childSpan := tracer.StartSpan(
				fmt.Sprintf("MultiAddNotes for chunk %d, count of bytes: %d", index, len(val)*int(unsafe.Sizeof(note.Note{}))),
				opentracing.ChildOf(span.Context()),
			)
			defer childSpan.Finish()

			query := sq.Insert(tableName).
				Columns("user_id", "classroom_id", "document_id").
				RunWith(r.db).
				PlaceholderFormat(sq.Dollar)

			for _, note := range val {
				query = query.Values(note.UserId, note.ClassroomId, note.DocumentId)
			}

			result, err := query.ExecContext(ctx)

			if err != nil {
				return err
			}

			rowsAffected, err := result.RowsAffected()

			if err != nil {
				return err
			}

			numberOfNotesCreated = numberOfNotesCreated + rowsAffected
			return nil
		}()

		if err != nil {
			return uint64(numberOfNotesCreated), err
		}
	}

	return uint64(numberOfNotesCreated), nil
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

func (r *repo) RemoveNote(ctx context.Context, id uint64) (error, bool) {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)

	if err != nil {
		return err, false
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err, false
	}

	if rowsAffected <= 0 {
		return nil, false
	}

	return nil, true
}
