package main

import (
	"github.com/hollson/redigor/alone"
)

func main() {
	cl := alone.NewClient(alone.Addr(":6379"))
	cl.IntSet("11",11)


	// cl := alone.NewClient(alone.Addr(":6379"))
	// rr, err := cl.String(func(c redis.Conn) (res interface{}, err error) {
	// 	return c.Do("get", "hello")
	// })
	// fmt.Println(rr, err)
	//
	// for {
	// 	// var cli = mode.New()
	// 	var cli = sentinel.NewClient(sentinel.Addrs([]string{":26379", ":26380", ":26381"}))
	//
	// 	result, err := cli.String(func(c redis.Conn) (res interface{}, err error) {
	// 		return c.Do("GET", "hello")
	// 	})
	//
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	log.Println(result)
	//
	// 	time.Sleep(time.Second * 5)
	// }
}
