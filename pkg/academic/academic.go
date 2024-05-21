package academic

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/pkg/user/link"
	"github.com/xorcare/pointer"
)

func GetAcademicLevels() (academics []model.AcademicLevel, err error) {
	err = database.Select(&academics, `SELECT * FROM academic_level ORDER BY created_at`)
	if err != nil {
		return nil, errx.SupportError
	}

	return academics, err
}

func GetAcademicCourses(academicId int) (courses []model.AcademicCourse, err error) {
	err = database.Select(&courses, `SELECT * FROM academic_course WHERE academic_level_id = ?`, academicId)
	if err != nil {
		return nil, errx.SupportError
	}
	return courses, err
}

func NewUserAcademicCourses(userId int, new []*model.UserAcademicCourseInput) (ret *bool, err error) {
	var preference model.UserAcademicCoursePreference

	err = database.Exec(`DELETE FROM user_academic_course WHERE user_id = ?`, userId)
	if err != nil {
		return nil, errx.SupportError
	}

	for i := 0; i < len(new); i++ {
		courseInput := new[i]
		if courseInput != nil {
			course := model.MapUserAcademicCourseInputToUserAcademicCourse(*courseInput, model.UserAcademicCourse{})
			course.UserId = userId

			_, err = database.Insert(course)
			if err != nil {
				return nil, errx.SupportError
			}
		}
	}
	preference.UserId = userId
	_, err = database.Insert(preference)
	if err != nil {
		return nil, errx.SupportError
	}

	return pointer.Bool(true), err
}

func GetTutorWithPreferredCourse(studentId int) (user model.User, err error) {
	var (
		course []model.AcademicCourse
	)

	course, err = GetUserPreferredCourse(studentId)
	if err != nil {
		return user, errx.CoursePreferenceError
	}

	user, err = GetTutorByCourse(course)
	if err != nil {
		return user, errx.SupportError
	}

	return user, nil
}

func SetUserAcademicLevel(userId, academicLevelId int) (err error) {
	var (
		academicCourse model.UserAcademicCourse
		courses        []model.AcademicCourse
	)

	courses, err = GetAcademicCourses(academicLevelId)
	if err != nil {
		return errx.LevelError
	}

	academicCourse.CourseId = courses[0].Id
	academicCourse.UserId = userId

	_, err = database.InsertOne(academicCourse)
	if err != nil {
		return errx.SupportError
	}

	return nil
}

func SetUserAcademicLevels(userId int, academicLevelIds []*int) (err error) {
	var academicCourse model.UserAcademicCourse

	for _, academicLevelId := range academicLevelIds {
		courses, err := GetAcademicCourses(*academicLevelId)
		if err != nil {
			return errx.LevelError
		}

		academicCourse = model.UserAcademicCourse{
			CourseId: courses[0].Id,
			UserId:   userId,
		}
		_, err = database.InsertOne(academicCourse)
		if err != nil {
			return errx.SupportError
		}
	}

	return nil
}

func GetUserAcademicLevels(userId int) (academicLevel []model.AcademicLevel, err error) {
	academicLevel, err = GetUserLevel(userId)
	if err != nil {
		return academicLevel, errx.UnknownLevelError
	}
	return academicLevel, nil
}

func UpdStudentAcademicCoursesPreferenceByParent(studentId int, new model.UserAcademicCoursePreferenceInput) (preference model.UserAcademicCoursePreference, err error) {
	preference, err = GetPreferences(studentId)
	if err != nil {
		return preference, errx.SupportError
	}

	preference = model.MapUserAcademicCoursePreferenceInputToUserAcademicCoursePreference(new, preference)

	err = database.Update(preference)
	if err != nil {
		return preference, errx.SupportError
	}
	return preference, nil
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

func GetTutorByCourse(courses []model.AcademicCourse) (topTutor model.User, err error) {
	var courseTutors []model.User
	tutorCounts := make(map[model.User]int)

	for _, course := range courses {

		err = database.Select(&courseTutors, `
            SELECT u.*
            FROM user u
            JOIN user_academic_course uac ON u.id = uac.user_id
            JOIN academic_course ac ON uac.course_id = ac.id
            JOIN authorization a ON u.id = a.user_id
            WHERE ac.name = ? AND a.level = ?`, course.Name, link.TutorAuthorizationLevel)
		if err != nil {
			return model.User{}, err
		}
		for _, tutor := range courseTutors {
			tutorCounts[tutor]++
		}
	}

	maxCount := 0
	for tutor, count := range tutorCounts {
		if count > maxCount {
			topTutor = tutor
			maxCount = count
		}
	}

	return topTutor, nil
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

func GetPreferences(userId int) (userAcademicCoursePreferences model.UserAcademicCoursePreference, err error) {
	err = database.Get(&userAcademicCoursePreferences,
		`SELECT * FROM user_academic_course_preference WHERE user_id = ?`, userId)
	if err != nil {
		return userAcademicCoursePreferences, err
	}
	return userAcademicCoursePreferences, nil
}
