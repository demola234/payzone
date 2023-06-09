package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"

	"github.com/demola234/payzone/api"
	db "github.com/demola234/payzone/db/sqlc"
	_ "github.com/demola234/payzone/doc/statik"
	"golang.org/x/crypto/sha3""
	"github.com/demola234/payzone/gapi"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	_ "github.com/demola234/payzone/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	toWallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.USDC), hdwallet.AddressIndex(1))
	toaddress, _ := toWallet.GetAddress()
	toaddressP2WPKH, _ := toWallet.GetKey().AddressP2WPKH()
	toaddressP2WPKHInP2SH, _ := toWallet.GetKey().AddressP2WPKHInP2SH()
	fmt.Println("USDC: ", toaddress, toaddressP2WPKH, toaddressP2WPKHInP2SH)

	// 2. Transfer USDC to the user
	transferUsdc(address, toaddress)
}

func transferUsdc(fromAccount string, toAccount string) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("0x" + fromAccount)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(toAccount)
	usdtTokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, usdtTokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

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
