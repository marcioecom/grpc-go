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

func (s *server) DownvoteCrypto(ctx context.Context, req *pb.DownvoteCryptoRequest) (*pb.DownvoteCryptoResponse, error) {
	var crypto models.Crypto

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid ObjectId: %v", err.Error()))
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$inc": bson.M{"down": 1, "total": -1}}

	err = s.collection.FindOneAndUpdate(s.mongoCtx, filter, update).Decode(&crypto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Crypto not found")
		}
		log.Fatal(err)
	}

	return &pb.DownvoteCryptoResponse{Ok: true}, nil
}

func (s *server) GetCrypto(ctx context.Context, req *pb.GetCryptoRequest) (*pb.GetCryptoResponse, error) {
	var crypto models.Crypto

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid ObjectId: %v", err.Error()))
	}

	err = s.collection.FindOne(s.mongoCtx, bson.M{"_id": oid}).Decode(&crypto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Crypto not found")
		}
		log.Fatal(err)
	}

	res := &pb.Crypto{
		Id:    crypto.ID.Hex(),
		Name:  crypto.Name,
		Up:    crypto.Up,
		Down:  crypto.Down,
		Total: crypto.Total,
	}

	return &pb.GetCryptoResponse{Crypto: res}, nil
}

func (s *server) GetAllCryptos(req *pb.GetAllCryptosRequest, stream pb.CryptoVotingService_GetAllCryptosServer) error {
	var cryptos []models.Crypto

	cursor, err := s.collection.Find(s.mongoCtx, bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error to find cryptos: %v", err))
	}

	if err = cursor.All(s.mongoCtx, &cryptos); err != nil {
		return status.Errorf(codes.Unavailable, fmt.Sprintf("Error to decode cryptos: %v", err))
	}

	for _, v := range cryptos {
		stream.Send(&pb.Crypto{
			Id:    v.ID.Hex(),
			Name:  v.Name,
			Up:    v.Up,
			Down:  v.Down,
			Total: v.Total,
		})
	}

	return nil
}
