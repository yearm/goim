package service

import (
	"context"
	"goim/pkg/proto/logic"
	"goim/pkg/proto/ws_conn"
)

func (s *Service) DeliverMessage(ctx context.Context, in *pb_ws_conn.DeliverMessageReq) {
	client := s.loadWsClient(in.Token)
	if client == nil {
		// 设备下线
		s.logicRpc.Offline(ctx, &pb_logic.OfflineReq{Token: in.Token})
		return
	}
	client.PushMessage(in.Message)
}
