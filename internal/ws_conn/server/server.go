package server

import (
	"goim/internal/ws_conn/service"
	"goim/pkg/proto/ws_conn"
	"google.golang.org/grpc"
)

type server struct {
	srv *service.Service
}

func RegisterServer(s *grpc.Server) {
	pb_ws_conn.RegisterWsConnServer(s, &server{})
}
