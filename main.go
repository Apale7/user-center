package main

import (
	"context"
	"net"

	config "user_center/config_loader"
	"user_center/dal/db"
	h "user_center/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	conn   *grpc.ClientConn
	ctx    context.Context
	cancel context.CancelFunc
	err    error
	addr   = ":9999"
)

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %+v", err)
	}
	server := grpc.NewServer()
	h.RegisterUserCenterServer(server, &UserCenterServer{})
	logrus.Infof("server listen at %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %+v", err)
	}
}

func init() {
	config.Init()
	port := config.Get("server_port")
	if port != "" {
		addr = ":" + port
	}
	db.Init()
}
