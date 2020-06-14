package server

import (
	"context"
	"goim/pkg/proto/logic"
	"goim/pkg/proto/ws_conn"
)

// 投递消息
func (s *Server) DeliverMessage(ctx context.Context, in *pb_ws_conn.DeliverMessageReq) (*pb_ws_conn.DeliverMessageResp, error) {
	client := s.loadWsClient(in.Token)
	if client == nil {
		// 设备下线
		s.logicRpc.Offline(ctx, &pb_logic.OfflineReq{Token: in.Token})
		return &pb_ws_conn.DeliverMessageResp{}, nil
	}
	client.PushMessage(in.Message)
	return &pb_ws_conn.DeliverMessageResp{}, nil
}
