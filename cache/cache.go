package cache

import (
	"encoding/json"

	"github.com/torkashvand/goshortener/config"

	"github.com/go-redis/redis"
	"github.com/torkashvand/goshortener/log"
)

var redisClient *redis.Client

// InitializeRedis connect to redis
func InitializeRedis() {
	config := config.Config()

	redisClient = redis.NewClient(&redis.Options{
		Addr:       config.GetString("REDIS_ADDRESS"),
		PoolSize:   config.GetInt("REDIS_POOL_SIZE"),
		MaxRetries: config.GetInt("REDIS_MAX_RETRIES"),
		Password:   config.GetString("REDIS_PASSWORD"),
		DB:         config.GetInt("REDIS_DB_NUMBER"),
	})

	if ping, err := redisClient.Ping().Result(); err == nil && len(ping) > 0 {
		log.Info("Connected to Redis")
	} else {
		log.Errorln("Redis Connection Failed")
	}
}

//GetValue is a wrapper for get command
func GetValue(key string) (interface{}, error) {
	var deserializedValue interface{}
	serializedValue, err := redisClient.Get(key).Result()
	json.Unmarshal([]byte(serializedValue), &deserializedValue)

	log.Errorln(err)

	return deserializedValue, err
}

//SetValue is a wrapper for set commnad
func SetValue(key string, value interface{}) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := redisClient.Set(key, string(serializedValue), 0).Err()

	log.Errorln(err)

	return true, err
}
