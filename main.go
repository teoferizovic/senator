package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/config"
	"github.com/teoferizovic/senator/routes"
)

func init() {

	user := config.GetEnvData("USER")
	fmt.Println(user)
	//fmt.Println(pass)
}

func main() {

	router := gin.Default()

	routes.Routes(router)

	router.Run(":8066")
}



