package sync

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/spiermar/zerocoin/blockchain"

	"github.com/spiermar/zerocoin/proto"
)

type syncServer struct{}

func (s *syncServer) SyncLatest(stream proto.Synchronization_SyncLatestServer) error {
	for {
		latestBlock, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		blockchain.AddBlockToChain(latestBlock)
	}
}

func (s *syncServer) SyncAll(stream proto.Synchronization_SyncAllServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		latestBlockHeld := blockchain.GetLatestBlock()
		latestBlockReceived := in.Blockchain[len(in.Blockchain)-1]
		if latestBlockReceived.Index > latestBlockHeld.Index {
			blockchain.ReplaceChain(in.Blockchain)
		}
	}
}

// InitServer initializes the sync server
func InitServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// var opts []grpc.ServerOption
	grpcServer := grpc.NewServer()
	proto.RegisterSynchronizationServer(grpcServer, &syncServer{})
	grpcServer.Serve(lis)
}
