package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
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
		user.GET("/ws", controller.UserSubscribe)
	}

	//routes for posts
	post := router.Group("/post", middleware.AuthMiddleware())
	{
		post.GET("/index",controller.UserTest)
	}

	//routes for articles
	article := router.Group("/article", middleware.AuthMiddleware())
	{
		article.GET("/index",controller.ArticleIndex)
	}

	//routes for comments
	comment := router.Group("/comment", middleware.AuthMiddleware())
	{
		comment.GET("/index",controller.CommentIndex)
	}

	//routes for chat
	m := melody.New()

	router.Use(static.Serve("/chat", static.LocalFile("./public", true)))

	router.GET("/chat/ws2", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
}

