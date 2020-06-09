package main

import (
	"goim/api/logic"
	"goim/conf"
)

func main() {
	conf.Init()
	logic.Init(conf.Conf)
}
