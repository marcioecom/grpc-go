package main

import (
	"context"
	"fmt"
	"log"

	"github.com/marcioecom/grpc-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	client := pb.NewCryptoVotingServiceClient(connection)

	// req := &pb.Crypto{
	// 	Name: "BTC",
	// }

	// res, err := client.CreateCrypto(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("Error to update crypto: %v", err)
	// }

	req := &pb.DownvoteCryptoRequest{
		Id: "62df0cf403c9472ca7bf3ee9",
	}

	res, err := client.DownvoteCrypto(context.Background(), req)
	if err != nil {
		log.Fatalf("Error to update crypto: %v", err)
	}

	fmt.Println(res)
}
