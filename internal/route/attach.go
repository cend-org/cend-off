package route

import (
	"duval/internal/route/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func attach(g *gin.Engine) (err error) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	//g.GET("/tok", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"token": authentication.CreateNewToken(),
	//	})
	//})

	apiGroup := g.Group("/api")
	err = api.ExportAttach(apiGroup)
	if err != nil {
		panic(err)
	}

	return err
}
