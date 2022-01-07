package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
	"net/http"
)

func CommentIndex(ctx *gin.Context) {

	id := ctx.Query("id")

	comments := []model.Comment{}
	err := *new(error)

	if id != "" {
		err = model.GetCommentById(&comments,id)
	} else {
		err = model.GetComments(&comments)
	}


	if len(comments) == 0 {
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
		"data": comments,
	})
	return

}