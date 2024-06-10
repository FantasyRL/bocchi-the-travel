package cache

import (
	"bocchi/config"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var (
	rFollow *redis.Client
)

func Init() {
	rFollow = redis.NewClient(&redis.Options{
		Addr:       config.Redis.Addr,
		ClientName: "Follow",
		DB:         2,
	})
}

func i64ToStr(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
