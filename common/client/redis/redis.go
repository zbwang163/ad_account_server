package redis

import "github.com/go-redis/redis"

var (
	Redis       = make(map[string]*redis.Client, 2)
	redisConfig = map[string]int{
		"ad.info.account_server": 1,
		"ad.info.content_server": 2,
	}
)

func InitRedis(psm string) {
	db, ok := redisConfig[psm]
	if !ok {
		return
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       db, // use default DB
	})
	Redis[psm] = client
}
