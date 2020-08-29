// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redigor

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 策略模式

// ModeInterface
type ModeInterface interface {
	fmt.Stringer
	GetConn() redis.Conn
	NewConn() (redis.Conn, error)
}

// DefaultDialOpts 默认连接配置
func DefaultDialOpts() []redis.DialOption {
	return []redis.DialOption{
		redis.DialConnectTimeout(time.Second),
		redis.DialReadTimeout(time.Second * 3),
		redis.DialWriteTimeout(time.Second * 3),
	}
}

// DefaultPoolOpts 默认连接池配置
func DefaultPoolOpts() []PoolOption {
	return []PoolOption{
		Wait(false),
		MaxIdle(2 * runtime.GOMAXPROCS(0)),
		IdleTimeout(time.Second * 15),
		TestOnBorrow(func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		}),
	}
}
