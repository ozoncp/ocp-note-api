package api

import (
	"context"

	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	log.Info().Msg("Create note ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	note := &note.Note{
		UserId:      uint32(request.UserId),
		ClassroomId: uint32(request.ClassroomId),
		DocumentId:  uint32(request.DocumentId),
	}

	noteId, err := a.repo.AddNote(ctx, note)

	if err != nil {
		log.Error().Err(err).Msg("failed to create note")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info().Msgf("Create note success (id: %d)", noteId)

	return &desc.CreateNoteV1Response{NoteId: noteId}, nil
}

func (a *api) DescribeNoteV1(ctx context.Context, request *desc.DescribeNoteV1Request) (*desc.DescribeNoteV1Response, error) {
	log.Info().Msg("Desribe note ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	note, err := a.repo.DescribeNote(ctx, uint64(request.NoteId))

	if err != nil {
		log.Error().Err(err).Msg("failed to get description note")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	log.Info().Msg("Desribe note success")

	return &desc.DescribeNoteV1Response{
		Note: &desc.Note{
			Id:          note.Id,
			UserId:      note.UserId,
			ClassroomId: note.ClassroomId,
			DocumentId:  note.DocumentId,
		},
	}, nil
}

func (a *api) ListNotesV1(ctx context.Context, request *desc.ListNotesV1Request) (*desc.ListNotesV1Response, error) {
	log.Info().Msg("List notes ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	notes, err := a.repo.ListNotes(ctx, uint64(request.Limit), uint64(request.Offset))

	if err != nil {
		log.Error().Err(err).Msg("failed to get notes")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var notesProto []*desc.Note

	for _, note := range notes {
		noteProto := &desc.Note{
			Id:          note.Id,
			UserId:      note.UserId,
			ClassroomId: note.ClassroomId,
			DocumentId:  note.DocumentId,
		}

		notesProto = append(notesProto, noteProto)
	}

	log.Info().Msg("List notes success")

	return &desc.ListNotesV1Response{Notes: notesProto}, nil
}

func (a *api) RemoveNoteV1(ctx context.Context, request *desc.RemoveNoteV1Request) (*desc.RemoveNoteV1Response, error) {
	log.Info().Msgf("Remove note (id: %d) ...", request.NoteId)

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := a.repo.RemoveNote(ctx, uint64(request.NoteId)); err != nil {
		log.Error().Err(err).Msg("failed to remove note")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info().Msgf("Remove note (id: %d) success", request.NoteId)

	return &desc.RemoveNoteV1Response{Found: true}, nil
}
