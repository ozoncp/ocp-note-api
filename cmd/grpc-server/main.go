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
	note "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

const (
	grpcPort  = ":82"
	httpPort  = ":8080"
	promPort  = ":9100"
	chunkSize = 2

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

	note.RegisterOcpNoteApiServer(grpcServer, api.NewOcpNoteApi(repo, dataProducer, chunkSize))

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
		log.Info().Msgf("metrics (http) listening on %s", promPort)

		if err = http.ListenAndServe(promPort, nil); err != nil {
			log.Error().Msgf("metrics (http) server fails: %v", err)
			return err
		}

		return nil
	})

	return group.Wait()
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msgf("failed to create grpc server")
	}
}
