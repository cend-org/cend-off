package planning

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
)

func CreateUserPlannings(ctx context.Context, input *model.CalendarPlanningInput) (*model.CalendarPlanning, error) {
	var (
		tok                   *token.Token
		err                   error
		calendarPlanning      model.CalendarPlanning
		calendarPlanningActor model.CalendarPlanningActor
	)

	calendarPlanning = model.MapCalendarPlanningInputToCalendarPlanning(*input, calendarPlanning)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &calendarPlanning, errx.UnAuthorizedError
	}

	calendarPlanning.AuthorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		return &calendarPlanning, errx.DbGetError
	}

	calendarId, err := database.InsertOne(calendarPlanning)
	if err != nil {
		return &calendarPlanning, errx.DbInsertError
	}

	calendarPlanningActor.AuthorizationId = calendarPlanning.AuthorizationId
	calendarPlanningActor.CalendarPlanningId = calendarId

	err = AddCalendarPlanningActor(calendarPlanningActor)
	if err != nil {
		return &calendarPlanning, errx.DbInsertError
	}

	return &calendarPlanning, nil
}

func AddUserIntoPlanning(ctx context.Context, calendarId int, selectedUserId int) (*model.CalendarPlanningActor, error) {
	var (
		tok                   *token.Token
		calendarPlanningActor model.CalendarPlanningActor
		err                   error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &calendarPlanningActor, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &calendarPlanningActor, errx.UnAuthorizedError
	}

	authorizationId, err := GetUserAuthorizationId(selectedUserId)
	if err != nil {
		return &calendarPlanningActor, errx.DbGetError
	}

	calendarPlanningActor.AuthorizationId = authorizationId
	calendarPlanningActor.CalendarPlanningId = calendarId

	err = AddCalendarPlanningActor(calendarPlanningActor)
	if err != nil {
		return &calendarPlanningActor, errx.DbInsertError
	}

	return &calendarPlanningActor, nil
}

func GetUserPlannings(ctx context.Context) (*model.CalendarPlanning, error) {
	var (
		tok              *token.Token
		err              error
		authorizationId  int
		calendarPlanning model.CalendarPlanning
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &calendarPlanning, errx.UnAuthorizedError
	}

	authorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		return &calendarPlanning, errx.DbGetError
	}

	calendarPlanning, err = GetPlanningById(authorizationId)
	if err != nil {
		return &calendarPlanning, errx.DbGetError
	}

	return &calendarPlanning, nil
}

func GetPlanningActors(ctx context.Context, calendarId int) ([]model.User, error) {
	var (
		err                    error
		calendarPlanningActors []model.User
	)
	calendarPlanningActors, err = GetPlanningActorByCalendarId(calendarId)
	if err != nil {
		return calendarPlanningActors, errx.DbGetError
	}

	return calendarPlanningActors, nil
}

func RemoveUserFromPlanning(ctx context.Context, calendarPlanningId int, selectedUserId int) (*bool, error) {
	var (
		selectedCalendarPlanningActor model.CalendarPlanningActor
		tok                           *token.Token
		err                           error
		status                        bool
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}

	selectedCalendarPlanningActor, err = GetSelectedPlanningActor(selectedUserId, calendarPlanningId)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = RemoveSelectedPlanningActor(selectedCalendarPlanningActor)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	status = true

	return &status, nil
}

func RemoveUserPlannings(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented: RemoveUserPlannings - removeUserPlannings"))
}

/*
	UTILS
*/

func GetUserAuthorizationId(UserId int) (id int, err error) {
	var userAuthorization model.Authorization
	err = database.Get(&userAuthorization, `SELECT * FROM authorization WHERE user_id = ?`, UserId)
	if err != nil {
		return state.ZERO, err
	}
	return userAuthorization.Id, nil
}

func GetPlanningById(authorizationId int) (calendarPlanning model.CalendarPlanning, err error) {
	err = database.Get(&calendarPlanning, `SELECT *  FROM calendar_planning WHERE authorization_id = ?`, authorizationId)
	if err != nil {
		return calendarPlanning, err
	}
	return calendarPlanning, err
}

func AddCalendarPlanningActor(calendarPlanningActor model.CalendarPlanningActor) (err error) {
	_, err = database.InsertOne(calendarPlanningActor)
	if err != nil {
		return err
	}
	return nil
}

func GetPlanningActorByCalendarId(calendarId int) (calendarPlanningActors []model.User, err error) {
	err = database.Select(&calendarPlanningActors,
		`SELECT user.* FROM user
              JOIN authorization auth ON user.id = auth.user_id
              JOIN calendar_planning_actor cpa ON auth.id = cpa.authorization_id
              JOIN calendar_planning cp ON cpa.calendar_planning_id = cp.id
     WHERE cp.id = ?`, calendarId)
	if err != nil {
		return calendarPlanningActors, err
	}
	return calendarPlanningActors, err
}

func GetSelectedPlanningActor(UserId int, calendarPlanningId int) (calendarPlanningActor model.CalendarPlanningActor, err error) {
	err = database.Get(&calendarPlanningActor,
		`SELECT cpa.*  FROM calendar_planning_actor cpa
                                  JOIN authorization auth ON cpa.authorization_id = auth.id
                                  JOIN calendar_planning cp ON cpa.calendar_planning_id = cp.id
                                  JOIN user ON auth.user_id = user.id
                                  WHERE user.id= ? AND cp.id = ?`, UserId, calendarPlanningId)
	if err != nil {
		return calendarPlanningActor, err
	}
	return calendarPlanningActor, nil
}

func RemoveSelectedPlanningActor(calendarPlanningActor model.CalendarPlanningActor) (err error) {
	err = database.Delete(calendarPlanningActor)
	if err != nil {
		return err
	}
	return nil
}
