package contract

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/pkg/user/authorization"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
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

func RemoveContract(ctx context.Context, contractId int) (*string, error) {
	var (
		tok      *token.Token
		contract model.Contract
		status   string
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
	status = "success"

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
	panic(fmt.Errorf("not implemented: RemoveUserPlannings - removeUserPlannings"))
}

func NewContractTimesheetDetail(ctx context.Context, input *model.ContractTimesheetDetailInput) (*model.ContractTimesheetDetail, error) {
	var (
		contractTimesheetDetail model.ContractTimesheetDetail
		contract                model.Contract
		tok                     *token.Token
		err                     error
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &contractTimesheetDetail, errx.UnAuthorizedError
	}

	contract, err = GetContractWithId(tok.UserId)
	if err != nil {
		return &contractTimesheetDetail, errx.DbGetError
	}

	contractTimesheetDetail = model.MapContractTimesheetDetailInputToContractTimesheetDetail(*input, contractTimesheetDetail)

	contractTimesheetDetail.ContractId = contract.Id

	_, err = database.InsertOne(contractTimesheetDetail)
	if err != nil {
		return &contractTimesheetDetail, err
	}

	return &contractTimesheetDetail, nil
}

func GetContractTimesheetDetail(ctx context.Context) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: GetContractTimesheetDetail - getContractTimesheetDetail"))
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
