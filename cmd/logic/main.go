package main

import (
	"github.com/sirupsen/logrus"
	"goim/api/logic"
	"goim/conf"
	"goim/pkg/logger"
	"goim/pkg/proto/logic"
)

func main() {
	conf.Init()
	logger.Init(
		pb_logic.ServerName,
		logrus.Level(conf.Conf.Logger.LogLevel),
		logrus.Level(conf.Conf.Logger.HookLevel),
		conf.Conf.Logger.ElasticHosts,
		conf.Conf.Logger.Index,
	)
	logic.Init(conf.Conf)
}
