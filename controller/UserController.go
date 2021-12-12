package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	if (model.User{}.Email == resultUser.Email) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No user with email:" + requestUser.Email,
		})
		return
	}

	//check if password is ok
	err := bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(requestUser.Password))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
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

	log.Info("Successfuly logged in, user:"+requestUser.Email)

	return

}

func UserLogout(ctx *gin.Context) {

	var err error
	token := ctx.Request.Header.Get("Authentication")

	err = service.AddTokenToBlackList(token)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Problem with Logged out",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out",
	})

	log.Info("Successfuly logged out")

	return
}

func UserTest(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Index",
	})
	return

}

func UserIndex(ctx *gin.Context) {

	id := ctx.Param("id")

	//check if user with id exits
	err, resultUser := model.GetByUserId(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No user with id:" + id,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resultUser,
	})
	return

}
