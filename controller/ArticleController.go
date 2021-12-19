package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
	"net/http"
)

func ArticleIndex(ctx *gin.Context) {

	id := ctx.Query("id")
	userId := ctx.Query("userId")

	articles := []model.Article{}
	err := *new(error)

	if id != "" && userId == "" {
		err = model.GetArticleById(&articles,id)
	} else if userId != "" && id == "" {
		err = model.GetArticleByUserId(&articles, userId)
	} else if  id != "" && userId != "" {
		err = model.GetArticleByIdAndUserId(&articles,id,userId)
	} else {
		err = model.GetArticles(&articles)
	}

	if len(articles) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found!",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": articles,
	})
	return

}
