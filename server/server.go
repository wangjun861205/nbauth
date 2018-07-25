package server

import (
	"net"

	"github.com/wangjun861205/nbauth/api"
	"google.golang.org/grpc"
)

func Serve(addr, logPath, mysqlAddr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	err = api.InitManager(mysqlAddr)
	if err != nil {
		return err
	}
	s, err := api.NewServer(logPath)
	if err != nil {
		return err
	}
	rpcServer := grpc.NewServer()
	api.RegisterAuthServiceServer(rpcServer, s)
	return rpcServer.Serve(l)
}
