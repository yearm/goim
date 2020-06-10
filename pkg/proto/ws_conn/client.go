package pb_ws_conn

import "goim/pkg/net/rpc"

func NewClient(c *rpc.ClientConfig) WsConnClient {
	conn := rpc.GetClientConn(c)
	return NewWsConnClient(conn)
}
