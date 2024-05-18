package academic

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/pkg/user/authorization"
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
	var coursePreference model.UserAcademicCoursePreference

	err = database.Exec(`DELETE FROM user_academic_course WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(new); i++ {
		courseInput := new[i]
		if courseInput != nil {
			course := model.MapUserAcademicCourseInputToUserAcademicCourse(*courseInput, model.UserAcademicCourse{})
			course.UserId = userId

			userAcademicCourseId, err := database.Insert(course)
			if err != nil {
				return nil, errx.DbInsertError
			}
			coursePreference.UserAcademicCourseId = int(userAcademicCourseId)

			_, err = database.Insert(coursePreference)
			if err != nil {
				return nil, errx.DbInsertError
			}

		}
	}

	return pointer.Bool(true), err
}

func GetTutorWithPreferredCourse(studentId int) (user model.User, err error) {
	var (
		course []model.AcademicCourse
	)

	course, err = GetUserPreferredCourse(studentId)
	if err != nil {
		return user, err
	}

	user, err = GetTutorByCourse(course)
	if err != nil {
		return user, err
	}

	return user, nil
}

func SetUserAcademicLevel(parentId, studentId, academicLevelId int) (err error) {
	var (
		academicCourse model.UserAcademicCourse
		courses        []model.AcademicCourse
	)

	if !authorization.IsUserParent(parentId) {
		return errx.UnAuthorizedError
	}

	courses, err = GetAcademicCourses(academicLevelId)
	if err != nil {
		return errx.DbGetError
	}

	academicCourse.CourseId = courses[0].Id
	academicCourse.UserId = studentId

	_, err = database.InsertOne(academicCourse)
	if err != nil {
		return errx.DbInsertError
	}

	return nil
}

func SetUserAcademicLevels(userId int, academicLevelIds []*int) (err error) {
	var academicCourse model.UserAcademicCourse

	for _, academicLevelId := range academicLevelIds {
		courses, err := GetAcademicCourses(*academicLevelId)
		if err != nil {
			return errx.DbGetError
		}
		academicCourse = model.UserAcademicCourse{
			CourseId: courses[0].Id,
			UserId:   userId,
		}
		_, err = database.InsertOne(academicCourse)
		if err != nil {
			return errx.DbInsertError
		}
	}

	return nil
}

func GetUserAcademicLevels(userId int) (academicLevel []model.AcademicLevel, err error) {
	academicLevel, err = GetUserLevel(userId)
	if err != nil {
		return academicLevel, errx.DbGetError
	}
	return academicLevel, nil
}

func UpdStudentAcademicCoursesPreferenceByParent(studentId int, new []*model.UserAcademicCoursePreferenceInput) (ret *bool, err error) {
	for i := 0; i < len(new); i++ {
		courseInput := new[i]
		if courseInput != nil {
			preference, err := GetPreferencesById(*courseInput.UserAcademicCourseId)
			if err != nil {
				return nil, errx.DbGetError
			}
			course := model.MapUserAcademicCoursePreferenceInputToUserAcademicCoursePreference(*courseInput, preference)

			err = database.Update(course)
			if err != nil {
				return nil, err
			}
		}
	}

	return pointer.Bool(true), err
}

/*
UTILS
*/

func GetUserPreferredCourse(userId int) (course []model.AcademicCourse, err error) {
	err = database.Select(&course, `SELECT ac.* FROM academic_course ac 
    	JOIN  user_academic_course uac ON ac.id = uac.course_id
    	WHERE  uac.user_id = ?`, userId)
	return course, nil
}

func GetTutorByCourse(courses []model.AcademicCourse) (user model.User, err error) {
	var tutors []model.User
	for _, course := range courses {
		err = database.Select(&tutors, `SELECT u.*
					FROM user u
							 JOIN user_academic_course uac ON u.id = uac.user_id
							 JOIN academic_course ac ON uac.course_id = ac.id
					WHERE ac.name = ? `, course.Name)
	}
	return user, nil
}

func GetUserLevel(UserId int) (academicLevel []model.AcademicLevel, err error) {
	err = database.Select(&academicLevel,
		`SELECT al.* FROM academic_level al
    			JOIN academic_course ac ON al.id = ac.academic_level_id
    			JOIN user_academic_course uac ON uac.course_id = ac.id
			WHERE uac.user_id = ?`, UserId)
	if err != nil {
		return academicLevel, err
	}

	return academicLevel, err
}

func GetPreferences(userId int) (userAcademicCoursePreferences []model.UserAcademicCoursePreference, err error) {
	err = database.Select(&userAcademicCoursePreferences,
		`SELECT uacp.*
		FROM user_academic_course_preference uacp
				 JOIN user_academic_course uac ON uacp.user_academic_course_id = uac.id
		WHERE uac.user_id = ?`, userId)
	if err != nil {
		return userAcademicCoursePreferences, err
	}
	return userAcademicCoursePreferences, nil
}

func GetPreferencesById(coursePreferenceId int) (preference model.UserAcademicCoursePreference, err error) {
	err = database.Get(&preference, `SELECT * FROM user_academic_course_preference WHERE user_academic_course_id = ? `, coursePreferenceId)
	if err != nil {
		return preference, err
	}
	return preference, nil
}
