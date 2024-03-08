package trans

import (
	"duval/internal/pkg/translator"
	"github.com/gin-gonic/gin"
)

func ExportAttach(g *gin.RouterGroup) (err error) {
	g.GET("/all", translator.GetAllTranslation)
	g.GET("/one/:language/:id", translator.GetTranslation)
	g.DELETE("/one/:language/:id", translator.DeleteTranslation)
	g.POST("/one", translator.NewTranslation)
	g.PUT("/one", translator.UpdateTranslation)

	return err
}
