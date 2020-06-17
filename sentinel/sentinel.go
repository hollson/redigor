package sentinel

import (
	"runtime"
	"time"
	"github.com/hollson/redigor/fzambia"
	"github.com/gomodule/redigo/redis"
	"github.com/hollson/redigor"
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

func (sm *sentinelMode) String() string { return "sentinel" }

func New(optFuncs ...OptFunc) redigor.ModeInterface {
	opts := options{
		addrs: []string{
			"127.0.0.1:26379",
		},
		masterName: "mymaster",
		poolOpts:   redigor.DefaultPoolOpts(),
		dialOpts:   redigor.DefaultDialOpts(),
	}
	for _, optFunc := range optFuncs {
		optFunc(&opts)
	}
	if len(opts.sentinelDialOpts) == 0 {
		opts.sentinelDialOpts = opts.dialOpts
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

func NewClient(optFuncs ...OptFunc) *redigor.Client {
	return redigor.New(New(optFuncs...))
}
