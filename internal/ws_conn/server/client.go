package server

import (
	"github.com/gorilla/websocket"
	"goim/pkg/util/stringi"
)

type WsClient struct {
	token     string // 唯一标识
	agent     string
	broadcast chan []byte
	conn      *websocket.Conn
}

func NewWsClient(conn *websocket.Conn, token string, agent string) *WsClient {
	return &WsClient{
		token:     token,
		agent:     agent,
		broadcast: make(chan []byte, 256),
		conn:      conn,
	}
}

func (w *WsClient) PushMessage(msg []byte) {
	w.broadcast <- msg
}

func (w *WsClient) HandlePackage(msg []byte) {
	if len(msg) == 1 && stringi.Bytes2str(msg) == "0" {
		w.PushMessage(msg)
	}
}
