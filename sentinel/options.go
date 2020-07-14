// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sentinel

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hollson/redigor"
)

type options struct {
	addrs            []string
	masterName       string
	poolOpts         []redigor.PoolOption
	dialOpts         []redis.DialOption
	sentinelDialOpts []redis.DialOption
}

type Option func(opts *options)

func WithAddress(value []string) Option {
	return func(opts *options) {
		opts.addrs = value
	}
}

func WithMasterName(value string) Option {
	return func(opts *options) {
		opts.masterName = value
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

func WithSentinelDial(value ...redis.DialOption) Option {
	return func(opts *options) {
		for _, dialOpt := range value {
			opts.sentinelDialOpts = append(opts.sentinelDialOpts, dialOpt)
		}
	}
}
