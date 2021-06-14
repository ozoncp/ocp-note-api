package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-note-api/core/api"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/internal/producer"
	note "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

const (
	grpcPort = ":82"

	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "inferno04"
	dbname   = "testdb"

	topic = "noteTopic"
)

func run() error {
	ctx := context.Background()

	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen: %v", err)
	}

	log.Info().Msgf("Starting server at localhost%v...", grpcPort)

	grpcServer := grpc.NewServer()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", psqlInfo)
	defer db.Close()

	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
	}

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
	}

	repo := repo.New(*db)
	dataProducer, err := producer.New(ctx, topic)

	if err != nil {
		log.Error().Err(err).Msg("failed to create a producer")
	}

	note.RegisterOcpNoteApiServer(grpcServer, api.NewOcpNoteApi(repo, dataProducer, 2))

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal().Err(err).Msgf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msgf("failed to create grpc server")
	}
}
