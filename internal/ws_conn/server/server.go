package server

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.etcd.io/etcd/clientv3"
	"goim/conf"
	"goim/pkg/logger"
	"goim/pkg/net/rpc"
	"goim/pkg/net/rpc/balancer"
	"goim/pkg/proto/logic"
	"goim/pkg/util/cmap"
	"goim/pkg/util/helper"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var intranetIp, _ = helper.GetIntranetIp()

type Server struct {
	c        *conf.Config
	manager  cmap.ConcurrentMap
	logicRpc pb_logic.LogicClient
}

func New(c *conf.Config) (s *Server) {
	return &Server{
		c:       c,
		manager: cmap.New(),
		logicRpc: pb_logic.NewClient(&rpc.ClientConfig{
			EtcdConfig: clientv3.Config{
				Endpoints: c.Etcd.Hosts,
			},
			RegistryDir: c.Etcd.Dir,
			ServiceName: pb_logic.ServerName,
			BalanceName: balancer.RoundRobin,
		}),
	}
}

func (s *Server) ReadPump(c *WsClient) {
	defer func() {
		// 设备离线
		s.logicRpc.Offline(context.TODO(), &pb_logic.OfflineReq{
			Token: c.token,
		})
		s.deleteWsClient(c.token)
		close(c.broadcast)
	}()

	// 设备登录
	_, err := s.logicRpc.SignIn(context.TODO(), &pb_logic.SignInReq{
		Token:  c.token,
		Agent:  c.agent,
		WsAddr: fmt.Sprintf("%s:%d", intranetIp, s.c.LogicAddrs.RpcPort),
	})
	if err != nil {
		return
	}

	s.storeWsClient(c.token, c)
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Logger.Error("readPump ReadMessage err:", err)
			}
			return
		}
		if message == nil {
			return
		}
		// 主要处理心跳包，不通过websocket发消息
		c.HandlePackage(message)
	}
}

func (s *Server) WritePump(c *WsClient) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.broadcast:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			logger.Logger.Infof("message write body:%s", message)
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 存储客户端连接
func (s *Server) storeWsClient(token string, wsConn *WsClient) {
	s.manager.Set(token, wsConn)
}

// 获取客户端连接
func (s *Server) loadWsClient(token string) *WsClient {
	value, ok := s.manager.Get(token)
	if ok {
		return value.(*WsClient)
	}
	return nil
}

// 删除客户端连接
func (s *Server) deleteWsClient(token string) {
	s.manager.Remove(token)
}
