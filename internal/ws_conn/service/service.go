package service

import (
	"github.com/gorilla/websocket"
	"go.etcd.io/etcd/clientv3"
	"goim/conf"
	"goim/pkg/net/rpc"
	"goim/pkg/net/rpc/balancer"
	"goim/pkg/proto/logic"
)

type Service struct {
	logicRpc pb_logic.LogicClient
}

func New(c *conf.Config) (s *Service) {
	return &Service{
		logicRpc: pb_logic.NewClient(&rpc.ClientConfig{
			EtcdConfig: clientv3.Config{
				Endpoints: c.Etcd.Hosts,
			},
			RegistryDir: c.Etcd.Dir,
			ServiceName: "logic_rpc",
			BalanceName: balancer.RoundRobin,
		}),
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
