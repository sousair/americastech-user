package main

import (
	"fmt"
	"net"

	"github.com/joho/godotenv"
	grpc_handlers "github.com/sousair/americastech-user/internal/presentation/grpc/handlers"
	"github.com/sousair/americastech-user/internal/presentation/grpc/pb"
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

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &grpc_handlers.Server{})

	fmt.Println("gRPC Server is running on port 9090")

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
