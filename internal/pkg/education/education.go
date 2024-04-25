package education

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
)

func SetUserEducationLevel(ctx context.Context, input *model.SubjectInput) (*model.Education, error) {
	panic(fmt.Errorf("not implemented: SetUserEducationLevel - setUserEducationLevel"))
}

func UpdateUserEducationLevel(ctx context.Context, input *model.SubjectInput) (*model.Education, error) {
	panic(fmt.Errorf("not implemented: UpdateUserEducationLevel - updateUserEducationLevel"))
}

func GetUserSubjects(ctx context.Context) ([]model.Subject, error) {
	panic(fmt.Errorf("not implemented: GetUserSubjects - getUserSubjects"))
}

func GetEducation(ctx context.Context) ([]model.Education, error) {
	panic(fmt.Errorf("not implemented: GetEducation - getEducation"))
}

func GetUserEducationLevel(ctx context.Context) (*model.Education, error) {
	panic(fmt.Errorf("not implemented: GetUserEducationLevel - getUserEducationLevel"))
}

func GetSchools(ctx context.Context) ([]model.School, error) {
	var schools []model.School
	var err error

	err = database.Select(&schools, `SELECT * FROM school ORDER BY created_at`)
	if err != nil {
		return nil, err
	}

	return schools, err
}

func GetSubjects(ctx context.Context, id int) ([]model.SchoolSubject, error) {
	var subjects []model.SchoolSubject
	var err error

	err = database.Select(&subjects, `SELECT * FROM school_subject WHERE school_number = ?`, id)
	if err != nil {
		return nil, err
	}

	return subjects, err
}

func GetSchool(ctx context.Context, id int) (*model.School, error) {
	var (
		err    error
		school model.School
	)

	err = database.Get(&school, `SELECT * FROM school WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return &school, err
}
