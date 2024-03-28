package planning

import (
	"duval/internal/authentication"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CalendarPlanning struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	AuthorizationId uint       `json:"authorization_id"`
	StartDateTime   time.Time  `json:"start_date_time"`
	EndDateTime     time.Time  `json:"end_date_time"`
	Description     string     `json:"description"`
}

type CalendarPlanningActor struct {
	Id                 uint       `json:"id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	AuthorizationId    uint       `json:"authorization_id"`
	CalendarPlanningId uint       `json:"calendar_planning_id"`
}

//Manage User Plannings

func CreateUserPlannings(ctx *gin.Context) {
	var (
		tok                   *authentication.Token
		err                   error
		calendarPlanning      CalendarPlanning
		calendarPlanningActor CalendarPlanningActor
	)

	err = ctx.ShouldBindJSON(&calendarPlanning)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	//set authorisation_id
	calendarPlanning.AuthorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	calendarId, err := database.InsertOne(calendarPlanning)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}
	//Insert current user to the Calendar Planning actor

	calendarPlanningActor.AuthorizationId = calendarPlanning.AuthorizationId
	calendarPlanningActor.CalendarPlanningId = calendarId
	_, err = database.InsertOne(calendarPlanningActor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, calendarPlanning)
}

func GetUserPlannings(ctx *gin.Context) {
	var (
		tok              *authentication.Token
		err              error
		authorizationId  uint
		calendarPlanning CalendarPlanning
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}
	authorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}
	calendarPlanning, err = GetPlanningById(authorizationId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, calendarPlanning)
}

func RemoveUserPlannings(ctx *gin.Context) {
	var (
		tok              *authentication.Token
		err              error
		authorizationId  uint
		calendarPlanning CalendarPlanning
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}
	authorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}
	calendarPlanning, err = GetPlanningById(authorizationId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	err = database.Delete(calendarPlanning)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

//Manage User Plannings

func GetPlanningActor(ctx *gin.Context) {

	return
}

func AddUserIntoPlanning(ctx *gin.Context) {
	return
}

func RemoveUserFromPlanning(ctx *gin.Context) {
	return
}

/*
	UTILS
*/

func GetUserAuthorizationId(userId uint) (id uint, err error) {
	var userAuthorization authorization.Authorization
	err = database.Get(&userAuthorization, `SELECT * FROM authorization WHERE authorization.user_id = ?`, userId)
	if err != nil {
		return 0, err
	}
	return userAuthorization.Id, nil
}

func GetPlanningById(authorizationId uint) (calendarPlanning CalendarPlanning, err error) {
	err = database.Get(&calendarPlanning, `SELECT *  FROM calendar_planning WHERE calendar_planning.authorization_id = ?`, authorizationId)
	if err != nil {
		return calendarPlanning, err
	}
	return calendarPlanning, err
}
