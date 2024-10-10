package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/rubberduckkk/credit-card/internal/app"
	"github.com/rubberduckkk/credit-card/internal/infra/config"
)

func main() {
	config.Load()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Get().Port))
	if err != nil {
		panic(err)
	}

	err = app.RunGRPCServer(listener)
	if err != nil {
		logrus.Errorf("run grpc server failed: %v", err)
	}
}
