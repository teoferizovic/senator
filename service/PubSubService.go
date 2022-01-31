package service

import (
	"github.com/go-redis/redis"
	"github.com/teoferizovic/senator/database"
)

func PublishData(data string) error {

	if err := database.Redis.Publish("newUser", data).Err(); err != nil {
		return err
	}

	return nil
}

func SubsribeData() *redis.PubSub {

	subscriber := database.Redis.Subscribe("newUser")

	return subscriber
}

