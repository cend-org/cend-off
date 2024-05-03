package contract

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/user/authorization"
	"time"
)

func NewContract(ctx context.Context, input *model.ContractInput) (*model.Contract, error) {
	var (
		tok      *token.Token
		contract model.Contract
		err      error
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &contract, errx.UnAuthorizedError
	}

	if !authorization.IsUserParent(tok.UserId) {
		if !authorization.IsUserTutor(tok.UserId) {
			return &contract, errx.UnAuthorizedError
		}
	}

	if authorization.IsUserParent(tok.UserId) {
		contract.ParentId = tok.UserId
	}

	if authorization.IsUserTutor(tok.UserId) {
		contract.TutorId = tok.UserId
	}

	contract = model.MapContractInputToContract(*input, contract)

	if !authorization.IsUserStudent(contract.StudentId) {
		return &contract, errx.Lambda(errors.New("not a student provided Id"))
	}

	if !authorization.IsUserParent(contract.ParentId) {
		return &contract, errx.Lambda(errors.New("not a parent provided Id"))
	}

	if !authorization.IsUserTutor(contract.TutorId) {
		return &contract, errx.Lambda(errors.New("not a tutor provided Id"))
	}

	contract.Id, err = database.InsertOne(contract)
	if err != nil {
		return &contract, errx.DbInsertError
	}

	return &contract, nil
}

func UpdContract(ctx context.Context, input *model.ContractInput, contractId int) (*model.Contract, error) {
	var (
		tok      *token.Token
		contract model.Contract
		err      error
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &contract, errx.UnAuthorizedError
	}

	if !authorization.IsUserParent(tok.UserId) {
		if !authorization.IsUserTutor(tok.UserId) {
			return &contract, errx.UnAuthorizedError
		}
	}

	if authorization.IsUserParent(tok.UserId) {
		contract.ParentId = tok.UserId
	}

	if authorization.IsUserTutor(tok.UserId) {
		contract.TutorId = tok.UserId
	}

	contract, err = GetContractWithId(contractId)
	if err != nil {
		return &contract, errx.DbGetError
	}
	contract = model.MapContractInputToContract(*input, contract)

	err = database.Update(contract)
	if err != nil {
		return &contract, errx.DbUpdateError
	}

	return &contract, nil
}

func RemoveContract(ctx context.Context, contractId int) (*bool, error) {
	var (
		tok      *token.Token
		contract model.Contract
		status   bool
		err      error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == 0 {
		return &status, errx.UnAuthorizedError
	}

	contract, err = GetContractWithId(contractId)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = database.Delete(contract)
	if err != nil {
		return &status, errx.DbDeleteError
	}
	status = true

	return &status, nil
}

func GetContracts(ctx context.Context) ([]model.Contract, error) {
	var (
		contract []model.Contract
		err      error
		tok      *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return contract, errx.UnAuthorizedError
	}

	contract, err = GetContractsWithId(tok.UserId)
	if err != nil {
		return contract, errx.DbInsertError
	}
	return contract, nil

}

func GetContract(ctx context.Context, contractId int) (*model.Contract, error) {
	var (
		contract model.Contract
		err      error
		tok      *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &contract, errx.UnAuthorizedError
	}

	if !authorization.IsUserParent(tok.UserId) {
		if !authorization.IsUserTutor(tok.UserId) {
			return &contract, errx.UnAuthorizedError
		}
	}

	contract, err = GetContractWithId(contractId)
	if err != nil {
		return &contract, errx.DbGetError
	}
	return &contract, nil
}

// Contract detail

func NewContractTimesheetDetail(ctx context.Context, input *model.ContractTimesheetDetailInput) (*model.ContractTimesheetDetail, error) {
	var (
		contractTimesheetDetail model.ContractTimesheetDetail
		tok                     *token.Token
		err                     error
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &contractTimesheetDetail, errx.UnAuthorizedError
	}
	if tok.UserId == state.ZERO {
		return &contractTimesheetDetail, errx.UnAuthorizedError
	}

	contractTimesheetDetail = model.MapContractTimesheetDetailInputToContractTimesheetDetail(*input, contractTimesheetDetail)

	_, err = database.InsertOne(contractTimesheetDetail)
	if err != nil {
		return &contractTimesheetDetail, err
	}

	return &contractTimesheetDetail, nil
}

func GetContractTimesheetDetail(ctx context.Context) ([]model.ContractTimesheetDetail, error) {
	var (
		tok      *token.Token
		contract []model.ContractTimesheetDetail
		err      error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return contract, errx.UnAuthorizedError
	}

	contract, err = GetContractTimesheetDetailsWithId(tok.UserId)
	if err != nil {
		return contract, errx.DbGetError
	}

	return contract, nil
}

func GetContractTimesheetDetailInfo(ctx context.Context, contractDetailId int) (*model.ContractTimesheetDetail, error) {
	var (
		tok      *token.Token
		contract model.ContractTimesheetDetail
		err      error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &contract, errx.UnAuthorizedError
	}

	contract, err = GetContractTimesheetDetailWithId(tok.UserId, contractDetailId)
	if err != nil {
		return &contract, errx.DbGetError
	}

	return &contract, nil
}

//Salary

func GetTotalSalary(ctx context.Context, studentId int, startDate time.Time, endDate time.Time) (*float64, error) {
	var (
		tok              *token.Token
		totalSalaryValue float64
		salaryValue      int
		hours            float64
		err              error
		timeSheetDetails []model.ContractTimesheetDetail
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &totalSalaryValue, errx.UnAuthorizedError
	}

	if !authorization.IsUserTutor(tok.UserId) {
		return &totalSalaryValue, errx.UnAuthorizedError
	}

	salaryValue, err = GetUserSalaryValue(tok.UserId, studentId)
	if err != nil {
		return &totalSalaryValue, errx.Lambda(err)
	}

	timeSheetDetails, err = GetContractDetailsWithUserId(startDate, endDate, studentId, tok.UserId)
	if err != nil {
		return &totalSalaryValue, errx.Lambda(err)
	}

	for _, timesheetDetail := range timeSheetDetails {
		hours = hours + timesheetDetail.Hours
	}

	totalSalaryValue, err = ComputeTotalSalaryValue(salaryValue, hours)
	if err != nil {
		return &totalSalaryValue, errx.Lambda(err)
	}
	return &totalSalaryValue, nil
}

/*
	UTILS
*/

func GetContractWithId(contractId int) (contract model.Contract, err error) {
	err = database.Get(&contract, `SELECT contract.* FROM contract where contract.id = ?`, contractId)
	if err != nil {
		return contract, err
	}

	return contract, nil
}

func GetContractsWithId(userId int) (contract []model.Contract, err error) {
	if authorization.IsUserParent(userId) {
		err = database.GetMany(&contract, `SELECT contract.* FROM contract where parent_id = ?`, userId)
		if err != nil {
			return contract, err
		}
	}

	if authorization.IsUserTutor(userId) {
		err = database.GetMany(&contract, `SELECT contract.* FROM contract where tutor_id = ?`, userId)
		if err != nil {
			return contract, err
		}
	}

	if authorization.IsUserStudent(userId) {
		err = database.GetMany(&contract, `SELECT contract.* FROM contract where student_id = ?`, userId)
		if err != nil {
			return contract, err
		}
	}

	return contract, nil
}

func GetContractTimesheetDetailWithId(userId int, contractDetailId int) (contractDetail model.ContractTimesheetDetail, err error) {
	err = database.Get(&contractDetail,
		`SELECT contract_timesheet_detail.* FROM contract_timesheet_detail
    JOIN contract ON contract_timesheet_detail.contract_id = contract.id
	WHERE tutor_id = ? AND contract_timesheet_detail.id = ?`, userId, contractDetailId)
	if err != nil {
		return contractDetail, err
	}
	return contractDetail, nil
}

func GetContractTimesheetDetailsWithId(userId int) (contractDetail []model.ContractTimesheetDetail, err error) {
	err = database.GetMany(&contractDetail,
		`SELECT contract_timesheet_detail.* FROM contract_timesheet_detail
   			JOIN contract ON contract_timesheet_detail.contract_id = contract.id
			WHERE tutor_id = ?`, userId)
	if err != nil {
		return contractDetail, err
	}
	return contractDetail, nil
}

func GetUserSalaryValue(tutorId, studentId int) (salaryValue int, err error) {
	err = database.Get(&salaryValue,
		`SELECT contract.salary_value
			FROM contract
			WHERE student_id = ?
			  AND tutor_id = ?`, studentId, tutorId)
	return salaryValue, nil
}

func GetContractDetailsWithUserId(startDate time.Time, endDate time.Time, studentId int, userId int) (timeSheetDetail []model.ContractTimesheetDetail, err error) {
	err = database.GetMany(&timeSheetDetail,
		`SELECT contract_timesheet_detail.*
			FROM contract_timesheet_detail
					 JOIN contract ON contract_timesheet_detail.contract_id = contract.id
			WHERE contract.tutor_id = ? AND contract.student_id = ?
			  AND contract_timesheet_detail.date
				BETWEEN ?
				AND ?`, userId, studentId, startDate.String(), endDate.String())
	if err != nil {
		return timeSheetDetail, err
	}

	return timeSheetDetail, nil
}

func ComputeTotalSalaryValue(salaryValue int, hours float64) (totalSalaryValue float64, err error) {
	totalSalaryValue = float64(salaryValue) * hours
	return totalSalaryValue, nil
}
