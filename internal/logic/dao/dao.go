package dao

import (
	"github.com/go-redis/redis"
	"goim/conf"
	rc "goim/pkg/cache/redis"
	"goim/pkg/database/mysql"
)

type Dao struct {
	c  *conf.Config
	db *mysql.DB
	rc *redis.Client
}

func New(c *conf.Config) (dao *Dao) {
	return &Dao{
		c:  c,
		db: mysql.NewMySQL(c.Mysql),
		rc: rc.NewRedis(c.Redis),
	}
}
