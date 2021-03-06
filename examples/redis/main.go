package main

import (
	"strings"

	"github.com/reugn/go-streams"

	"github.com/go-redis/redis"
	ext "github.com/reugn/go-streams/extension"
	"github.com/reugn/go-streams/flow"
)

//docker exec -it pubsub bash
//https://redis.io/topics/pubsub
func main() {
	config := &redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	}
	source, err := ext.NewRedisSource(config, "test")
	streams.Check(err)
	flow1 := flow.NewMap(toUpper, 1)
	sink := ext.NewRedisSink(config, "test2")

	source.Via(flow1).To(sink)
}

var toUpper = func(in interface{}) interface{} {
	msg := in.(*redis.Message)
	return strings.ToUpper(msg.Payload)
}
