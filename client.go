// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redigor

import (
	"github.com/gomodule/redigo/redis"
)

// SubscribeFunc 订阅回调函数
type SubscribeFunc func(c redis.PubSubConn) (err error)

// ExecuteFunc 普通回调函数
type ExecuteFunc func(c redis.Conn) (res interface{}, err error)

type Client struct {
	mode ModeInterface
}

func New(mode ModeInterface) *Client {
	return &Client{mode: mode}
}

// Mode 当前客户端使用模式
// alone 单机或者代理入口模式
// cluster Redis-Cluster集群模式
// sentinel Redis-Sentinel哨兵模式
func (c *Client) Mode() string {
	return c.mode.String()
}

func (c *Client) Execute(fn ExecuteFunc) (res interface{}, err error) {
	conn := c.mode.GetConn()
	defer conn.Close()
	if res, err = fn(conn); err != nil {
		if _, ok := err.(redis.Error); ok {
			return
		} else if newConn, newErr := c.mode.NewConn(); newErr != nil {
			return
		} else {
			defer newConn.Close()
			res, err = fn(newConn)
		}
	}
	return
}

func (c *Client) Do(cmd string, args ...interface{}) (res interface{}, err error) {

	conn := c.mode.GetConn()
	defer conn.Close()
	if res, err = conn.Do(cmd, args...); err != nil {
		if _, ok := err.(redis.Error); ok {
			return
		} else if newConn, newErr := c.mode.NewConn(); newErr != nil {
			return
		} else {
			defer newConn.Close()
			res, err = newConn.Do(cmd, args...)
		}
	}
	return
}

func (c *Client) Subscribe(fn SubscribeFunc) error {
	conn, err := c.mode.NewConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(redis.PubSubConn{Conn: conn})
}