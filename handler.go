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
	fmt.Println("CheckToken called")
	return handler.CheckToken(ctx, req)
}

func (*UserCenterServer) Refresh(ctx context.Context, req *h.RefreshRequest) (*h.RefreshResponse, error) {
	fmt.Println("Refresh called")
	return handler.Refresh(ctx, req)
}

func (*UserCenterServer) CreateGroup(ctx context.Context, req *h.CreateGroupRequest) (*h.CreateGroupResponse, error) {
	fmt.Println("CreateGroup called")
	return handler.CreateGroup(ctx, req)
}

func (*UserCenterServer) GetGroup(ctx context.Context, req *h.GetGroupRequest) (*h.GetGroupResponse, error) {
	fmt.Println("GetGroup called")
	return handler.GetGroup(ctx, req)
}

func (*UserCenterServer) GetGroupMembers(ctx context.Context, req *h.GetGroupMembersRequest) (*h.GetGroupMembersResponse, error) {
	fmt.Println("GetGroupMembers called")
	return handler.GetGroupMembers(ctx, req)
}

func (*UserCenterServer) JoinGroup(ctx context.Context, req *h.JoinGroupRequest) (*h.JoinGroupResponse, error) {
	fmt.Println("JoinGroup called")
	return handler.JoinGroup(ctx, req)
}

func (*UserCenterServer) ExitGroup(ctx context.Context, req *h.ExitGroupRequest) (*h.ExitGroupResponse, error) {
	fmt.Println("ExitGroup called")
	return handler.ExitGroup(ctx, req)
}
