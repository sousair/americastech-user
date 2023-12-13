package main

import (
	"fmt"
	"net"

	"github.com/joho/godotenv"
	crypto_provider "github.com/sousair/americastech-user/internal/infra/cryptography"
	grpc_handlers "github.com/sousair/americastech-user/internal/presentation/grpc/handlers"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("/.env")

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		panic(err)
	}

	cryptoProvider := crypto_provider.NewCryptoProvider()
	server := grpc.NewServer()

	grpc_handlers.NewUserServiceServer(server, cryptoProvider)

	fmt.Println("gRPC Server is running on port 9090")

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
