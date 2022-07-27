package main

import (
	"context"
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/marcioecom/grpc-go/database/models"
	"github.com/marcioecom/grpc-go/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.CryptoVotingServiceClient

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	client = pb.NewCryptoVotingServiceClient(connection)

	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format:   "[${time}]:${status}- [${method}] ${path} ${latency} \n",
		TimeZone: "America/Sao_Paulo",
	}))

	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})

	app.Get("/", getAllCryptos)
	app.Get("/:id", getCrypto)
	app.Post("/", createCrypto)
	app.Put("/up/:id", upvoteCrypto)
	app.Put("/down/:id", downvoteCrypto)

	app.Listen(":3001")
}

func createCrypto(c *fiber.Ctx) error {
	var crypto models.Crypto
	err := c.BodyParser(&crypto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error parsing JSON: %v", err),
		})
	}

	req := &pb.Crypto{
		Name: crypto.Name,
	}

	res, err := client.CreateCrypto(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error sending grpc message: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func getCrypto(c *fiber.Ctx) error {
	id := c.Params("id")

	req := &pb.GetCryptoRequest{Id: id}

	res, err := client.GetCrypto(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error sending grpc message: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func getAllCryptos(c *fiber.Ctx) error {
	var cryptos []models.Crypto

	req := &pb.GetAllCryptosRequest{}

	stream, err := client.GetAllCryptos(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error sending grpc message: %v", err),
		})
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("Error: %v", err),
			})
		}

		oid, _ := primitive.ObjectIDFromHex(res.Id)
		item := models.Crypto{
			ID:    oid,
			Name:  res.Name,
			Up:    res.Up,
			Down:  res.Down,
			Total: res.Total,
		}
		cryptos = append(cryptos, item)
	}

	return c.Status(fiber.StatusCreated).JSON(cryptos)
}

func upvoteCrypto(c *fiber.Ctx) error {
	id := c.Params("id")

	req := &pb.UpvoteCryptoRequest{Id: id}

	res, err := client.UpvoteCrypto(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error sending grpc message: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func downvoteCrypto(c *fiber.Ctx) error {
	id := c.Params("id")

	req := &pb.DownvoteCryptoRequest{Id: id}

	res, err := client.DownvoteCrypto(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Error sending grpc message: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
