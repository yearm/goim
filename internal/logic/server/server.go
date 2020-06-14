package server

import (
	"go.etcd.io/etcd/clientv3"
	"goim/conf"
	"goim/internal/logic/dao"
	"goim/pkg/net/rpc"
	"goim/pkg/net/rpc/balancer"
	"goim/pkg/proto/ws_conn"
)

type Server struct {
	c         *conf.Config
	dao       *dao.Dao
	wsConnRpc pb_ws_conn.WsConnClient
}

func New(c *conf.Config) (s *Server) {
	return &Server{
		c:   c,
		dao: dao.New(c),
		wsConnRpc: pb_ws_conn.NewClient(&rpc.ClientConfig{
			EtcdConfig: clientv3.Config{
				Endpoints: c.Etcd.Hosts,
			},
			RegistryDir: c.Etcd.Dir,
			ServiceName: pb_ws_conn.ServerName,
			BalanceName: balancer.AddrPicker,
		}),
	}
}
