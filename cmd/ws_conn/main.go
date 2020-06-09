package main

import (
	"goim/api/ws_conn"
	"goim/conf"
)

func main() {
	conf.Init()
	ws_conn.Init(conf.Conf)
}
