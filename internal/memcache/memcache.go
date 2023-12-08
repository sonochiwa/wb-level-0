package memcache

import (
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var MC *cache.Cache

func init() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	MC = cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
}
