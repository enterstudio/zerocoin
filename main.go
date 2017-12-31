package main

import (
	"flag"

	"github.com/spiermar/zerocoin/api"
	"github.com/spiermar/zerocoin/blockchain"
	"github.com/spiermar/zerocoin/sync"
)

var (
	grpcPort = flag.Int("grpc-port", 10000, "The gRPC server port")
	httpPort = flag.Int("http-port", 1323, "The HTTP server port")
)

func main() {
	flag.Parse()

	blockchain.GenerateGenesisBlock("Hello, World!")

	api.InitServer(*httpPort)

	sync.InitServer(*grpcPort)
}
