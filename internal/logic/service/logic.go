package service

import (
	"context"
	"goim/pkg/proto/logic"
)

// 设备登录
func (s *Service) SignIn(ctx context.Context, in *pb_logic.SignInReq) (*pb_logic.SignInResp, error) {
	return &pb_logic.SignInResp{}, nil
}

// 设备离线
func (s *Service) Offline(ctx context.Context, in *pb_logic.OfflineReq) (*pb_logic.OfflineResp, error) {
	return &pb_logic.OfflineResp{}, nil
}

// 发送消息
func (s *Service) SendMessage(ctx context.Context, in *pb_logic.SendMessageReq) (*pb_logic.SendMessageResp, error) {
	return &pb_logic.SendMessageResp{}, nil
}
