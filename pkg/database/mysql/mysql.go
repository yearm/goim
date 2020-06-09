package mysql

import (
	"fmt"
	"time"
)

type Config struct {
	Addr        string        `mapstructure:"addr"`
	DSN         string        `mapstructure:"dsn"`
	ReadDSN     []string      `mapstructure:"readDsn"`
	Active      int           `mapstructure:"active"`      // 最大连接数
	Idle        int           `mapstructure:"idle"`        // 最大空闲连接
	IdleTimeout time.Duration `mapstructure:"idleTimeout"` // 最大的连接生存期
}

func NewMySQL(c *Config) (db *DB) {
	db, err := Open(c)
	if err != nil {
		panic(fmt.Errorf("open mysql error(%v)", err))
	}
	return
}
