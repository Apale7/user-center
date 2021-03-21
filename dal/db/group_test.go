package db

import (
	"context"
	"reflect"
	"testing"
	"user_center/dal/db/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ctx = context.Background()

func TestJoinGroup(t *testing.T) {
	err := JoinGroup(ctx, model.UserGroup{UserID: 1, GroupID: 3})
	if err != nil {
		t.FailNow()
	}
}

func TestGetGroupMembers(t *testing.T) {
	resp, err := GetGroupMembers(ctx, 1)
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	logrus.Println(resp)
}

func TestGetGroup(t *testing.T) {
	g, err := GetGroup(ctx, model.Group{Model: gorm.Model{ID: 2}}, 1, true)
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	logrus.Infof("%+v\n", g)
}

func TestCreateGroup(t *testing.T) {
	type args struct {
		ctx       context.Context
		groupInfo model.Group
	}
	tests := []struct {
		name      string
		args      args
		wantGroup model.Group
		wantErr   bool
	}{
		{
			name: "create group 738",
			args: args{
				ctx: ctx,
				groupInfo: model.Group{
					OwnerID:   1,
					GroupName: "555",
				},
			},
			wantGroup: model.Group{
				Model: gorm.Model{
					ID: 4,
				},
				OwnerID:   1,
				GroupName: "555",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroup, err := CreateGroup(tt.args.ctx, tt.args.groupInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroup.ID, tt.wantGroup.ID) {
				t.Errorf("CreateGroup() = %v, want %v", gotGroup, tt.wantGroup)
			} else {
				logrus.Infof("got: %+v", gotGroup)
			}
		})
	}
}

func TestExitGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		ug  model.UserGroup
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete id: 3",
			args: args{
				ctx: ctx,
				ug:  model.UserGroup{UserID: 1, GroupID: 3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExitGroup(tt.args.ctx, tt.args.ug); (err != nil) != tt.wantErr {
				t.Errorf("ExitGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
