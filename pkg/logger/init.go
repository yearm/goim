package logger

import (
	"github.com/sirupsen/logrus"
	"goim/pkg/util/helper"
)

func Init(serverName string, level logrus.Level, hookLevel logrus.Level, elasticHosts []string, index string) {
	var hook logrus.Hook
	if hookLevel != 0 && len(elasticHosts) > 0 && index != "" {
		hosts, err := helper.GetIntranetIp()
		if err != nil {
			panic(err)
		}
		hook = NewElasticHook(&ElasticHookConfig{
			ElasticAddr: elasticHosts,
			ServerName:  serverName,
			Host:        hosts,
			Index:       index,
			HookLevel:   hookLevel,
		})
	}
	NewLogger(&Config{
		Level: level,
		Hook:  hook,
	})
}
