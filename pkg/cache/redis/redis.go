package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	Addr        string        `mapstructure:"addr"`
	Pwd         string        `mapstructure:"pwd"`
	MaxRetries  int           `mapstructure:"maxRetries"`
	PoolSize    int           `mapstructure:"poolSize"`
	IdleTimeout time.Duration `mapstructure:"idleTimeout"`
	Db          int           `mapstructure:"db"`
}

func NewRedis(c *Config) (rc *redis.Client) {
	rc = redis.NewClient(&redis.Options{
		Network:     "tcp",
		Addr:        c.Addr,
		Password:    c.Pwd,
		DB:          c.Db,
		MaxRetries:  c.MaxRetries,
		PoolSize:    c.PoolSize,
		IdleTimeout: time.Duration(c.IdleTimeout),
	})
	_, err := rc.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("connnect redis error(%v)", err))
	}
	return
}
