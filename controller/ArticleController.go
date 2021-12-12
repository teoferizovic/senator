package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
	"net/http"
)

func ArticleIndex(ctx *gin.Context) {


	err, resultArticles := model.GetArticles()

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
