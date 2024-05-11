package router

import (
	mediafile "github.com/cend-org/duval/pkg/media"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExtendRoute(r *gin.Engine) {
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello ok",
		})
	})
	r.POST("/upload", mediafile.Upload)
}
