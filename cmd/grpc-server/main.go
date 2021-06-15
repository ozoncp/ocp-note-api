package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-note-api/core/api"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/internal/metrics"
	"github.com/ozoncp/ocp-note-api/internal/producer"
	"github.com/ozoncp/ocp-note-api/internal/tracer"
	note "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

const (
	grpcPort    = ":7002"
	httpPort    = ":8080"
	messagePort = ":9100"
	chunkSize   = 2

	host     = "database"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"

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

	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
		return err
	}

	repo := repo.New(*db, chunkSize)
	dataProducer, err := producer.New(ctx, topic)

	if err != nil {
		log.Error().Err(err).Msg("failed to create a producer")
		return err
	}

	note.RegisterOcpNoteApiServer(grpcServer, api.NewOcpNoteApi(repo, dataProducer))

	var group errgroup.Group

	group.Go(func() error {
		log.Info().Msg("serving grpc requests...")
		return grpcServer.Serve(listen)
	})

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	group.Go(func() error {
		if err := note.RegisterOcpNoteApiHandlerFromEndpoint(ctx, gwmux, grpcPort, opts); err != nil {
			log.Error().Msgf("register gateway fails: %v", err)
			return err
		}

		mux := http.NewServeMux()
		mux.Handle("/", gwmux)

		log.Info().Msgf("http server listening on %s", httpPort)
		if err = http.ListenAndServe(httpPort, mux); err != nil {
			log.Error().Msgf("http gateway server fails: %v", err)
			return err
		}

		return nil
	})

	group.Go(func() error {
		metrics.RegisterMetrics()

		http.Handle("/metrics", promhttp.Handler())
		log.Info().Msgf("metrics (http) listening on %s", messagePort)

		if err = http.ListenAndServe(messagePort, nil); err != nil {
			log.Error().Msgf("metrics (http) server fails: %v", err)
			return err
		}

		return nil
	})

	return group.Wait()
}

func main() {
	tracer.InitTracing("ocp_note_api")

	if err := run(); err != nil {
		log.Fatal().Err(err).Msgf("failed to run service")
	}
}
