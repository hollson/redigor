// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redigor

import (
	"github.com/gomodule/redigo/redis"
)

func (c *Client) IntGet(val int) (int, error) {
	return c.Int(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("GET", val)
	})
}

func (c *Client) IntSet(key string,val int) (int, error) {
	return c.Int(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("SET", key,val)
	})
}

func (c *Client) IntsGet(val ...int) ([]int, error) {
	return c.Ints(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("GET", val)
	})
}

func (c *Client) IntsSet(val ...int) ([]int, error) {
	return c.Ints(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("SET", val)
	})
}



// func (c *Client) IntMap(val string) (map[string]int, error) {
// 	return redis.IntMap(c.Execute(fn))
// }
//
// func (c *Client) Int64(val string) (int64, error) {
// 	return redis.Int64(c.Execute(fn))
// }
//
// func (c *Client) Int64s(val string) ([]int64, error) {
// 	return redis.Int64s(c.Execute(fn))
// }
// func (c *Client) Int64Map(val string) (map[string]int64, error) {
// 	return redis.Int64Map(c.Execute(fn))
// }
//
// func (c *Client) Uint64(val string) (uint64, error) {
// 	return redis.Uint64(c.Execute(fn))
// }

// func (c *Client) Bool(val string) (bool, error) {
// 	return redis.Bool(c.Execute(fn))
// }
//
// func (c *Client) String(val string) (string, error) {
// 	return redis.String(c.Execute(fn))
// }
//
// func (c *Client) StringMap(val string) (map[string]string, error) {
// 	return redis.StringMap(c.Execute(fn))
// }
//
// func (c *Client) Strings(val string) ([]string, error) {
// 	return redis.Strings(c.Execute(fn))
// }
//
// func (c *Client) Bytes(val string) ([]byte, error) {
// 	return redis.Bytes(c.Execute(fn))
// }
//
// func (c *Client) ByteSlices(val string) ([][]byte, error) {
// 	return redis.ByteSlices(c.Execute(fn))
// }
//
// func (c *Client) Positions(val string) ([]*[2]float64, error) {
// 	return redis.Positions(c.Execute(fn))
// }
//
// func (c *Client) Float64(val string) (float64, error) {
// 	return redis.Float64(c.Execute(fn))
// }
//
// func (c *Client) Float64s(val string) ([]float64, error) {
// 	return redis.Float64s(c.Execute(fn))
// }
//
// func (c *Client) Values(val string) ([]interface{}, error) {
// 	return redis.Values(c.Execute(fn))
// }
