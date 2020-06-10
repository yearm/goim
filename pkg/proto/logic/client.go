package pb_logic

import "goim/pkg/net/rpc"

func NewClient(c *rpc.ClientConfig) LogicClient {
	conn := rpc.GetClientConn(c)
	return NewLogicClient(conn)
}
