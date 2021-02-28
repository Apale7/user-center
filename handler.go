package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user_center/handler"
	h "user_center/proto/user-center"
)

type UserCenterServer struct{}

func (*UserCenterServer) Register(ctx context.Context, req *h.RegisterRequest) (*h.RegisterResponse, error) {
	fmt.Println("Register called")
	return handler.Register(ctx, req)
}

func (*UserCenterServer) Login(ctx context.Context, req *h.LoginRequest) (*h.LoginResponse, error) {
	fmt.Println("Login called")
	return handler.Login(ctx, req)
}

func (*UserCenterServer) Delete(ctx context.Context, req *h.DeleteRequest) (*h.DeleteResponse, error) {
	fmt.Println("Delete called")
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (*UserCenterServer) CheckToken(ctx context.Context, req *h.CheckTokenRequest) (*h.CheckTokenResponse, error) {
	return handler.CheckToken(ctx, req)
}
