package main

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-note-api/core/api"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
)

const (
	grpcPort = ":82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen: %v", err)
	}

	log.Info().Msgf("Starting server at localhost%v...", grpcPort)

	s := grpc.NewServer()
	desc.RegisterOcpNoteApiServer(s, api.NewOcpNoteApi())

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msgf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msgf("failed to create grpc server")
	}
}
