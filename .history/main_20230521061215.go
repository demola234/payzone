package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/demola234/payzone/api"
	db "github.com/demola234/payzone/db/sqlc"
	_ "github.com/demola234/payzone/doc/statik"
	"github.com/demola234/payzone/gapi"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	_ "github.com/demola234/payzone/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	configs, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}
	if configs.Enviroment == "developement" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	conn, err := sql.Open(configs.DBDriver, configs.DBSource)
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}

	store := db.NewStore(conn)
	go runGateWayServer(configs, store)
	runGRPCServer(configs, store)

}

func runGateWayServer(configs utils.Config, store db.Store) {
	server, err := gapi.NewServer(configs, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	gRPCMux := runtime.NewServeMux(jsonOpt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterPayzoneHandlerServer(ctx, gRPCMux, server)
	if err != nil {
		log.Info().Msg("cannot register gateway server: ")
	}

	mux := http.NewServeMux()
	mux.Handle("/", gRPCMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msg("cannot create statik file system")
	}

	fs := http.FileServer(statikFS)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", fs))

	// Listen and serve
	listener, err := net.Listen("tcp", configs.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start HTTP listener")
	}

	
	log.Info().Msgf("starting HTTP server on %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Msg("cannot start HTTP GateWay server")
	}
}

func runGRPCServer(configs utils.Config, store db.Store) {
	server, err := gapi.NewServer(configs, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	gRPCLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	gRPCServer := grpc.NewServer(gRPCLogger)
	pb.RegisterPayzoneServer(gRPCServer, server)

	// Register reflection service on gRPC server.
	reflection.Register(gRPCServer)

	// Listen and serve
	listener, err := net.Listen("tcp", configs.GRPCServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start grpc listener")
	}
	log.Info().Msgf("starting gRPC server on %s", configs.GRPCServerAddress)
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start grpc server")
	}
}

func runGinServer(configs utils.Config, store db.Store) {
	server, err := api.NewServer(configs, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	err = server.Start(configs.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start http server")
	}
}
