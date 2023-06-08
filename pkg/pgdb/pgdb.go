package pgdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	defaultKey = "default"
	driverName = "pgx"
)

var _dbPool = NewConnPool(context.Background())

func Pool() *ConnPool {
	return _dbPool
}

type ConnPool struct {
	mu  sync.Mutex
	ctx context.Context

	pool map[string]*sqlx.DB
}

func NewConnPool(ctx context.Context) *ConnPool {
	return &ConnPool{
		ctx:  ctx,
		pool: map[string]*sqlx.DB{},
	}
}

func (c *ConnPool) AddConn(key, dsn string) (conn *sqlx.DB, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		key = defaultKey
	}

	if c.pool[key] != nil {
		conn = c.pool[key]
		return
	}

	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return
	}

	if conn, err = sqlx.ConnectContext(c.ctx, driverName, stdlib.RegisterConnConfig(connConfig)); err != nil {
		return
	}

	c.pool[key] = conn
	return
}

func (c *ConnPool) GetConn(key string) (conn *sqlx.DB, err error) {
	if key == "" {
		key = defaultKey
	}

	if c.pool[key] == nil {
		err = fmt.Errorf("conn with key=%s not exists", key)
		return
	}

	conn = c.pool[key]
	return
}

func (c *ConnPool) CloseAll() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, v := range c.pool {
		_ = v.Close()
	}
}
