package collection

import (
	"duval/internal/pkg/media"
	"github.com/gin-gonic/gin"
)

func FileRoutes(g *gin.RouterGroup) {
	g.POST("/upload", media.Upload)
	g.Static("/public", "public")
	return
}
