// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cluster

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/hollson/redigor"
)

type options struct {
	nodes    []string // e.g. 127.0.0.1:30001
	waitTime time.Duration
	poolOpts []redigor.PoolOption
	dialOpts []redis.DialOption
}

type Option func(opts *options)

func WithNodes(value []string) Option {
	return func(opts *options) {
		opts.nodes = value
	}
}

func WithWait(value time.Duration) Option {
	return func(opts *options) {
		opts.waitTime = value
	}
}

func WithPool(value ...redigor.PoolOption) Option {
	return func(opts *options) {
		for _, poolOpt := range value {
			opts.poolOpts = append(opts.poolOpts, poolOpt)
		}
	}
}

func WithDial(value ...redis.DialOption) Option {
	return func(opts *options) {
		for _, dialOpt := range value {
			opts.dialOpts = append(opts.dialOpts, dialOpt)
		}
	}
}
