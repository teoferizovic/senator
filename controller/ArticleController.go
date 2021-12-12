package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
	"net/http"
)

func ArticleIndex(ctx *gin.Context) {

	id := ctx.Query("id")
	userId := ctx.Query("userId")

	resultArticles := []model.Article{}
	err := *new(error)

	if id != "" && userId == "" {
		err, resultArticles = model.GetArticleById(id)
	} else if userId != "" && id == "" {
		err, resultArticles = model.GetArticleByUserId(userId)
	} else if  id != "" && userId != "" {
		err, resultArticles = model.GetArticleByIdAndUserId(id,userId)
	} else {
		err, resultArticles = model.GetArticles()
	}

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No data found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resultArticles,
	})
	return

}
