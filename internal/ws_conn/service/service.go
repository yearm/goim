package service

import (
	"github.com/gorilla/websocket"
	"goim/conf"
)

type Service struct {
	// rpc
}

func New(c *conf.Config) (s *Service) {
	return &Service{
	}
}

type WsClient struct {
	Uuid string // 唯一标识
	Conn *websocket.Conn
}

func NewWsClient(conn *websocket.Conn, uuid string) *WsClient {
	return &WsClient{
		Uuid: uuid,
		Conn: conn,
	}
}
