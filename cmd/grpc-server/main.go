package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-note-api/core/api"
	"github.com/ozoncp/ocp-note-api/core/repo"
	note "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

var grpcPort int

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "inferno04"
	dbname   = "testdb"
)

func init() {
	flag.IntVar(&grpcPort, "port", 7002, "GRPC server port")
}

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	grpcEndpoint := fmt.Sprintf("localhost:%d", grpcPort)

	lis, err := net.Listen("tcp", grpcEndpoint)

	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
	}

	log.Info().Msgf("Starting server at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

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
	note.RegisterOcpNoteApiServer(grpcServer, api.NewOcpNoteApi(repo, 2))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}
}
