package contract

import (
	"duval/internal/authentication"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	PerHour     = 0
	PerDay      = 1
	PerWeek     = 2
	PerTwoWeeks = 3
	PerMeet     = 4

	Cash         = 1
	BankTransfer = 2
)

type Contract struct {
	Id            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	TutorId       uint       `json:"tutor_id"`
	ParentId      uint       `json:"parent_id"`
	StudentId     uint       `json:"student_id"`
	StartDate     time.Time  `json:"start_date"`
	EndDate       time.Time  `json:"end_date"`
	PaymentType   uint       `json:"payment_type"`
	SalaryValue   float64    `json:"salary_value"`
	PaymentMethod uint       `json:"payment_method"`
}

type ContractTimesheetDetail struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	ContractId uint       `json:"contract_id"`
	Date       time.Time  `json:"date"`
	Hours      time.Time  `json:"hours"`
}

// Contract

func AddNewContract(ctx *gin.Context) {
	var (
		err      error
		tok      *authentication.Token
		contract Contract
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = ctx.ShouldBindJSON(&contract)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	if authorization.IsUserStudent(tok.UserId) {
		contract.StudentId = tok.UserId

		if contract.ParentId == state.ZERO {
			contract.ParentId = contract.StudentId
		}
	}

	if authorization.IsUserParent(tok.UserId) {
		contract.ParentId = tok.UserId
	}

	if authorization.IsUserTutor(tok.UserId) {
		contract.TutorId = tok.UserId
	}

	if !authorization.IsUserTutor(contract.TutorId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(errors.New("invalid tutor")),
		})
		return
	}

	if !authorization.IsUserParent(contract.ParentId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(errors.New("invalid parent")),
		})
		return
	}

	contractId, err := database.InsertOne(contract)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	contract.Id = contractId
	ctx.AbortWithStatusJSON(http.StatusOK, contract)
}

func GetContract(ctx *gin.Context) {
	var (
		err      error
		tok      *authentication.Token
		contract Contract
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	if authorization.IsUserStudent(tok.UserId) {
		contract, err = GetContractByStudent(tok.UserId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbGetError,
			})
			return
		}
	}

	if authorization.IsUserParent(tok.UserId) {
		contract, err = GetContractByParent(tok.UserId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbGetError,
			})
			return
		}
	}

	if authorization.IsUserTutor(tok.UserId) {
		contract, err = GetContractByTutor(tok.UserId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.DbGetError,
			})
			return
		}
	}

	if authorization.IsUserProfessor(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, contract)
}

func RenewContract(ctx *gin.Context) {
	var (
		err      error
		tok      *authentication.Token
		contract Contract
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = ctx.ShouldBindJSON(&contract)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	if authorization.IsUserStudent(tok.UserId) {
		contract.StudentId = tok.UserId

		if contract.ParentId == state.ZERO {
			contract.ParentId = contract.StudentId
		}
	}

	if authorization.IsUserParent(tok.UserId) {
		contract.ParentId = tok.UserId
	}

	if authorization.IsUserTutor(tok.UserId) {
		contract.TutorId = tok.UserId
	}

	if !authorization.IsUserTutor(contract.TutorId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(errors.New("invalid tutor")),
		})
		return
	}

	if !authorization.IsUserParent(contract.ParentId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(errors.New("invalid parent")),
		})
		return
	}

	err = database.Update(contract)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, contract)
}

func RemoveContract(ctx *gin.Context) {
	var (
		err      error
		tok      *authentication.Token
		contract Contract
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	if authorization.IsUserProfessor(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	err = database.Delete(contract)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "Contract deleted successfully",
	})
}

//Contract Timesheet Detail

func AddNewContractTimesheetDetail(ctx *gin.Context) {
	var (
		err                     error
		tok                     *authentication.Token
		contract                Contract
		contractTimesheetDetail ContractTimesheetDetail
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	if !authorization.IsUserTutor(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	contract, err = GetContractById(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	err = ctx.ShouldBindJSON(&contractTimesheetDetail)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	contractTimesheetDetail.ContractId = contract.Id
	contractTimesheetDetail.Hours, err = GetDuration(contractTimesheetDetail.Date)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(errors.New("errors getting duration")),
		})
		return
	}

	contractTimesheetDetailId, err := database.InsertOne(contractTimesheetDetail)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}
	contractTimesheetDetail.Id = contractTimesheetDetailId
	ctx.AbortWithStatusJSON(http.StatusOK, contractTimesheetDetail)

}

func GetContractTimesheetDetail(ctx *gin.Context) {
	var (
		contractTimesheetDetail ContractTimesheetDetail
		err                     error
		tok                     *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	contractTimesheetDetail, err = GetContractDetailById(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, contractTimesheetDetail)
}

func RemoveContractTimesheetDetail(ctx *gin.Context) {
	var (
		contractTimesheetDetail ContractTimesheetDetail
		err                     error
		tok                     *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	contractTimesheetDetail, err = GetContractDetailById(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	err = database.Delete(contractTimesheetDetail)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbDeleteError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": " Contract timesheet detail  deleted successfully",
	})
}

//Contract total hour

func GetTotalWorkHour(ctx *gin.Context) {
	var (
		totalHour       time.Time
		err             error
		contractDetails []ContractTimesheetDetail
		tok             *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = database.GetMany(&contractDetails, `SELECT contract_timesheet_detail.* FROM contract_timesheet_detail
                                   JOIN contract ON contract_timesheet_detail.contract_id = contract.id
                                   WHERE contract.tutor_id = ?`, tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}
	GetTotalDuration(contractDetails)
	ctx.AbortWithStatusJSON(http.StatusOK, totalHour)
}

/*
	UTILS
*/

func GetContractByStudent(userId uint) (contract Contract, err error) {
	err = database.Get(&contract, `SELECT contract.* FROM contract WHERE contract.student_id = ?`, userId)
	if err != nil {
		return contract, err
	}
	return contract, nil
}

func GetContractByParent(userId uint) (contract Contract, err error) {
	err = database.Get(&contract, `SELECT contract.* FROM contract WHERE contract.parent_id = ?`, userId)
	if err != nil {
		return contract, err
	}
	return contract, nil
}

func GetContractByTutor(userId uint) (contract Contract, err error) {
	err = database.Get(&contract, `SELECT contract.* FROM contract WHERE contract.tutor_id = ?`, userId)
	if err != nil {
		return contract, err
	}
	return contract, nil
}

func GetContractById(userId uint) (contract Contract, err error) {
	if authorization.IsUserStudent(userId) {
		contract, err = GetContractByStudent(userId)
		if err != nil {
			return contract, err
		}
	}

	if authorization.IsUserParent(userId) {
		contract, err = GetContractByParent(userId)
		return contract, err
	}

	if authorization.IsUserTutor(userId) {
		contract, err = GetContractByTutor(userId)
		return contract, err
	}
	return contract, nil
}

func GetContractDetailById(userId uint) (contract ContractTimesheetDetail, err error) {
	if authorization.IsUserStudent(userId) {
		contract, err = GetContractDetailByStudent(userId)
		if err != nil {
			return contract, err
		}
	}

	if authorization.IsUserParent(userId) {
		contract, err = GetContractDetailByParent(userId)
		return contract, err
	}

	if authorization.IsUserTutor(userId) {
		contract, err = GetContractDetailByTutor(userId)
		return contract, err
	}
	return contract, nil
}

func GetContractDetailByStudent(userId uint) (contract ContractTimesheetDetail, err error) {
	err = database.Get(&contract, `SELECT contract_timesheet_detail.* FROM contract_timesheet_detail
                                   JOIN contract ON contract_timesheet_detail.contract_id = contract.id
                                   WHERE contract.student_id = ?`, userId)
	if err != nil {
		return contract, err
	}
	return contract, nil
}

func GetContractDetailByParent(userId uint) (contract ContractTimesheetDetail, err error) {
	err = database.Get(&contract, `SELECT contract_timesheet_detail.* FROM contract_timesheet_detail
                                   JOIN contract ON contract_timesheet_detail.contract_id = contract.id
                                   WHERE contract.parent_id = ?`, userId)
	if err != nil {
		return contract, err
	}
	return contract, nil
}

func GetContractDetailByTutor(userId uint) (contract ContractTimesheetDetail, err error) {
	err = database.Get(&contract, `SELECT contract_timesheet_detail.* FROM contract_timesheet_detail
                                   JOIN contract ON contract_timesheet_detail.contract_id = contract.id
                                   WHERE contract.tutor_id = ?`, userId)
	if err != nil {
		return contract, err
	}
	return contract, nil
}

func GetDuration(startTime time.Time) (total time.Time, err error) {
	var (
		duration time.Duration
	)
	duration = time.Now().Sub(startTime)
	hours := duration.Hours()
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	currentDuration := fmt.Sprintf("%02d:%02d:%02d", int(hours), minutes, seconds)

	total, err = time.Parse(time.TimeOnly, currentDuration)
	if err != nil {
		return total, err
	}
	return total, nil
}

func GetTotalDuration(contractDetails []ContractTimesheetDetail) (string, error) {
	time.Sleep(100)
	var duration int
	for _, contractDetail := range contractDetails {
		duration = contractDetail.Hours.Second()
	}
	seconds := duration
	minutes := duration % 60
	hours := duration % 3600
	return fmt.Sprintf("%02d:%02d:%02d", int(hours), int(minutes), int(seconds)), nil

}
