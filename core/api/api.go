package api

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/internal/metrics"
	"github.com/ozoncp/ocp-note-api/internal/producer"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpNoteApiServer
	repo         repo.Repo
	dataProducer producer.Producer
}

func NewOcpNoteApi(repo repo.Repo, dataProducer producer.Producer) desc.OcpNoteApiServer {
	return &api{
		repo:         repo,
		dataProducer: dataProducer,
	}
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

	message := producer.CreateMessage(producer.Create, noteId, time.Now())
	err = a.dataProducer.Send(message)

	if err != nil {
		log.Warn().Msgf("failed to send message about creating a note to kafka: %v", err)
	}

	metrics.CreateCounterInc("Create")

	return &desc.CreateNoteV1Response{NoteId: noteId}, nil
}

func (a *api) MultiCreateNotesV1(ctx context.Context, request *desc.MultiCreateNotesV1Request) (*desc.MultiCreateNotesV1Response, error) {
	log.Info().Msg("Multi create notes ...")

	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
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

	numberOfNotesCreated, err := a.repo.MultiAddNotes(ctx, notes)

	if err != nil {
		log.Error().Err(err).Msg("failed to multi create notes")
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	note := &note.Note{
		Id:          uint64(request.Note.Id),
		UserId:      uint32(request.Note.UserId),
		ClassroomId: uint32(request.Note.ClassroomId),
		DocumentId:  uint32(request.Note.DocumentId),
	}

	err, found := a.repo.UpdateNote(ctx, note)

	if err != nil {
		log.Error().Err(err).Msg("failed to update note")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		log.Error().Err(err).Msg("not found note")
		return nil, status.Error(codes.NotFound, "not found note")
	}

	log.Info().Msgf("Update note (id: %d) success", request.Note.Id)

	message := producer.CreateMessage(producer.Update, note.Id, time.Now())
	err = a.dataProducer.Send(message)

	if err != nil {
		log.Warn().Msgf("failed to send message about updating a note to kafka: %v", err)
	}

	metrics.UpdateCounterInc("Update")

	return &desc.UpdateNoteV1Response{Found: true}, nil
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
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err, found := a.repo.RemoveNote(ctx, uint64(request.NoteId))

	if err != nil {
		log.Error().Err(err).Msg("failed to remove note")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		log.Error().Err(err).Msg("not found note")
		return nil, status.Error(codes.NotFound, "not found note")
	}

	log.Info().Msgf("Remove note (id: %d) success", request.NoteId)

	message := producer.CreateMessage(producer.Remove, uint64(request.NoteId), time.Now())
	err = a.dataProducer.Send(message)

	if err != nil {
		log.Warn().Msgf("failed to send message about deleting a note to kafka: %v", err)
	}

	metrics.RemoveCounterInc("Remove")

	return &desc.RemoveNoteV1Response{Found: true}, nil
}
