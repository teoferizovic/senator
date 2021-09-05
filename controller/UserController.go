package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
)

func UserLogin(c *gin.Context) {
	fmt.Println(model.GetUsers())
	c.String(200, "login")
}

func UserRegister(c *gin.Context) {
	c.String(200, "login")
}

func UserLogout(c *gin.Context) {
	c.String(200, "logout")
}
