package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/teoferizovic/senator/config"
	"github.com/teoferizovic/senator/database"
	"github.com/teoferizovic/senator/database/migration"
	"github.com/teoferizovic/senator/routes"
	"github.com/teoferizovic/senator/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	//"github.com/sirupsen/logrus"
)

func init() {

	var err error

	//db variables
	dbHost := config.GetEnvData("DB_HOST")
	dbUser := config.GetEnvData("DB_USER")
	dbPass := config.GetEnvData("DB_PASSWORD")
	dbPort := config.GetEnvData("DB_PORT")
	dbName := config.GetEnvData("DB_NAME")

	//redis variables
	redisHost := config.GetEnvData("REDIS_HOST")
	redisPort := config.GetEnvData("REDIS_PORT")
	redisPassword := config.GetEnvData("REDIS_PASSWORD")
	redisDbNum, _ := strconv.Atoi(config.GetEnvData("REDIS_DB_NUMBER"))

	//mysql connection
	dsn := dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//redis connection
	database.Redis = redis.NewClient(&redis.Options{
		Addr:     redisHost+":"+redisPort, // host:port of the redis server
		Password: redisPassword, // no password set
		DB:       redisDbNum,
	})

	//init logger
	service.InitiLogger()

	//execute migration
	migration.Migrate()

}

func main() {

	router := gin.Default()

	routes.Routes(router)

	router.Run(":"+config.GetEnvData("APP_PORT"))

}


