package service

import (
	"context"
	"goim/pkg/proto/ws_conn"
)

// 投递消息
func (s *Service) DeliverMessage(ctx context.Context, in *pb_ws_conn.DeliverMessageReq) (*pb_ws_conn.DeliverMessageResp, error) {
	return &pb_ws_conn.DeliverMessageResp{}, nil
}
