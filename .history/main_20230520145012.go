package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/demola234/payzone/api"
	db "github.com/demola234/payzone/db/sqlc"
	_ "github.com/demola234/payzone/doc/statik"
	"github.com/demola234/payzone/gapi"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	_ "github.com/demola234/payzone/utils"
	"github.com/foxnut/go-hdwallet"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
)

// Create a usdt wallet for a user
func main() {
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)
	if err != nil {
		panic(err)
	}
	// configs, err := utils.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config: ", err)
	// }

	// conn, err := sql.Open(configs.DBDriver, configs.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db: ", err)
	// }

	// store := db.NewStore(conn)
	// go runGateWayServer(configs, store)
	// runGRPCServer(configs, store)
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.USDC), hdwallet.AddressIndex(1))
	address, _ := wallet.GetAddress()
	addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
	addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()
	fmt.Println("USDC: ", address, addressP2WPKH, addressP2WPKHInP2SH)
}

func transferUsdc(fromAccount string, toAccount string, amount float64) {
	

}

func runGateWayServer(configs utils.Config, store db.Store) {
	server, err := gapi.NewServer(configs, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
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
		log.Fatal("cannot register gateway server: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gRPCMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("cannot create statik file system: ", err)
	}

	fs := http.FileServer(statikFS)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", fs))

	// Listen and serve
	listener, err := net.Listen("tcp", configs.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start HTTP listener: ", err)
	}

	log.Printf("starting HTTP server on %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP GateWay server: ", err)
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
