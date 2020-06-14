package service

import (
	"context"
	"goim/pkg/proto/logic"
)

func (s *Service) SignIn(ctx context.Context, in *pb_logic.SignInReq) {
	// TODO do something get uid
	var uid string
	s.dao.SetDevice(uid, in)
}

func (s *Service) Offline(ctx context.Context, in *pb_logic.OfflineReq) {
	// TODO do something get uid
	var uid string
	s.dao.DelDevice(uid, in)
}

func (s *Service) SendMessage(ctx context.Context, in *pb_logic.SendMessageReq) {

}
