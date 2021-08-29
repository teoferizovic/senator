package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/controller"
)

func Routes(router *gin.Engine){

	user := router.Group("/user")
	{
		user.POST("/register", controller.UserRegister)
		user.POST("/login", controller.UserLogin)
		user.PUT("/logout", controller.UserLogout)
	}

}

