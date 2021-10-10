package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/controller"
	"github.com/teoferizovic/senator/middleware"
)

func Routes(router *gin.Engine){

	user := router.Group("/user")
	{
		user.POST("/register", controller.UserRegister)
		user.POST("/login", controller.UserLogin)
		user.PUT("/logout", controller.UserLogout)
		user.GET("/index",middleware.AuthMiddleware(), controller.UserIndex)
	}

	post := router.Group("/post", middleware.AuthMiddleware())
	{
		post.GET("/index",controller.UserIndex)
	}

}

