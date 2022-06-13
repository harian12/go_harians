package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterV1(r *gin.Engine) {
	v1 := r.Group("v1/")
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
