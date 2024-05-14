package academic

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/xorcare/pointer"
)

func GetAcademicLevels() (academics []model.AcademicLevel, err error) {
	err = database.Select(&academics, `SELECT * FROM academic_level ORDER BY created_at`)
	if err != nil {
		return nil, err
	}

	return academics, err
}

func GetAcademicCourses(academicId int) (courses []model.AcademicCourse, err error) {
	err = database.Select(&courses, `SELECT * FROM academic_course WHERE academic_level_id = ?`, academicId)
	if err != nil {
		return nil, err
	}
	return courses, err
}

func NewUserAcademicCourses(userId int, new []*model.UserAcademicCourseInput) (ret *bool, err error) {
	err = database.Exec(`DELETE FROM user_academic_course WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(new); i++ {
		courseInput := new[i]
		if courseInput != nil {
			course := model.MapUserAcademicCourseInputToUserAcademicCourse(*courseInput, model.UserAcademicCourse{})
			course.UserId = userId

			_, err = database.Insert(course)
			if err != nil {
				return nil, err
			}
		}
	}

	return pointer.Bool(true), err
}
