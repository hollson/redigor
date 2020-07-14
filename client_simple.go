// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// 一些简单的扩展封装
package redigor

import (
	"github.com/gomodule/redigo/redis"
)

func (c *Client) IntSet(key string, val int) (interface{}, error) {
	return c.Execute(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("SET", key, val)
	})
}

func (c *Client) IntGet(key string) (int, error) {
	return redis.Int(c.Execute(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("GET", key)
	}))
}

func (c *Client) StringSet(key, val string) (interface{}, error) {
	return c.Do("SET", key, val)

	// return c.Execute(func(c redis.Conn) (res interface{}, err error) {
	// 	return c.Do("SET", key, val)
	// })
}

func (c *Client) StringGet(key string) (string, error) {
	return redis.String(c.Do("GET", key))
	// return redis.Int(c.Execute(func(c redis.Conn) (res interface{}, err error) {
	// 	return c.Do("GET",key)
	// }))
}

// func (c *Client) IntsGet(keys ...string) ([]int, error) {
// 	return redis.Ints(c.Execute(func(c redis.Conn) (res interface{}, err error) {
// 		return c.Do("MGET",keys)
// 	}))
// }
