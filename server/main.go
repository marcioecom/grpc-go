package server

import (
	"context"
	"fmt"
	"log"

	"github.com/marcioecom/grpc-go/database/models"
	"github.com/marcioecom/grpc-go/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	collection *mongo.Collection
	mongoCtx   context.Context
}

func New(collection *mongo.Collection, mongoCtx context.Context) *server {
	return &server{
		collection: collection,
		mongoCtx:   mongoCtx,
	}
}

func (s *server) CreateCrypto(ctx context.Context, req *pb.Crypto) (*pb.CreateCryptoResponse, error) {
	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is required")
	}

	if len(req.GetName()) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "Name must be at least 3 characters")
	}

	cryptoData := models.Crypto{
		Name: req.GetName(),
	}

	result, err := s.collection.InsertOne(s.mongoCtx, cryptoData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid := result.InsertedID.(primitive.ObjectID)
	req.Id = oid.Hex()
	return &pb.CreateCryptoResponse{Crypto: req}, nil
}

func (s *server) UpvoteCrypto(ctx context.Context, req *pb.UpvoteCryptoRequest) (*pb.UpvoteCryptoResponse, error) {
	var crypto models.Crypto

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid ObjectId: %v", err.Error()))
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$inc": bson.M{"up": 1, "total": 1}}

	err = s.collection.FindOneAndUpdate(s.mongoCtx, filter, update).Decode(&crypto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Crypto not found")
		}
		log.Fatal(err)
	}

	return &pb.UpvoteCryptoResponse{Ok: true}, nil
}
