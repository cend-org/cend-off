package router

import (
	"github.com/cend-org/duval/internal/router/docs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExtendRoute(r *gin.Engine) (err error) {
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello ok",
		})
	})

	for i := 0; i < len(RootRoutesGroup); i++ {
		group := r.Group(RootRoutesGroup[i].Group)
		err = docs.GenerateDocumentation(group, RootRoutesGroup[i].Paths)
		if err != nil {
			panic(err)
		}
	}

	return err
}
