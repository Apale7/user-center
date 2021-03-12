package db

import (
	"context"
	"testing"
	"user_center/dal/db/model"

	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func TestJoinGroup(t *testing.T) {
	err := JoinGroup(ctx, model.UserGroup{UserID: 1, GroupID: 1})
	if err == nil {
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
	g, err := GetGroup(ctx, model.Group{GroupName: "te"})
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	logrus.Infof("%+v\n", g)
}
