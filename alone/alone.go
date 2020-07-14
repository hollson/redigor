// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alone

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hollson/redigor"
)

type aloneMode struct{ pool redis.Pool }

func (am *aloneMode) GetConn() redis.Conn {
	return am.pool.Get()
}

func (am *aloneMode) NewConn() (redis.Conn, error) {
	return am.pool.Dial()
}

func (am *aloneMode) String() string { return "alone" }

func New(optFuncs ...Option) redigor.ModeInterface {
	opts := options{
		addr:     "127.0.0.1:6381",
		dialOpts: redigor.DefaultDialOpts(),
		poolOpts: redigor.DefaultPoolOpts(),
	}
	for _, optFunc := range optFuncs {
		optFunc(&opts)
	}
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", opts.addr, opts.dialOpts...)
		},
	}
	for _, poolOptFunc := range opts.poolOpts {
		poolOptFunc(&pool)
	}
	return &aloneMode{pool: pool}
}

func NewClient(optFuncs ...Option) *redigor.Client {
	return redigor.New(New(optFuncs...))
}
