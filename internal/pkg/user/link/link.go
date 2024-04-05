package link

import (
	"duval/internal/authentication"
	"duval/internal/pkg/user"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	// AuthorizationLevel

	StudentAuthorizationLevel   = 0
	ParentAuthorizationLevel    = 1
	TutorAuthorizationLevel     = 2
	ProfessorAuthorizationLevel = 3

	//Link_type

	StudentParent    = 0
	StudentTutor     = 1
	StudentProfessor = 2
)

type UserAuthorizationLink struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	LinkType  uint       `json:"link_type"`
}

type UserAuthorizationLinkActor struct {
	Id                      uint       `json:"id"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
	DeletedAt               *time.Time `json:"deleted_at"`
	UserAuthorizationLinkId uint       `json:"user_authorization_link_id"`
	AuthorizationId         uint       `json:"authorization_id"`
}

// Parent Handler

func GetUserParent(ctx *gin.Context) {
	var (
		tok     *authentication.Token
		parents []user.User
		err     error
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
	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbGetError,
			})
			return
		}
	}

	parents, err = GetLink(auth.Id, ParentAuthorizationLevel, StudentParent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, parents)

}

func AddParentToUser(ctx *gin.Context) {
	var (
		tok                     *authentication.Token
		parent                  user.User
		userAuthorizationLinkId uint
		err                     error
	)
	time.Sleep(100)
	// Select User

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	//check if user is a student
	if !authorization.IsUserStudent(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "User is not a student ",
		})
		return
	}

	// Select Parent from body
	err = ctx.ShouldBindJSON(&parent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	// Check if parent  doesn't exist in the database based on name and family name then  create a user named parent if not
	currentParent, err := GetUserByUserName(parent)
	if currentParent.Id == state.ZERO {
		//	Create parent with email parent+1@cend.intra
		currentParent, err = CreateNewUser(parent)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbInsertError,
			})
			return
		}

	}

	//Check if link already exist if not then create new link and add creator into link actor by default
	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	userAuthorizationLinkId, err = GetUserLink(StudentParent, auth.Id)
	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(StudentParent, tok.UserId, tok.UserLevel)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbInsertError,
			})
			return
		}
	}

	//Check if parent is already added to the user
	currentParentAuth, err := authorization.GetUserAuthorization(currentParent.Id, ParentAuthorizationLevel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	_, err = GetLink(currentParentAuth.Id, ParentAuthorizationLevel, StudentParent)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DuplicateUserError,
		})
		return
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, currentParent.Id, ParentAuthorizationLevel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, currentParent)
}

func RemoveUserParent(ctx *gin.Context) {
	var (
		parent user.User
		actor  UserAuthorizationLinkActor
		tok    *authentication.Token
		err    error
	)
	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	//Check if user is authorized to delete a parent
	if tok.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	//Select parent from body
	err = ctx.ShouldBindJSON(&parent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	//Get selected parent user_authorization_link_actor
	actor, err = GetSelectedUserLinkActor(parent, StudentParent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}
	//Delete selected parent form user_authorization_link_actor
	err = DeleteUserLinkActor(actor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}
}

// Tutor Handler

func GetUserTutor(ctx *gin.Context) {
	var (
		tok   *authentication.Token
		tutor []user.User
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

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbGetError,
			})
			return
		}
	}

	tutor, err = GetLink(auth.Id, TutorAuthorizationLevel, StudentTutor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, tutor)
}

func AddTutorToUser(ctx *gin.Context) {
	var (
		err                     error
		tok                     *authentication.Token
		userAuthorizationLinkId uint
		tutor                   user.User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	//check if user is a student
	if !authorization.IsUserStudent(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "User is not a student ",
		})
		return
	}

	err = ctx.ShouldBindJSON(&tutor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	userAuthorizationLinkId, err = GetUserLink(StudentTutor, auth.Id)
	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(StudentParent, tok.UserId, tok.UserLevel)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbInsertError,
			})
			return
		}
	}

	currentTutor, err := GetUserByUserName(tutor)

	//Check if parent is already added to the user
	currentTutorAuth, err := authorization.GetUserAuthorization(currentTutor.Id, TutorAuthorizationLevel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	_, err = GetLink(currentTutorAuth.Id, TutorAuthorizationLevel, StudentTutor)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DuplicateUserError,
		})
		return
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, currentTutor.Id, TutorAuthorizationLevel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, currentTutor)
}

func RemoveUserTutor(ctx *gin.Context) {
	var (
		tutor user.User
		actor UserAuthorizationLinkActor
		tok   *authentication.Token
		err   error
	)
	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	//Check if user is authorized to delete a tutor
	if tok.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	//Select tutor from body
	err = ctx.ShouldBindJSON(&tutor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	//Get selected tutor user_authorization_link_actor
	actor, err = GetSelectedUserLinkActor(tutor, StudentTutor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}
	//Delete selected tutor form user_authorization_link_actor
	err = DeleteUserLinkActor(actor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}
}

// Professor Handler

func GetUserProfessor(ctx *gin.Context) {
	var (
		tok       *authentication.Token
		professor []user.User
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

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbGetError,
			})
			return
		}
	}

	professor, err = GetLink(auth.Id, ProfessorAuthorizationLevel, StudentProfessor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, professor)
}

func RemoveUserProfessor(ctx *gin.Context) {
	var (
		professor user.User
		actor     UserAuthorizationLinkActor
		tok       *authentication.Token
		err       error
	)
	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	//Check if user is authorized to delete a professor
	if tok.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	//Select professor from body
	err = ctx.ShouldBindJSON(&professor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	//Get selected professor user_authorization_link_actor
	actor, err = GetSelectedUserLinkActor(professor, StudentProfessor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}
	//Delete selected professor form user_authorization_link_actor
	err = DeleteUserLinkActor(actor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}
}

/*
	UTILS
*/

func SetUserAuthorizationLink(linkType uint, userId uint, userLevel uint) (userAuthorizationLinkId uint, err error) {
	var (
		userAuthorizationLink UserAuthorizationLink
	)

	userAuthorizationLink.LinkType = linkType

	userAuthorizationLinkId, err = database.InsertOne(userAuthorizationLink)
	if err != nil {
		return userAuthorizationLinkId, err
	}
	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, userId, userLevel)
	if err != nil {
		return userAuthorizationLinkId, err
	}
	return userAuthorizationLinkId, nil
}

func SetUserAuthorizationLinkActor(linkId uint, userId uint, level uint) (err error) {
	var userAuthorizationLinkActor UserAuthorizationLinkActor

	auth, err := authorization.GetUserAuthorization(userId, level)
	if err != nil {
		return err
	}
	userAuthorizationLinkActor.AuthorizationId = auth.Id
	userAuthorizationLinkActor.UserAuthorizationLinkId = linkId

	if err != nil {
		return err
	}
	_, err = database.InsertOne(userAuthorizationLinkActor)
	if err != nil {
		return err
	}

	return nil
}

func CreateNewUser(user user.User) (currentUser user.User, err error) {

	user.Email = "parent+1@cend.intern"
	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return user, err
	}

	if user.Name == state.EMPTY {
		user.Name = user.Matricule
	}

	if user.NickName == state.EMPTY {
		user.NickName = user.Matricule
	}

	user.Id, err = database.InsertOne(user)
	if err != nil {
		return user, err
	}

	err = authorization.NewUserAuthorization(user.Id, uint(ParentAuthorizationLevel))
	if err != nil {
		return user, err
	}

	_, err = authentication.GetTokenString(user.Id)
	if err != nil {
		return user, err
	}
	currentUser = user
	return currentUser, nil
}

func GetLink(authId uint, authorizationLevel uint, linkType uint) (link []user.User, err error) {
	err = database.GetMany(&link, `SELECT user.* FROM user
                       JOIN authorization ON user.id = authorization.user_id
                       JOIN user_authorization_link_actor ON authorization.id = user_authorization_link_actor.authorization_id
                       JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
WHERE user_authorization_link.id =  (
    SELECT user_authorization_link_actor.user_authorization_link_id
    FROM user_authorization_link_actor
             JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
    WHERE user_authorization_link_actor.authorization_id = ? AND user_authorization_link.link_type = ?
    )
AND authorization.level = ?`, authId, linkType, authorizationLevel)
	if err != nil {
		return link, err
	}

	return link, nil
}

func GetUserByUserName(currentUser user.User) (user user.User, err error) {
	err = database.Get(&user, `SELECT user.* FROM user WHERE user.name = ? and user.family_name = ?`, currentUser.Name, currentUser.FamilyName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserLink(linkType uint, authorizationId uint) (linkId uint, err error) {
	var userLink UserAuthorizationLink
	err = database.Get(&userLink,
		`SELECT user_authorization_link.* FROM user_authorization_link
                                  JOIN user_authorization_link_actor ON user_authorization_link.id = user_authorization_link_actor.user_authorization_link_id
                                  WHERE user_authorization_link.link_type = ? AND user_authorization_link_actor.authorization_id = ?`, linkType, authorizationId)
	if err != nil {
		return 0, err
	}
	return userLink.Id, nil
}

func GetSelectedUserLinkActor(user user.User, linkType uint) (actor UserAuthorizationLinkActor, err error) {
	err = database.Get(&actor,
		`SELECT user_authorization_link_actor.*
FROM user_authorization_link_actor
JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
JOIN authorization ON user_authorization_link_actor.authorization_id = authorization.id
JOIN user ON authorization.user_id = user.id
WHERE user.family_name = ? AND  user.name = ? AND user_authorization_link.link_type = ?`, user.FamilyName, user.Name, linkType)
	if err != nil {
		return actor, err
	}
	return actor, nil
}

func DeleteUserLinkActor(userAuthorizationLinkActor UserAuthorizationLinkActor) (err error) {
	err = database.Delete(userAuthorizationLinkActor)
	if err != nil {
		return err
	}
	return nil
}
