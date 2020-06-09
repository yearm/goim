package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync/atomic"
	"time"
)

type DB struct {
	write *conn
	read  []*conn
	idx   int64
}

type conn struct {
	*sql.DB
	conf *Config
}

type Tx struct {
	db     *conn
	tx     *sql.Tx
	c      context.Context
	cancel func()
}

func Open(c *Config) (*DB, error) {
	d, err := connect(c, c.DSN)
	if err != nil {
		return nil, err
	}
	w := &conn{DB: d, conf: c}
	rs := make([]*conn, 0, len(c.ReadDSN))
	for _, rd := range c.ReadDSN {
		d, err := connect(c, rd)
		if err != nil {
			return nil, err
		}
		r := &conn{DB: d, conf: c}
		rs = append(rs, r)
	}
	return &DB{write: w, read: rs}, nil
}

func connect(c *Config, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(c.Active)
	d.SetMaxIdleConns(c.Idle)
	d.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return d, nil
}

func (d *DB) readIndex() int {
	if len(d.read) == 0 {
		return 0
	}
	v := atomic.AddInt64(&d.idx, 1)
	return int(v) % len(d.read)
}

func (d *DB) QueryRow() {

}

func (d *DB) Query() {

}

func (d *DB) Exec() {

}

func (d *DB) Prepare() {

}
