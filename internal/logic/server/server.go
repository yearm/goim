package server

import (
	"goim/internal/logic/service"
	"goim/pkg/proto/logic"
	"google.golang.org/grpc"
)

type server struct {
	svr *service.Service
}

func RegisterServer(s *grpc.Server) {
	pb_logic.RegisterLogicServer(s, &server{})
}
