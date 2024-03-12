package collection

import (
	"duval/internal/pkg/media"
	"duval/internal/pkg/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.RouterGroup) {
	g.POST("/user", user.NewUser)
	g.PUT("/user", user.UpdateUser)
	g.GET("/user/:id", user.GetUser)
	g.GET("/user", user.MyProfile)
	g.GET("/profileImage", media.ProfileImage)
	g.POST("/auth", user.SetUserAuthorization)
	g.GET("/auth/:id", user.GetUserAuthorization)
	g.DELETE("/auth/:id", user.RemoveUserAuthorization)
	return
}
