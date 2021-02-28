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

func (*UserCenterServer) Login(context.Context, *h.LoginRequest) (*h.LoginResponse, error) {
	fmt.Println("Login called")
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (*UserCenterServer) Delete(context.Context, *h.DeleteRequest) (*h.DeleteResponse, error) {
	fmt.Println("Delete called")
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (*UserCenterServer) CheckToken(context.Context, *h.CheckTokenRequest) (*h.CheckTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
