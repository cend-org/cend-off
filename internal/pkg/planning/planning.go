package planning

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
)

func CreateUserPlannings(ctx context.Context, input *model.CalendarPlanningInput) (*model.CalendarPlanning, error) {
	panic(fmt.Errorf("not implemented: CreateUserPlannings - createUserPlannings"))
}

func AddUserIntoPlanning(ctx context.Context, calendarID int, selectedUserID int) (*model.CalendarPlanningActor, error) {
	panic(fmt.Errorf("not implemented: AddUserIntoPlanning - addUserIntoPlanning"))
}

func GetUserPlannings(ctx context.Context) (*model.CalendarPlanning, error) {
	panic(fmt.Errorf("not implemented: GetUserPlannings - getUserPlannings"))
}

func GetPlanningActors(ctx context.Context, calendarID int) ([]model.User, error) {
	panic(fmt.Errorf("not implemented: GetPlanningActors - getPlanningActors"))
}

//removeUserFromPlanning(calendarPlanningId: ID!, selectedUserId: ID!): String
