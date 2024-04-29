package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExtendRoute(r *gin.Engine) {
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello ok",
		})
	})
}
