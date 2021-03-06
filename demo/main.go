// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/hollson/redigor/alone"
	"github.com/hollson/redigor/cluster"
	"github.com/hollson/redigor/sentinel"
)

func main() {
	// aloneDemo()
	sentinelDemo()
	// sentinelErr()
}

func aloneDemo() {
	cli := alone.NewClient(alone.WithAddress(":6379"))
	log.Println(cli.Do("PING"))
}

// DefaultDialOpts 默认连接配置
func DefaultDialOpts() []redis.DialOption {
	return []redis.DialOption{
		redis.DialPassword("123456"),
		redis.DialConnectTimeout(time.Second),
		redis.DialReadTimeout(time.Second * 3),
		redis.DialWriteTimeout(time.Second * 3),
	}
}

func sentinelDemo() {
	// e.g. 127.0.0.1:26379
	var sentinels = []string{":26379", ":26380", ":26381"}
	// 默认MasterName为mymaster

	sentinel.WithDial()
	sentinel.WithSentinelDial()
	cli := sentinel.NewClient(sentinel.WithAddress(sentinels))

	// // Do 操作
	// log.Println(cli.Do("SET", "hello", "world"))
	// log.Println(redis.String(cli.Do("GET", "hello")))
	// log.Println(cli.Do("PING"))

	_,err:=cli.Execute(func(c redis.Conn) (res interface{}, err error) {
		c.Send("SET","XXX","222222")
		// c.Flush()
		return nil,c.Flush()
		// return c.Do("SET","XYZ","123")
	})
if err!=nil{
	fmt.Println(err)
}

	// Execute
	log.Println(cli.Execute(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("PING")
	}))

	// 简单封装
	log.Println(cli.StringGet("hello"))
}

// 哨兵故障转移
func sentinelErr() {
	var sentinels = []string{":12345", ":26379", ":26380"}
	cli := sentinel.NewClient(sentinel.WithAddress(sentinels), sentinel.WithMasterName("mymaster"))

	// 先循环GetConn，否则再NewConn
	fmt.Println(cli.Do("PING")) // err + pong
	fmt.Println(cli.Do("PING")) // pong
	fmt.Println(cli.Do("PING")) // pong

	fmt.Println()
	sentinels = []string{":26379", ":26380"}
	cli = sentinel.NewClient(sentinel.WithAddress(sentinels), sentinel.WithMasterName("unknown"))
	fmt.Println(cli.Do("PING"))
}

// 集群模式(待验证)
func clusterDemo() {
	cli := cluster.NewClient(cluster.WithNodes([]string{}))
	_ = cli
}
