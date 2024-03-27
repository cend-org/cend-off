package link

import (
	"duval/internal/authentication"
	"duval/internal/pkg/user"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// AuthorizationLevel

	StudentAuthorizationLevel   = 0
	ParentAuthorizationLevel    = 1
	TutorAuthorizationLevel     = 2
	ProfessorAuthorizationLevel = 3

	//Link_type

	Student_Parent    = 0
	Student_Tutor     = 1
	Student_Professor = 2
)

func GetUserParent(ctx *gin.Context) {
	var (
		tok    *authentication.Token
		parent user.User
		err    error
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	if tok.UserId == state.ZERO {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.UnAuthorizedError,
			})
			return
		}
	}
	err = database.Get(&parent, `SELECT user.* FROM user
   								JOIN authorization ON user.id = authorization.user_id
   								JOIN user_authorization_link_actor ON authorization.id = user_authorization_link_actor.authorization_id
   								JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
             						WHERE authorization.level = ? AND user_authorization_link.link_type = ?`, ParentAuthorizationLevel, Student_Parent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, parent)

}

func GetUserTutor(ctx *gin.Context) {
	var (
		tok   *authentication.Token
		tutor user.User
		err   error
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	if tok.UserId == state.ZERO {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.UnAuthorizedError,
			})
			return
		}
	}

	err = database.Get(&tutor, `SELECT user.* FROM user
   								JOIN authorization ON user.id = authorization.user_id
   								JOIN user_authorization_link_actor ON authorization.id = user_authorization_link_actor.authorization_id
   								JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
             						WHERE authorization.level = ? AND user_authorization_link.link_type = ?`, TutorAuthorizationLevel, Student_Tutor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, tutor)
}

func GetUserProfessor(ctx *gin.Context) {
	var (
		tok       *authentication.Token
		professor user.User
		err       error
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	if tok.UserId == state.ZERO {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.UnAuthorizedError,
			})
			return
		}
	}

	err = database.Get(&professor, `SELECT user.* FROM user
   								JOIN authorization ON user.id = authorization.user_id
   								JOIN user_authorization_link_actor ON authorization.id = user_authorization_link_actor.authorization_id
   								JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
             						WHERE authorization.level = ? AND user_authorization_link.link_type = ?`, ProfessorAuthorizationLevel, Student_Professor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, professor)
}
