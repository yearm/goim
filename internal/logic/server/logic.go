package server

import (
	"context"
	"goim/pkg/proto/logic"
)

// 设备登录
func (s *server) SignIn(ctx context.Context, in *pb_logic.SignInReq) (*pb_logic.SignInResp, error) {
	s.svr.SignIn(ctx, in)
	return &pb_logic.SignInResp{WsAddr: in.WsAddr}, nil
}

// 设备离线
func (s *server) Offline(ctx context.Context, in *pb_logic.OfflineReq) (*pb_logic.OfflineResp, error) {
	s.svr.Offline(ctx, in)
	return &pb_logic.OfflineResp{}, nil
}

// 发送消息
func (s *server) SendMessage(ctx context.Context, in *pb_logic.SendMessageReq) (*pb_logic.SendMessageResp, error) {
	s.svr.SendMessage(ctx, in)
	return &pb_logic.SendMessageResp{}, nil
}
