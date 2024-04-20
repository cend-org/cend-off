package planning

import (
	"context"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
)

func CreateUserPlannings(ctx *context.Context, input *model.NewCalendarPlanning) (*model.CalendarPlanning, error) {
	var (
		tok                   *authentication.Token
		err                   error
		calendarPlanning      model.CalendarPlanning
		calendarPlanningActor model.CalendarPlanningActor
	)
	calendarPlanning.StartDateTime = input.StartDateTime
	calendarPlanning.EndDateTime = input.EndDateTime
	calendarPlanning.Description = input.Description

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &calendarPlanning, errx.Lambda(err)
	}

	calendarPlanning.AuthorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		return &calendarPlanning, errx.Lambda(err)
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

func GetUserPlannings(ctx *context.Context) (*model.CalendarPlanning, error) {
	var (
		tok              *authentication.Token
		err              error
		authorizationId  uint
		calendarPlanning model.CalendarPlanning
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &calendarPlanning, errx.Lambda(err)
	}

	authorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		return &calendarPlanning, errx.Lambda(err)
	}

	calendarPlanning, err = GetPlanningById(authorizationId)
	if err != nil {
		return &calendarPlanning, errx.Lambda(err)
	}

	return &calendarPlanning, nil
}

func RemoveUserPlannings(ctx *context.Context) (*string, error) {
	var (
		tok              *authentication.Token
		err              error
		authorizationId  uint
		calendarPlanning model.CalendarPlanning
		status           string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	authorizationId, err = GetUserAuthorizationId(tok.UserId)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	calendarPlanning, err = GetPlanningById(authorizationId)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	err = database.Delete(calendarPlanning)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	status = "success"
	return &status, nil
}

func AddUserIntoPlanning(ctx *context.Context, calendarId int, selectedUserId int) (*model.CalendarPlanningActor, error) {
	var (
		tok                   *authentication.Token
		calendarPlanningActor model.CalendarPlanningActor
		err                   error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &calendarPlanningActor, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &calendarPlanningActor, errx.UnAuthorizedError
	}

	authorizationId, err := GetUserAuthorizationId(uint(selectedUserId))
	if err != nil {
		return &calendarPlanningActor, errx.Lambda(err)
	}

	calendarPlanningActor.AuthorizationId = authorizationId
	calendarPlanningActor.CalendarPlanningId = uint(calendarId)

	err = AddCalendarPlanningActor(calendarPlanningActor)
	if err != nil {
		return &calendarPlanningActor, errx.DbInsertError
	}

	return &calendarPlanningActor, nil
}

func GetPlanningActors(ctx *context.Context, calendarId int) ([]*model.User, error) {
	var (
		err                           error
		calendarPlanningActors        []model.User
		currentCalendarPlanningActors []*model.User
	)

	calendarPlanningActors, err = GetPlanningActorByCalendarId(uint(calendarId))
	if err != nil {
		return currentCalendarPlanningActors, errx.DbGetError
	}

	for _, actor := range calendarPlanningActors {
		currentCalendarPlanningActors = append(currentCalendarPlanningActors, &model.User{
			Id:          actor.Id,
			Name:        actor.Name,
			FamilyName:  actor.FamilyName,
			NickName:    actor.NickName,
			Description: actor.Description,
			CoverText:   actor.CoverText,
			Profile:     actor.Profile,
			AddOnTitle:  actor.AddOnTitle,
		})
	}
	return currentCalendarPlanningActors, errx.DbGetError
}

func RemoveUserFromPlanning(ctx *context.Context, calendarPlanningId int, selectedUserId int) (*string, error) {
	var (
		selectedCalendarPlanningActor model.CalendarPlanningActor
		tok                           *authentication.Token
		err                           error
		status                        string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}

	selectedCalendarPlanningActor, err = GetSelectedPlanningActor(uint(selectedUserId), uint(calendarPlanningId))
	if err != nil {
		return &status, errx.DbGetError
	}

	err = RemoveSelectedPlanningActor(selectedCalendarPlanningActor)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	return &status, nil
}

/*
	UTILS
*/

func GetUserAuthorizationId(userId uint) (id uint, err error) {
	var userAuthorization model.Authorization
	err = database.Get(&userAuthorization, `SELECT * FROM authorization WHERE authorization.user_id = ?`, userId)
	if err != nil {
		return 0, err
	}
	return userAuthorization.Id, nil
}

func GetPlanningById(authorizationId uint) (calendarPlanning model.CalendarPlanning, err error) {
	err = database.Get(&calendarPlanning, `SELECT *  FROM calendar_planning WHERE calendar_planning.authorization_id = ?`, authorizationId)
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

func GetPlanningActorByCalendarId(calendarId uint) (calendarPlanningActors []model.User, err error) {
	err = database.GetMany(&calendarPlanningActors,
		`SELECT user.* FROM user
              JOIN authorization ON user.id = authorization.user_id
              JOIN calendar_planning_actor ON authorization.id = calendar_planning_actor.authorization_id
              JOIN calendar_planning ON calendar_planning_actor.calendar_planning_id = calendar_planning.id
     WHERE calendar_planning.id = ?`, calendarId)
	if err != nil {
		return calendarPlanningActors, err
	}
	return calendarPlanningActors, err
}

func GetSelectedPlanningActor(userId uint, calendarPlanningId uint) (calendarPlanningActor model.CalendarPlanningActor, err error) {
	err = database.Get(&calendarPlanningActor,
		`SELECT calendar_planning_actor.*  FROM calendar_planning_actor
                                  JOIN authorization ON calendar_planning_actor.authorization_id = authorization.id
                                  JOIN calendar_planning ON calendar_planning_actor.calendar_planning_id = calendar_planning.id
                                  JOIN user ON authorization.user_id = user.id
                                  WHERE user.id= ? AND calendar_planning.id = ?`, userId, calendarPlanningId)
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
