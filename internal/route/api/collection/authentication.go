package collection

import (
	"duval/internal/pkg/user"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(g *gin.RouterGroup) {
	g.POST("/login", user.Login)
	g.POST("/password", user.NewPassword)
	g.GET("/password", user.GetUserPasswordHistory)
	return
}
