package redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	redisClient *redis.Client
)

func init() {
	if err := initClient(); err != nil {
		panic(err)
	}
	log.Info("连接Redis成功")
}

func initClient() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
