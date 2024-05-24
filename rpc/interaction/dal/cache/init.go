package cache

import (
	"bocchi/config"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var (
	rLike    *redis.Client
	rComment *redis.Client
	//rFollow  *redis.Client
	//rMessage *redis.Client
)

func Init() {
	rLike = redis.NewClient(&redis.Options{
		Addr:       config.Redis.Addr,
		ClientName: "Like",
		DB:         0,
	})
	rComment = redis.NewClient(&redis.Options{
		Addr:       config.Redis.Addr,
		ClientName: "Comment",
		DB:         1,
	})
}

func i64ToStr(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
