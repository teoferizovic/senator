package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/database"
	"github.com/teoferizovic/senator/model"
	"github.com/teoferizovic/senator/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UserRegister(ctx *gin.Context) {

	var requestUser model.User

	//check if right credeitals are sent
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newUser, err := model.CreateUser(&requestUser)

	//if email still exists or insert problems
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//return 200
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created user with email:" + newUser.Email,
	})
	return
}

func UserLogin(ctx *gin.Context) {

	var requestUser model.User

	//check if right credeitals are sent
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//check if user with email exits
	resultUser := model.GetByEmail(&requestUser)

	if (model.User{}) == resultUser {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "No user with email:" + requestUser.Email,
		})
		return
	}

	//check if password is ok
	err := bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(requestUser.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	tokenString, err := service.CreateToken(requestUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Token for user is:" + tokenString,
	})
	return

}

func UserLogout(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out",
	})
	return
}

func UserIndex(ctx *gin.Context) {

	pong, err := database.Redis.Ping().Result()
	fmt.Println(pong, err)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Index",
	})
	return

}
