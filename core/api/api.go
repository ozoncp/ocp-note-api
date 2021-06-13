package api

import (
	"context"

	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/internal/utils"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpNoteApiServer
	repo      repo.Repo
	chunkSize uint32
}

func NewOcpNoteApi(repo repo.Repo, chunkSize uint32) desc.OcpNoteApiServer {
	return &api{
		repo:      repo,
		chunkSize: chunkSize,
	}
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (a *api) CreateNoteV1(ctx context.Context, request *desc.CreateNoteV1Request) (*desc.CreateNoteV1Response, error) {
	log.Info().Msg("Create note ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, err
	}

	note := &note.Note{
		UserId:      uint32(request.UserId),
		ClassroomId: uint32(request.ClassroomId),
		DocumentId:  uint32(request.DocumentId),
	}

	noteId, err := a.repo.AddNote(ctx, note)

	if err != nil {
		log.Error().Err(err).Msg("failed to create note")
		return nil, err
	}

	log.Info().Msgf("Create note success (id: %d)", noteId)

	return &desc.CreateNoteV1Response{NoteId: noteId}, nil
}

func (a *api) MultiCreateNotesV1(ctx context.Context, request *desc.MultiCreateNotesV1Request) (*desc.MultiCreateNotesV1Response, error) {
	log.Info().Msg("Multi create notes ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, err
	}

	var notes []note.Note

	for _, val := range request.Notes {

		note := &note.Note{
			UserId:      uint32(val.UserId),
			ClassroomId: uint32(val.ClassroomId),
			DocumentId:  uint32(val.DocumentId),
		}

		notes = append(notes, *note)
	}

	chunks := utils.SplitNoteSlice(notes, a.chunkSize)

	var (
		successPos                  = 0
		numberOfNotesCreated uint64 = 0
	)

	for _, val := range chunks {

		num, err := a.repo.AddNotes(ctx, val)

		if err != nil {
			log.Error().Err(err).Msg("failed to multi create notes")
			return nil, err
		}

		successPos += len(val)
		numberOfNotesCreated += num
	}

	log.Info().Msgf("Multi create notes success")

	return &desc.MultiCreateNotesV1Response{
		NumberOfNotesCreated: numberOfNotesCreated,
	}, nil
}

func (a *api) UpdateNoteV1(ctx context.Context, request *desc.UpdateNoteV1Request) (*desc.UpdateNoteV1Response, error) {
	log.Info().Msgf("Update note (id: %d) ...", request.Note.Id)

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, err
	}

	note := &note.Note{
		Id:          uint64(request.Note.Id),
		UserId:      uint32(request.Note.UserId),
		ClassroomId: uint32(request.Note.ClassroomId),
		DocumentId:  uint32(request.Note.DocumentId),
	}

	if err := a.repo.UpdateNote(ctx, note); err != nil {
		log.Error().Err(err).Msg("failed to update note")
		return &desc.UpdateNoteV1Response{Found: false}, nil
	}

	log.Info().Msgf("Update note (id: %d) success", request.Note.Id)

	return &desc.UpdateNoteV1Response{Found: true}, nil
}

func (a *api) DescribeNoteV1(ctx context.Context, request *desc.DescribeNoteV1Request) (*desc.DescribeNoteV1Response, error) {
	log.Info().Msg("Desribe note ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, err
	}

	note, err := a.repo.DescribeNote(ctx, uint64(request.NoteId))

	if err != nil {
		log.Error().Err(err).Msg("failed to get description note")
		return nil, err
	}

	log.Info().Msg("Desribe note success")

	return &desc.DescribeNoteV1Response{
		Note: &desc.Note{
			Id:          int64(note.Id),
			UserId:      int32(note.UserId),
			ClassroomId: int32(note.ClassroomId),
			DocumentId:  int32(note.DocumentId),
		},
	}, nil
}

func (a *api) ListNotesV1(ctx context.Context, request *desc.ListNotesV1Request) (*desc.ListNotesV1Response, error) {
	log.Info().Msg("List notes ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, err
	}

	notes, err := a.repo.ListNotes(ctx, uint64(request.Limit), uint64(request.Offset))

	if err != nil {
		log.Error().Err(err).Msg("failed to get notes")
		return nil, err
	}

	var notesProto []*desc.Note

	for _, note := range notes {
		noteProto := &desc.Note{
			Id:          int64(note.Id),
			UserId:      int32(note.UserId),
			ClassroomId: int32(note.ClassroomId),
			DocumentId:  int32(note.DocumentId),
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
		return nil, err
	}

	if err := a.repo.RemoveNote(ctx, uint64(request.NoteId)); err != nil {
		log.Error().Err(err).Msg("failed to remove note")
		return &desc.RemoveNoteV1Response{Found: false}, nil
	}

	log.Info().Msgf("Remove note (id: %d) success", request.NoteId)

	return &desc.RemoveNoteV1Response{Found: true}, nil
}
