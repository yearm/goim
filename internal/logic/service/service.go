package service

import (
	"goim/conf"
	"goim/internal/logic/dao"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
	// rpc
}

func New(c *conf.Config) (s *Service) {
	return &Service{
		c:   c,
		dao: dao.New(c),
	}
}
