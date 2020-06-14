package main

import (
	"github.com/sirupsen/logrus"
	"goim/api/ws_conn"
	"goim/conf"
	"goim/pkg/logger"
	"goim/pkg/proto/ws_conn"
)

func main() {
	conf.Init()
	logger.Init(
		pb_ws_conn.ServerName,
		logrus.Level(conf.Conf.Logger.LogLevel),
		logrus.Level(conf.Conf.Logger.HookLevel),
		conf.Conf.Logger.ElasticHosts,
		conf.Conf.Logger.Index,
	)
	ws_conn.Init(conf.Conf)
}
