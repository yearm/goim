package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goim/pkg/cache/redis"
	"goim/pkg/database/etcd"
	"goim/pkg/database/mysql"
	"os"
)

var (
	Conf = &Config{}
)

type Config struct {
	Mysql       *mysql.Config `mapstructure:"mysql"`
	Redis       *redis.Config `mapstructure:"redis"`
	Etcd        *etcd.Config  `mapstructure:"etcd"`
	LogicAddrs  *Addrs        `mapstructure:"logic"`
	WsConnAddrs *Addrs        `mapstructure:"wsConn"`
	Logger      *LoggerConfig `mapstructure:"logger"`
}

func Init() {
	// 获取项目的执行路径
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath(fmt.Sprintf("%s/%s", dir, "conf")) //设置配置文件的搜索目录
	read(v)

	// 监听配置变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		// 重新加载配置
		read(v)
	})
}

// 加载配置
func read(v *viper.Viper) {
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("conf.Init() error(%v)", err))
	}
}
