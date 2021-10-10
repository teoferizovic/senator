package database

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// DBCon is the connection handle, for the mysql database, and for Redis
var (
	DBCon *gorm.DB
	Redis *redis.Client
)