package api

import (
	"context"

	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpNoteApiServer
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (a *api) CreateNoteV1(ctx context.Context, request *desc.CreateNoteV1Request) (*desc.CreateNoteV1Response, error) {
	log.Print("Create note ", request)

	return nil, nil
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

func NewOcpNoteApi() desc.OcpNoteApiServer {
	return &api{}
}
