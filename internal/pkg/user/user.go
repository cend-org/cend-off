package user

import (
	"database/sql"
	"duval/internal/authentication"
	"duval/internal/pkg/user/authorization"
	"duval/internal/pkg/user/password"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Name       string     `json:"name"`
	FamilyName string     `json:"family_name"`
	NickName   string     `json:"nick_name"`
	Email      string     `json:"email"`
	Matricule  string     `json:"matricule"`
	Age        uint       `json:"age"`
	BirthDate  time.Time  `json:"birth_date"`
	Sex        int        `json:"sex"`
	Lang       int        `json:"language"`
	Status     int        `json:"status"`
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

	ctx.JSON(http.StatusOK, user)
}

/*
NewUser creates a new record of user in the system
*/
func NewUser(ctx *gin.Context) {
	var (
		user User
		err  error
	)

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
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

	user.Id, err = database.InsertOne(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
UpdateUser updates the designed user by the id field.
user.id should be provided and user.email should not be empty.
*/
func UpdateUser(ctx *gin.Context) {
	var (
		user User
		err  error
	)

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Invalid mail format",
		})
		return
	}

	if user.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "user id is required for the operation",
		})
		return
	}

	existing, err := GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if existing.Id > 0 && existing.Id != user.Id {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "an other user with the same email already exists !",
		})
		return
	}
	err = database.Update(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
GetUser returns the data of the user designed by the id provided in the url params
*/
func GetUser(ctx *gin.Context) {
	var (
		user User
		err  error
		id   int
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Client.Get(&user, `SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
MyProfile get all the  user's data connected
*/
func MyProfile(ctx *gin.Context) {
	var (
		tok  *authentication.Token
		err  error
		user User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "authentiaction failed",
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

/*
Login takes auth.email and auth.password as parameters.
*/
func Login(ctx *gin.Context) {
	var (
		err  error
		usr  User
		tok  authentication.Token
		auth authentication.Auth
		pass password.Password
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

	err = database.Get(&pass, `SELECT * FROM password WHERE user_id = ? ORDER BY created_at desc  LIMIT 1`, usr.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if !utils.CheckPasswordHash(auth.Password, pass.Psw) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password error",
		})
		return
	}

	tok.UserId = usr.Id

	tokStr, err := authentication.NewAccessToken(tok)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokStr,
	})

	return
}

func NewPassword(ctx *gin.Context) {
	var pass password.Password
	var err error

	err = ctx.ShouldBindJSON(&pass)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if pass.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password should be bound to user",
		})
		return
	}

	if strings.TrimSpace(pass.Psw) == state.EMPTY {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "new password value cannot be empty",
		})
		return
	}

	if utils.PasswordHasValidLength(pass.Psw) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "maximum value of password is 18",
		})
		return
	}

	pass.ContentHash, err = utils.CreateContentHash(pass.Psw)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password should be bound to user",
		})
		return
	}

	pass.Psw, err = utils.HashPassword(pass.Psw)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password should be bound to user",
		})
		return
	}

	_, err = database.Insert(pass)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
	return
}

func GetUserPasswordHistory(ctx *gin.Context) {
	var (
		passwords []password.Password
		err       error
		tok       *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	err = database.Select(&passwords, `SELECT * FROM password WHERE user_id = ? ORDER BY  created_at desc `, tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, passwords)
	return
}

func GetUserAuthorization(ctx *gin.Context) {
	var (
		auth []authorization.Authorization
		err  error
		tok  *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
	}

	auth, err = authorization.GetUserAuthorizations(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, auth)
	return
}

func RemoveUserAuthorization(ctx *gin.Context) {
	var (
		authorizationLevel int
		err                error
		tok                *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
	}

	authorizationLevel, err = strconv.Atoi(ctx.Param("as"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
	}

	err = authorization.DeleteUserAuthorization(tok.UserId, uint(authorizationLevel))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	ctx.Status(http.StatusOK)
	return
}

func RemoveAllUserAuthorization(context *gin.Context) {
	var (
		err error
		tok *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
	}

	err = authorization.DeleteUserAuthorizations(tok.UserId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	context.Status(http.StatusOK)
	return
}

/*

	UTILITIES

*/

func GetUserByEmail(email string) (user User, err error) {
	err = database.Client.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserWithId(id uint) (user User, err error) {
	err = database.Client.Get(&user, `SELECT * FROM USER WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	return user, err
}
