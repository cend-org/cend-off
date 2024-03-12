package collection

import (
	"duval/internal/pkg/phone"
	"github.com/gin-gonic/gin"
)

func PhoneNumberRoutes(g *gin.RouterGroup) {
	g.POST("/phone/:user_id", phone.NewPhoneNumber)
	g.PUT("/phone/", phone.UpdateUserPhoneNumber)
	g.GET("/phone/:user_id", phone.GetUserPhoneNumber)
	return
}
