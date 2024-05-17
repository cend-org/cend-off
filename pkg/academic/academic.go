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

func GetTutorWithPreferredCourse(studentId int) (user model.User, err error) {
	var (
		course model.AcademicCourse
	)

	course, err = GetUserPreferredCourse(studentId)
	if err != nil {
		return user, err
	}

	user, err = GetTutorByCourseId(course.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

/*
UTILS
*/

func GetUserPreferredCourse(userId int) (course model.AcademicCourse, err error) {
	err = database.Get(&course, `SELECT ac.* FROM academic_course ac 
    	JOIN  user_academic_course uac ON ac.id = uac.course_id
    	WHERE  uac.user_id = ?`, userId)
	return course, nil
}

func GetTutorByCourseId(courseId int) (user model.User, err error) {
	err = database.Get(&user, `SELECT u.* FROM user u JOIN user_academic_course uac ON u.id = uac.user_id WHERE uac.course_id = ? `, courseId)
	return user, nil
}
