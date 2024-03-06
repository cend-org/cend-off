package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExportAttach(g *gin.RouterGroup) (err error) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	for _, route := range routes {
		route(g)
	}

	return err
}
