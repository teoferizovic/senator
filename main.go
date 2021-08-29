package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/routes"
)

func main() {

	router := gin.Default()

	routes.Routes(router)

	router.Run(":8066")
}

