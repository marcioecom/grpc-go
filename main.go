package main

import (
	"fmt"
	"log"
	"net"

	db "github.com/marcioecom/grpc-go/database"
	"github.com/marcioecom/grpc-go/pb"
	"github.com/marcioecom/grpc-go/server"
	"google.golang.org/grpc"
)

func main() {
	db.Init()

	mongoCtx := db.GetContext()
	collection := db.GetCollection("cryptos")
	upvoteService := server.New(collection, mongoCtx)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCryptoVotingServiceServer(grpcServer, upvoteService)

	fmt.Println("Server running on :50051")
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
