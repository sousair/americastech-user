package grpc_handlers

import (
	"context"

	crypto_provider "github.com/sousair/americastech-user/internal/infra/cryptography"
	"github.com/sousair/americastech-user/internal/presentation/grpc/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) ValidateUserToken(ctx context.Context, req *pb.ValidateUserTokenRequest) (*pb.ValidateUserTokenResponse, error) {
	userToken := req.GetToken()

	if userToken == "" {
		return &pb.ValidateUserTokenResponse{
			Valid: false,
		}, nil
	}

	cryptoProvider := crypto_provider.NewCryptoProvider()

	_, err := cryptoProvider.VerifyAuthToken(userToken)

	if err != nil {
		return &pb.ValidateUserTokenResponse{
			Valid: false,
		}, nil

	}

	return &pb.ValidateUserTokenResponse{
		Valid: true,
	}, nil
}
