package user

import (
	"duval/internal/authentication"
	"duval/internal/utils"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	LinkTypeParentStudent    = 0
	LinkTypeTutorStudent     = 1
	LinkTypeProfessorStudent = 2
)

type PersonLink struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	InitiatorId uint       `json:"main_user_id"`
	UserId      uint       `json:"act_user_id"`
	LinkType    uint       `json:"link_type"`
}

func GetUserParents(ctx *gin.Context) {
	var (
		parents []User
		err     error
		tok     *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	query := `SELECT u.* 
				FROM person_link pl 
				    JOIN user u ON pl.initiator_id =  u.id 
				WHERE pl.user_id = ? AND pl.link_type = ?`

	err = database.Select(&parents, query, tok.UserId, LinkTypeParentStudent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, parents)
	return
}
