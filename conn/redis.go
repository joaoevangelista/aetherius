package conn

import (
	"encoding/json"
	"os"
	"time"

	"gopkg.in/go-redis/cache.v1"
	"gopkg.in/redis.v3"
)

func InitRedis() *cache.Codec {

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			os.Getenv("REDIS_HOST"): os.Getenv("REDIS_PORT"),
		},
		Password:     os.Getenv("REDIS_PASSWORD"),
		DialTimeout:  3 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})

	return &cache.Codec{
		Ring: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}
}
