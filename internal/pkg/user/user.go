package user

import (
	"database/sql"
	"duval/internal/authentication"
	"duval/internal/pkg/code"
	"duval/internal/pkg/user/authorization"
	"duval/internal/pkg/user/password"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	StatusNew        = 0
	StatusUnverified = 1
	StatusActive     = 2
)

type User struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Name            string     `json:"name"`
	FamilyName      string     `json:"family_name"`
	NickName        string     `json:"nick_name"`
	Email           string     `json:"email"`
	Matricule       string     `json:"matricule"`
	Age             uint       `json:"age"`
	BirthDate       time.Time  `json:"birth_date"`
	Sex             int        `json:"sex"`
	Lang            int        `json:"language"`
	Status          int        `json:"status"`
	ProfileImageXid string     `json:"profile_image_xid"`
}

/*

	ROUTES

*/

func Register(ctx *gin.Context) {
	var (
		err                error
		user               User
		authorizationLevel int
	)

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	authorizationLevel, err = strconv.Atoi(ctx.Param("as"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	if !utils.IsValidEmail(user.Email) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.InvalidEmailError,
		})
		return
	}

	_, err = GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	if user.Id > state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{
			Message: errx.DuplicateUserError,
		})
		return
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	if user.Name == state.EMPTY {
		user.Name = user.Matricule
	}

	if user.NickName == state.EMPTY {
		user.NickName = user.Matricule
	}

	user.Id, err = database.InsertOne(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	err = authorization.NewUserAuthorization(user.Id, uint(authorizationLevel))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	tokenStr, err := authentication.GetTokenString(user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
	return
}

func VerifyUserEmailValidationCode(ctx *gin.Context) {
	var (
		err            error
		tok            *authentication.Token
		validationCode int
		user           User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	validationCode, err = strconv.Atoi(ctx.Param("code"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	err = code.IsUserVerificationCodeValid(tok.UserId, validationCode)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	user.Status = StatusActive

	err = database.Update(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func SendUserEmailValidationCode(ctx *gin.Context) {
	var (
		err  error
		tok  *authentication.Token
		user User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = code.NewUserVerificationCode(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	user.Status = StatusUnverified

	err = database.Update(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func MyProfile(ctx *gin.Context) {
	var (
		tok  *authentication.Token
		err  error
		user User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

func UpdMyProfile(ctx *gin.Context) {
	var (
		err error
		tok *authentication.Token
		usr User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	usr.Id = tok.UserId
	err = database.Update(usr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, usr)
	return
}

func Login(ctx *gin.Context) {
	var (
		err  error
		usr  User
		auth authentication.Auth
	)

	err = ctx.ShouldBindJSON(&auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	// GET USER DATA
	usr, err = GetUserByEmail(auth.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if !password.IsPasswordValid(usr.Id, auth.Password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	tokenStr, err := authentication.GetTokenString(usr.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})

	return
}

func NewPassword(ctx *gin.Context) {
	var (
		pass password.Password
		err  error
		tok  *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = ctx.ShouldBindJSON(&pass)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	err = password.CreatePassword(tok.UserId, pass)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
	return
}

func GetPasswordHistory(ctx *gin.Context) {
	var (
		passwords []password.Password
		tok       *authentication.Token
		err       error
	)

	time.Sleep(100)
	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	err = database.GetMany(&passwords,
		`SELECT password.*
			FROM password
			WHERE password.user_id = ?
			ORDER BY password.created_at DESC`, tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, passwords)
}

/*

	UTILITIES

*/

func GetUserByEmail(email string) (user User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserWithId(id uint) (user User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	return user, err
}
