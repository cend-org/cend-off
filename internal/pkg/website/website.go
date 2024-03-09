package website

import (
	"duval/internal/utils"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Website struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Xid       string     `json:"xid"`
}

func GetWebsites(ctx *gin.Context) {
	var (
		err  error
		webs []Website
	)

	err = database.Select(&webs, `SELECT * FROM website ORDER BY created_at desc `)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, webs)
	return
}
