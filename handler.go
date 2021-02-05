package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	h "user_center/proto/user-center"
)

type UserCenterServer struct{}

func (*UserCenterServer) Register(context.Context, *h.RegisterRequest) (*h.RegisterResponse, error) {
	fmt.Println("Register called")
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (*UserCenterServer) Login(context.Context, *h.RegisterRequest) (*h.LoginResponse, error) {
	fmt.Println("Login called")
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (*UserCenterServer) Logout(context.Context, *h.LogoutRequest) (*h.LogoutResponse, error) {
	fmt.Println("Logout called")
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
