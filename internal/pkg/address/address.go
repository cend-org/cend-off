package address

import (
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Id          uint       `json:"id"`
	Country     string     `json:"country"`
	City        string     `json:"city"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Street      string     `json:"street"`
	FullAddress string     `json:"full_address"`
	XID         string     `json:"xid"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type UserAddress struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UserId      uint       `json:"user_id"`
	AddressId   uint       `json:"address_id"`
	AddressType string     `json:"address_type"`
}

func NewAddress(ctx *gin.Context) {

	var (
		userId      int
		address     Address
		userAddress UserAddress
		err         error
	)
	userId, err = strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to retrieve params",
		})
	}

	err = ctx.ShouldBindJSON(&address)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}
	
	address.Id, err = database.InsertOne(address)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	// Link new address to the current user
	userAddress.UserId = uint(userId)
	userAddress.AddressId = address.Id
	_, err = database.InsertOne(userAddress)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to link user to address database",
		})
		return
	}

	ctx.JSON(http.StatusOK, address)
	return
}

/*
UPDATE ADDRESS OF A USER BY PROVIDING ID IN THE BODY
*/
func UpdateUserAddress(ctx *gin.Context) {
	var (
		address Address
		err     error
	)

	err = ctx.ShouldBindJSON(&address)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}
	if address.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "address id is required for the operation",
		})
		return
	}

	err = database.Update(address)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbUpdateError,
		})
		return
	}

	ctx.JSON(http.StatusOK, address)
	return
}

/*
GET USER ADDRESS  BASED ON user_id PROVIDED IN PARAMS
*/
func GetUserAddress(ctx *gin.Context) {
	var (
		userId  int
		address Address
		err     error
	)

	userId, err = strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Get(&address, `SELECT address.*
    FROM address JOIN user_address 
    ON address.id = user_address.address_id 
    WHERE user_address.user_id = ?`, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGEtError,
		})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

/*
REMOVE USER ADDRESS  BASED ON user_id PROVIDED IN PARAMS
*/

func RemoveUserAddress(ctx *gin.Context) {
	var (
		userId  int
		address Address
		userAddress UserAddress
		err     error
	)

	userId, err = strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Get(&address, `SELECT address.*
    FROM address JOIN user_address 
    ON address.id = user_address.address_id 
    WHERE user_address.user_id = ?`, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGEtError,
		})
		return
	}

	err = database.Delete(address )
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}
	// and remove user_address

	err = database.Get(&userAddress, `SELECT * FROM user_address where user_id = ?`, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGEtError,
		})
		return
	}
	err = database.Delete(userAddress)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}
	ctx.JSON(http.StatusOK, "Address removed successfuly!")
}
