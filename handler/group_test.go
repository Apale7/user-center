package handler

import (
	"context"
	"testing"
	user_center "user_center/proto/user-center"
)

var ctx = context.Background()

func TestCreateGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user_center.CreateGroupRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.CreateGroupResponse
		wantErr  bool
	}{
		{
			name:     "create group 737",
			args:     args{ctx: ctx, req: &user_center.CreateGroupRequest{OwnerId: 1, GroupName: "737"}},
			wantResp: &user_center.CreateGroupResponse{Group: &user_center.Group{GroupName: "737"}},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := CreateGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.Group.GroupName != tt.wantResp.Group.GroupName {
				t.Errorf("CreateGroup() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user_center.GetGroupRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.GetGroupResponse
		wantErr  bool
	}{
		{
			name:    "test get group1",
			args:    args{ctx: ctx, req: &user_center.GetGroupRequest{GroupInfo: &user_center.Group{GroupName: "7"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotResp.Groups) != 2 {
				t.Errorf("GetGroup() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetGroupMembers(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user_center.GetGroupMembersRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.GetGroupMembersResponse
		wantErr  bool
	}{
		{
			name:    "get group members from group1",
			args:    args{ctx: ctx, req: &user_center.GetGroupMembersRequest{GroupId: 1}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetGroupMembers(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotResp.Members) != 2 {
				t.Errorf("GetGroupMembers() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestJoinGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user_center.JoinGroupRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.JoinGroupResponse
		wantErr  bool
	}{
		{
			name:    "Duplicate key 1-1",
			args:    args{ctx: ctx, req: &user_center.JoinGroupRequest{UserId: 2, GroupId: 1}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := JoinGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("JoinGroup() = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}

func TestExitGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user_center.ExitGroupRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.ExitGroupResponse
		wantErr  bool
	}{
		{
			name:    "Duplicate key 1-1",
			args:    args{ctx: ctx, req: &user_center.ExitGroupRequest{UserId: 1, GroupId: 1}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ExitGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExitGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("ExitGroup() = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}
