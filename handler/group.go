package handler

import (
	"context"
	"user_center/dal/db"
	"user_center/dal/db/model"
	"user_center/dto"
	user_center "user_center/proto/user-center"

	"github.com/Apale7/common/constdef"
)

func CreateGroup(ctx context.Context, req *user_center.CreateGroupRequest) (resp *user_center.CreateGroupResponse, err error) {
	if !checkGroupInfo(req) {
		return nil, constdef.ErrInvalidParams
	}
	resp = &user_center.CreateGroupResponse{}
	group, err := db.CreateGroup(ctx, model.Group{OwnerID: req.OwnerId, GroupName: req.GetGroupName()})
	if err != nil {
		return
	}
	resp.Group = dto.ToGRPCGroup(&group)
	return
}

func checkGroupInfo(req *user_center.CreateGroupRequest) bool {
	if req == nil {
		return false
	}
	if req.GroupName == "" {
		return false
	}
	if req.OwnerId <= 0 {
		return false
	}
	return true
}

func GetGroup(ctx context.Context, req *user_center.GetGroupRequest) (resp *user_center.GetGroupResponse, err error) {
	resp = &user_center.GetGroupResponse{}
	groups, err := db.GetGroup(ctx, *dto.ToModelGroup(req.GroupInfo))
	if err != nil {
		return
	}
	resp.Groups = make([]*user_center.Group, 0, len(groups))
	for _, g := range groups {
		resp.Groups = append(resp.Groups, dto.ToGRPCGroup(&g))
	}
	return
}
func GetGroupMembers(ctx context.Context, req *user_center.GetGroupMembersRequest) (resp *user_center.GetGroupMembersResponse, err error) {
	resp = &user_center.GetGroupMembersResponse{}
	members, err := db.GetGroupMembers(ctx, uint(req.GroupId))
	if err != nil {
		return
	}
	resp.Members = make([]*user_center.UserExtra, 0, len(members))
	for _, m := range members {
		resp.Members = append(resp.Members, dto.ModelUserExtra2RPCUserExtra(&m))
	}
	return
}
func JoinGroup(ctx context.Context, req *user_center.JoinGroupRequest) (resp *user_center.JoinGroupResponse, err error) {
	err = db.JoinGroup(ctx, model.UserGroup{UserID: req.UserId, GroupID: req.GroupId})
	return
}
func ExitGroup(ctx context.Context, req *user_center.ExitGroupRequest) (resp *user_center.ExitGroupResponse, err error) {
	err = db.ExitGroup(ctx, model.UserGroup{UserID: req.UserId, GroupID: req.GroupId})
	return
}
