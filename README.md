# Golang Upvote Service

Cryptocurrency upvote microservice using golang and gRPC

## 📋 Prerequisites
Versions I've used:
- Go 1.18.3
- MongoDB 5.0.9
- Protoc 3.6.1

## 🔧 Installation

Clone the repository

```
$ git clone https://github.com/marcioecom/grpc-go.git
```

Access the project folder, and download the Go dependencies

```
$ go get ./...
```

## 🚀 Running

```
$ docker-compose up -d
```

## 📄 Endpoints

### Create Crypto
> POST http://localhost:3001/ 

Body:
~~~~json
{
  "name": "BTC"
}
~~~~
Response:
~~~~json
{
  "crypto": {
    "id": "62e212efafcd6a39f91a9c33",
    "name": "BTC"
  }
}
~~~~

### Get crypto
> GET http://localhost:3001/:id

Response:
~~~~json
{
  "crypto": {
    "id": "62e1e29b17ffc1b9e583c3c1",
    "name": "BTC",
    "up": 4,
    "down": 1,
    "total": 3
  }
}
~~~~

### Get all cryptos
> GET http://localhost:3001/

Response:
~~~~json
[
  {
    "id": "62e212efafcd6a39f91a9c33",
    "name": "BTC",
    "up": 4,
    "down": 1,
    "total": 3
  },
  {
    "id": "62e21310afcd6a39f91a9c34",
    "name": "ETH",
    "up": 2,
    "down": 1,
    "total": 1
  }
]
~~~~

### Upvote crypto
> PUT http://localhost:3001/up/:id

Response:
~~~~json
{
  "ok": true
}
~~~~

### Downvote crypto
> PUT http://localhost:3001/down/:id

Response:
~~~~json
{
  "ok": true
}
~~~~

## 🛠️ Build with

* [Golang](https://go.dev/) - Linguagem usada para desenvolver
* [gRPC](https://grpc.io/) - Framework RPC de alta performance
* [MongoDB](https://www.mongodb.com/) - Banco de dados não relacional
