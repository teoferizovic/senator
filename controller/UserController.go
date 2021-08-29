package controller

import "github.com/gin-gonic/gin"

func UserLogin(c *gin.Context) {
	c.String(200, "login")
}

func UserRegister(c *gin.Context) {
	c.String(200, "login")
}

func UserLogout(c *gin.Context) {
	c.String(200, "logout")
}
