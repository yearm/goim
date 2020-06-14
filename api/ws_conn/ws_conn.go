package ws_conn

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goim/internal/ws_conn/service"
	"goim/pkg/logger"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type connWsReq struct {
	Token string
	Agent string
}

func connWs(ctx *gin.Context) {
	in := new(connWsReq)
	if err := ctx.Bind(in); err != nil {
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	client := service.NewWsClient(conn, in.Token, in.Agent)
	go srv.ReadPump(client)
	go srv.WritePump(client)
}
