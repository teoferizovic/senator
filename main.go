package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/teoferizovic/senator/database"
	"github.com/teoferizovic/senator/routes"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func init() {

	var err error

	//mysql connection
	dsn := "root:root@tcp(127.0.0.1:3306)/proba?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//redis connection
	database.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "", // no password set
		DB:       7,  // use default DB
	})

}

func main() {

	router := gin.Default()

	routes.Routes(router)

	router.Run(":8066")
}



