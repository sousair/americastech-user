package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	jwt_provider "github.com/sousair/americastech-user/internal/infra/jwt"
	grpc_handlers "github.com/sousair/americastech-user/internal/presentation/grpc/handlers"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("/.env")

	if err != nil {
		panic(err)
	}

	var (
		userSecret = os.Getenv("USER_TOKEN_SECRET")

		port = os.Getenv("PORT")
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		panic(err)
	}

	jwtProvider := jwt_provider.NewJwtProvider(userSecret)
	server := grpc.NewServer(grpc.EmptyServerOption{})

	grpc_handlers.NewUserServiceServer(server, jwtProvider)

	fmt.Println("gRPC Server is running on port 9090")

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
