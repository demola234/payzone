package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/demola234/payzone/api"
	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/gapi"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configs, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(configs.DBDriver, configs.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	// runGinServer(configs, store)
	runGRPCServer(configs, store)

}

func runGateWayServer(configs utils.Config, store db.Store) {
	server, err := gapi.NewServer(configs, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	gRPCMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err pb.RegisterPayzoneHandlerServer(ctx, gRPCMux, server)

	// Listen and serve
	listener, err := net.Listen("tcp", configs.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot start grpc listener: ", err)
	}
	log.Printf("starting gRPC server on %s", configs.GRPCServerAddress)
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server: ", err)
	}
}

func runGRPCServer(configs utils.Config, store db.Store) {
	server, err := gapi.NewServer(configs, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterPayzoneServer(gRPCServer, server)

	// Register reflection service on gRPC server.
	reflection.Register(gRPCServer)

	// Listen and serve
	listener, err := net.Listen("tcp", configs.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot start grpc listener: ", err)
	}
	log.Printf("starting gRPC server on %s", configs.GRPCServerAddress)
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server: ", err)
	}
}

func runGinServer(configs utils.Config, store db.Store) {
	server, err := api.NewServer(configs, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(configs.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start http server: ", err)
	}
}
