package grpc_handlers

import (
	"context"

	"github.com/sousair/americastech-user/internal/application/providers/cryptography"
	"github.com/sousair/americastech-user/internal/presentation/grpc/pb"
	"google.golang.org/grpc"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	cryptoProvider cryptography.CryptoProvider
}

func NewUserServiceServer(grpcServer *grpc.Server, cryptoProvider cryptography.CryptoProvider) {
	userServer := &UserServiceServer{
		cryptoProvider: cryptoProvider,
	}

	pb.RegisterUserServiceServer(grpcServer, userServer)
}

func (h UserServiceServer) ValidateUserToken(ctx context.Context, req *pb.ValidateUserTokenRequest) (*pb.ValidateUserTokenResponse, error) {
	userToken := req.GetToken()

	if userToken == "" {
		return &pb.ValidateUserTokenResponse{
			Valid: false,
		}, nil
	}

	_, err := h.cryptoProvider.VerifyAuthToken(userToken)

	if err != nil {
		return &pb.ValidateUserTokenResponse{
			Valid: false,
		}, nil

	}

	return &pb.ValidateUserTokenResponse{
		Valid: true,
	}, nil
}
