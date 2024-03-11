package phone

import (
	"duval/internal/utils"
	"duval/pkg/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PhoneNumber struct {
	Id                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	MobilePhoneNumber string     `json:"mobile_phone_number"`
	IsUrgency         bool       `json:"is_urgency"`
}

type UserPhoneNumber struct {
	UserId        uint `json:"user_id"`
	PhoneNumberId uint `json:"phone_number_id"`
}

/*

	ROUTES Handlers

*/

/*
ADD NEW PHONE NUMBER TO A USER BY PORVIDING user.id in params AND RETRUN PHONE NUMBER
*/
func NewPhoneNumber(ctx *gin.Context) {
	var (
		phone PhoneNumber
		err   error
	)

	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = ctx.ShouldBindJSON(&phone)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if !utils.PhoneValidator(phone.MobilePhoneNumber) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Invalid format of phone number",
		})
		return
	}

	phone.Id, err = database.InsertOne(phone)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Phone number already exist in the database",
		})
		return
	}
	// Link phone to user .
	_, err = database.Client.Exec(`INSERT INTO user_phone_number VALUES (? , ?)`, userId, phone.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to link phone number to user",
		})
		return
	}

	ctx.JSON(http.StatusOK, phone)
	return
}

/*
UPDATE PHONE NUMBER OF A USER BY PORVIDING user.id in params AND LIST OF USER PHONE NUMBER
*/
func UpdateUserPhoneNumber(ctx *gin.Context) {
	var (
		id              int
		phone           PhoneNumber
		newPhone        PhoneNumber
		userPhoneNumber UserPhoneNumber
		err             error
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = ctx.ShouldBindJSON(&newPhone)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if !utils.PhoneValidator(newPhone.MobilePhoneNumber) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Invalid format of phone number",
		})
		return
	}

	err = database.Client.Get(&userPhoneNumber, `SELECT * FROM user_phone_number WHERE user_id = ?`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to fetch user phone number  user_id unknown",
		})
		return
	}

	err = database.Client.Get(&phone, `SELECT * FROM phone_number WHERE id = ?`, userPhoneNumber.PhoneNumberId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to fetch phone number  user_id unknown",
		})
		return
	}

	if phone.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "No phone number found please add new",
		})
	}

	newMobilePhone, err := database.Client.Exec(`UPDATE phone_number SET mobile_phone_number = ? WHERE id = ?`, newPhone.MobilePhoneNumber, userPhoneNumber.PhoneNumberId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to update mobile phone number",
		})
		return
	}

	ctx.JSON(http.StatusOK, newMobilePhone)
	return
}

/*
GET USER PHONE NUMBER BASED ON user_id PROVIDED IN PARAMS
*/

func GetUserPhoneNumber(ctx *gin.Context) {
	var (
		id              int
		phone           PhoneNumber
		userPhoneNumber UserPhoneNumber
		err             error
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}
	err = database.Client.Get(&userPhoneNumber, `SELECT * FROM user_phone_number WHERE user_id = ?`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to fetch user phone number  user_id unknown",
		})
		return
	}

	err = database.Client.Get(&phone, `SELECT * FROM phone_number WHERE id = ?`, userPhoneNumber.PhoneNumberId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to fetch phone number  user_id unknown",
		})
		return
	}

	ctx.JSON(http.StatusOK, phone)
}

/*
UTILS
*/
