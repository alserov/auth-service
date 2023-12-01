package auth

import (
	"context"
	authv1 "github.com/alserov/auth-service/grpc/gen/auth"
	"google.golang.org/grpc"
)

type serverAPI struct {
	authv1.UnimplementedAuthServer /*to run without implementation all methods*/
}

func Register(gRPC *grpc.Server) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	panic("")
}

func (s *serverAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	panic("")
}
