package api

import (
	"context"

	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpNoteApiServer
	repo repo.Repo
}

func NewOcpNoteApi(repo repo.Repo) desc.OcpNoteApiServer {
	return &api{repo: repo}
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (a *api) CreateNoteV1(ctx context.Context, request *desc.CreateNoteV1Request) (*desc.CreateNoteV1Response, error) {
	log.Print("Create note ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, err
	}

	note := &note.Note{
		UserId:      request.UserId,
		ClassroomId: request.ClassroomId,
		DocumentId:  request.DocumentId,
	}

	noteId, err := a.repo.AddNote(ctx, note)

	if err != nil {
		log.Error().Err(err).Msg("failed to create note")
		return nil, err
	}

	log.Info().Msgf("Create note success (id: %d)", noteId)

	return &desc.CreateNoteV1Response{NoteId: noteId}, nil
}

func (a *api) DescribeNoteV1(ctx context.Context, request *desc.DescribeNoteV1Request) (*desc.DescribeNoteV1Response, error) {
	log.Print("Desribe note", request)

	return nil, nil
}

func (a *api) ListNotesV1(ctx context.Context, request *desc.ListNotesV1Request) (*desc.ListNotesV1Response, error) {
	log.Print("List notes")

	return nil, nil
}

func (a *api) RemoveNoteV1(ctx context.Context, request *desc.RemoveNoteV1Request) (*desc.RemoveNoteV1Response, error) {
	log.Print("Remove note")

	return nil, nil
}
