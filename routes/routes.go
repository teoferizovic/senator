package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/controller"
	"github.com/teoferizovic/senator/middleware"
)

func Routes(router *gin.Engine){

	//routes for users
	user := router.Group("/user")
	{
		user.POST("/register", controller.UserRegister)
		user.POST("/login", controller.UserLogin)
		user.PUT("/logout", controller.UserLogout)
		user.GET("/test",middleware.AuthMiddleware(), controller.UserTest)
		user.GET("/index/:id", controller.UserIndex)
	}

	//routes for posts
	post := router.Group("/post", middleware.AuthMiddleware())
	{
		post.GET("/index",controller.UserTest)
	}

	//routes for articles
	article := router.Group("/article")
	{
		article.GET("/index",controller.ArticleIndex)
	}

	//routes for comments
	comment := router.Group("/comment", middleware.AuthMiddleware())
	{
		comment.GET("/index",controller.CommentIndex)
	}
}

