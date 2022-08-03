package server

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/marcioecom/grpc-go/database"
	"github.com/marcioecom/grpc-go/pb"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	database.Init()
	mongoCtx := database.GetContext()
	collection := database.GetCollection("crypto_test")
	collection.DeleteMany(mongoCtx, bson.M{"name": "testCrypto"})
	service := New(collection, mongoCtx)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterCryptoVotingServiceServer(s, service)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func createClient(ctx context.Context, t *testing.T) pb.CryptoVotingServiceClient {
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewCryptoVotingServiceClient(conn)

	return client
}

var cryptoTestId string

func TestCreateCrypto(t *testing.T) {
	ctx := context.Background()
	client := createClient(ctx, t)

	req := &pb.Crypto{Name: "testCrypto"}

	result, err := client.CreateCrypto(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create new crypto: %v", err)
	}

	cryptoTestId = result.Crypto.GetId()
	log.Printf("Created crypto: %v", result)
}

func TestGetCrypto(t *testing.T) {
	ctx := context.Background()
	client := createClient(ctx, t)

	req := &pb.GetCryptoRequest{Id: cryptoTestId}

	result, err := client.GetCrypto(ctx, req)
	if err != nil {
		t.Fatalf("Failed to get crypto: %v", err)
	}

	log.Printf("Got crypto: %v", result)
}

func TestUpvoteCrypto(t *testing.T) {
	ctx := context.Background()
	client := createClient(ctx, t)

	req := &pb.UpvoteCryptoRequest{Id: cryptoTestId}

	result, err := client.UpvoteCrypto(ctx, req)
	if err != nil {
		t.Fatalf("Failed to upvote crypto: %v", err)
	}

	log.Printf("Upvoted crypto: %v", result)
}

func TestDownvoteCrypto(t *testing.T) {
	ctx := context.Background()
	client := createClient(ctx, t)

	req := &pb.DownvoteCryptoRequest{Id: cryptoTestId}

	result, err := client.DownvoteCrypto(ctx, req)
	if err != nil {
		t.Fatalf("Failed to downvote crypto: %v", err)
	}

	log.Printf("Downvoted crypto: %v", result)
}

func TestClearCryptoVotes(t *testing.T) {
	ctx := context.Background()
	client := createClient(ctx, t)

	req := &pb.ClearVotesCryptoRequest{Id: cryptoTestId}

	result, err := client.ClearVotes(ctx, req)
	if err != nil {
		t.Fatalf("Failed to clear crypto votes: %v", err)
	}

	log.Printf("Clear crypto votes: %v", result)
}
