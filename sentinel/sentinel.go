// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sentinel

import (
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/hollson/redigor"
	"github.com/hollson/redigor/sentinel/fzambia"
)

type sentinelMode struct {
	pool *redis.Pool
}

func (sm *sentinelMode) GetConn() redis.Conn {
	return sm.pool.Get()
}

func (sm *sentinelMode) NewConn() (redis.Conn, error) {
	return sm.pool.Dial()
}

// 使用方法来约束字段
func (sm *sentinelMode) String() string { return "sentinel" }

func New(opt ...Option) redigor.ModeInterface {
	opts := options{
		addrs: []string{
			"127.0.0.1:26379",
		},
		masterName: "mymaster",
		poolOpts:   redigor.DefaultPoolOpts(),
		dialOpts:   redigor.DefaultDialOpts(),
		sentinelDialOpts:redigor.DefaultDialOpts(),
	}
	for _, optFunc := range opt {
		optFunc(&opts)
	}
	st := &fzambia.Sentinel{
		Addrs:      opts.addrs,
		MasterName: opts.masterName,
		Pool: func(addr string) *redis.Pool {
			return &redis.Pool{
				Wait:    true,
				MaxIdle: runtime.GOMAXPROCS(0),
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", addr, opts.sentinelDialOpts...)
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) (err error) {
					_, err = c.Do("PING")
					return
				},
			}
		},
	}
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			addr, err := st.MasterAddr()
			if err != nil {
				return nil, err
			}
			return redis.Dial("tcp", addr, opts.dialOpts...)
		},
	}
	for _, poolOptFunc := range opts.poolOpts {
		poolOptFunc(pool)
	}
	return &sentinelMode{pool: pool}
}

func NewClient(optFunc ...Option) *redigor.Client {
	return redigor.New(New(optFunc...))
}
