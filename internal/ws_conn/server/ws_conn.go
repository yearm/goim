package server

import (
	"context"
	"goim/pkg/proto/ws_conn"
)

// 投递消息
func (s *server) DeliverMessage(ctx context.Context, in *pb_ws_conn.DeliverMessageReq) (*pb_ws_conn.DeliverMessageResp, error) {
	s.srv.DeliverMessage(ctx, in)
	return &pb_ws_conn.DeliverMessageResp{}, nil
}
