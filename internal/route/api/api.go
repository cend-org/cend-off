package api

import (
	"duval/internal/route/api/collection"
	"github.com/gin-gonic/gin"
)

// Routes defines a list of gin function that append the /api group.
var Routes = []func(g *gin.RouterGroup){
	collection.UserRoutes,
	collection.AuthRoutes,
	collection.FileRoutes,
	collection.ChatRoutes,
	collection.PhoneNumberRoutes,
}
