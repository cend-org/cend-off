package route

import (
	"duval/internal/route/api"
	"duval/internal/route/translation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func attach(g *gin.Engine) (err error) {
	g.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Connected")
	})

	apiGroup := g.Group("/api")
	err = api.ExportAttach(apiGroup)
	if err != nil {
		panic(err)
	}

	translationGroup := g.Group("/translation")
	err = trans.ExportAttach(translationGroup)
	if err != nil {
		panic(err)
	}

	return err
}
