package services

import (
	"RestApiWithRedisAndRabbitMQ/config"
	"RestApiWithRedisAndRabbitMQ/models"
	"context"
	"encoding/json"
	"log"
)

var ctx = context.Background()

func SetRedis(user models.User) int64{
	redis:= config.InitRedis() // create the postgres db connection
	var status int64  = 1
	json, _ := json.Marshal(user)
	err := redis.Set(ctx, user.IIN, string(json), 0).Err()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return 0
	}
	return status
}


func GetRedis(key string) (models.User, int64){
	redis:= config.InitRedis() // create the postgres db connection
	var user models.User
	val, err := redis.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return user, 0
	}
	json.Unmarshal([]byte(val), &user)
	return user,1
}

func DeleteRedis(key string) (int64){
	redis:= config.InitRedis() // create the postgres db connection
	val, err := redis.Del(ctx, key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return 0
	}
	return val
}
