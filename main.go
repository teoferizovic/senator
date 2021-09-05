package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/database"
	"github.com/teoferizovic/senator/routes"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func init() {

	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/proba?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func main() {

	router := gin.Default()

	routes.Routes(router)

	router.Run(":8066")
}



