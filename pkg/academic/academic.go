package academic

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/user"
	"github.com/cend-org/duval/pkg/user/authorization"
	"github.com/xorcare/pointer"
)

const (
	// AuthorizationLevel

	StudentAuthorizationLevel   = 0
	ParentAuthorizationLevel    = 1
	TutorAuthorizationLevel     = 2
	ProfessorAuthorizationLevel = 3

	//Link_type

	StudentParent    = 0
	StudentTutor     = 1
	StudentProfessor = 2
)

/*

	LEVELS

*/

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

func GetAcademicLevels() (academics []model.AcademicLevel, err error) {
	err = database.Select(&academics, `SELECT * FROM academic_level ORDER BY created_at`)
	if err != nil {
		return nil, errx.SupportError
	}

	return academics, err
}

/*

	COURSE

*/

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

func GetAcademicCourses(academicId int) (courses []model.AcademicCourse, err error) {
	err = database.Select(&courses, `SELECT * FROM academic_course WHERE academic_level_id = ?`, academicId)
	if err != nil {
		return nil, errx.SupportError
	}
	return courses, err
}

/*

	LINK

*/

// STUDENT

func AddStudent(linkType, userLinkId, level int, email string) (studentId int, err error) {
	var (
		userAuthorizationLinkId int
		student                 model.User
	)

	student, err = GetUserWithEmail(email)
	if err != nil {
		return studentId, errx.UnknownStudentError
	}

	studentId = student.Id

	auth, err := authorization.GetUserAuthorization(student.Id, StudentAuthorizationLevel)
	if err != nil {
		return studentId, errx.UnknownStudentError
	}

	userAuthorizationLinkId, err = GetUserLink(linkType, auth.Id)
	if userAuthorizationLinkId != state.ZERO {
		currentParentAuth, err := authorization.GetUserAuthorization(userLinkId, level)
		if err != nil {
			return studentId, errx.ParentError
		}

		parents, err := GetLink(currentParentAuth.Id, level, linkType)
		if len(parents) > 0 {
			return studentId, errx.DuplicateUserError
		}
	}

	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(linkType, student.Id, StudentAuthorizationLevel)
		if err != nil {
			return studentId, errx.SupportError
		}
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, userLinkId, level)
	if err != nil {
		return studentId, errx.ParentError
	}

	return studentId, nil
}

func CreateStudent(name, familyName string) (student model.User, err error) {
	var email = fmt.Sprintf("%s.%s@cend.intern", name, familyName)

	_, err = user.NewStudent(email)
	if err != nil {
		return student, err
	}

	student, err = GetUserWithEmail(email)
	if err != nil {
		return student, err
	}

	studentInput := model.UserInput{
		Name:       &name,
		FamilyName: &familyName,
	}

	err = UpdateStudent(student.Id, studentInput)
	if err != nil {
		return student, err
	}

	return student, nil
}

func UpdateStudent(studentId int, profile model.UserInput) (err error) {
	var (
		user model.User
	)
	user, err = GetUserWithId(studentId)
	if err != nil {
		return errx.UnknownStudentError
	}

	user = model.MapUserInputToUser(profile, user)

	err = database.Update(user)
	if err != nil {
		return errx.SupportError
	}

	return nil
}

// TUTOR

func AddStudentToLink(userLinkId int, name, familyName string) (*model.User, error) {
	var (
		err            error
		currentStudent model.User
	)

	currentStudent, err = GetUserByUserName(name, familyName)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		currentStudent, err = CreateStudent(name, familyName)
		if err != nil {
			return nil, errx.SupportError
		}
	}

	_, err = AddStudent(StudentParent, userLinkId, ParentAuthorizationLevel, currentStudent.Email)
	if err != nil {
		return nil, errx.SupportError
	}
	return &currentStudent, nil
}

func AddStudentToTutor(tutorId int, studentId int) (*model.User, error) {
	var (
		err            error
		currentStudent model.User
	)

	currentStudent, err = GetUserWithId(studentId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		if err != nil {
			return nil, errx.UnknownLevelError
		}
	}

	_, err = AddStudent(StudentTutor, tutorId, TutorAuthorizationLevel, currentStudent.Email)
	if err != nil {
		return nil, errx.SupportError
	}
	return &currentStudent, nil
}

func GetOtherTutorWithPreferredCourse(studentId int, tutorId int) (user model.User, err error) {
	var (
		course []model.AcademicCourse
	)

	course, err = GetUserPreferredCourse(studentId)
	if err != nil {
		return user, errx.CoursePreferenceError
	}

	user, err = GetOtherTutorByCourse(course, tutorId)
	if err != nil {
		return user, errx.SupportError
	}

	return user, nil
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

// PARENT

func IsStudentParentLinked(parentId, userId int) bool {
	var userLink model.UserAuthorizationLink
	var actor model.UserAuthorizationLinkActor
	var linkType = StudentParent

	var err error
	userLink, err = GetLinkById(userId, linkType)
	if err != nil {
		return false
	}

	err = database.Get(&actor, `SELECT ua_la.* FROM user_authorization_link_actor ua_la  JOIN  authorization a ON ua_la.authorization_id = a.id WHERE ua_la.user_authorization_link_id = ? AND a.user_id = ?`, userLink.Id, parentId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return true
}

// PROFESSOR

func AddStudentToProfessor(profId int, studentId int) (*model.User, error) {
	var (
		err            error
		currentStudent model.User
	)

	currentStudent, err = GetUserWithId(studentId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		if err != nil {
			return nil, errx.UnknownLevelError
		}
	}

	_, err = AddStudent(StudentProfessor, profId, ProfessorAuthorizationLevel, currentStudent.Email)
	if err != nil {
		return nil, errx.SupportError
	}
	return &currentStudent, nil
}

func GetProfessorStudent(name string) (users []model.User, err error) {
	arg := "%" + name + "%"
	query := fmt.Sprintf(`SELECT *  FROM user  WHERE name LIKE '%s' OR family_name LIKE '%s' `, arg, arg)
	err = database.Select(&users, query)
	if err != nil {
		return users, errx.UnknownUserError
	}
	return users, nil
}

/*


	PREFERENCES

*/

func GetUserPreferredCourse(userId int) (course []model.AcademicCourse, err error) {
	err = database.Select(&course, `SELECT ac.* FROM academic_course ac 
    	JOIN  user_academic_course uac ON ac.id = uac.course_id
    	WHERE  uac.user_id = ?`, userId)
	return course, nil
}

/*

	Appointment

*/

func SetAppointment(userId int, new model.AppointmentInput) (model.Appointment, error) {
	var (
		appointment     model.Appointment
		userAppointment model.UserAppointment
		err             error
	)
	appointment = model.MapAppointmentInputToAppointment(new, appointment)

	appointmentId, err := database.InsertOne(appointment)
	if err != nil {
		return appointment, errx.SupportError
	}

	userAppointment.UserId = userId
	userAppointment.AppointmentId = appointmentId

	_, err = database.InsertOne(userAppointment)
	if err != nil {
		return appointment, errx.SupportError
	}

	return appointment, nil
}

/*

	UTILS

*/

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
            WHERE ac.name = ? AND a.level = ?`, course.Name, TutorAuthorizationLevel)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return topTutor, err
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

func GetOtherTutorByCourse(courses []model.AcademicCourse, tutorId int) (topTutor model.User, err error) {
	var courseTutors []model.User
	tutorCounts := make(map[model.User]int)

	for _, course := range courses {
		err = database.Select(&courseTutors, `
            SELECT u.*
            FROM user u
            JOIN user_academic_course uac ON u.id = uac.user_id
            JOIN academic_course ac ON uac.course_id = ac.id
            JOIN authorization a ON u.id = a.user_id
            WHERE ac.name = ? AND a.level = ? AND u.id NOT IN (?)`, course.Name, TutorAuthorizationLevel, tutorId)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return topTutor, err
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

func GetLinkById(userId int, linkType int) (userLink model.UserAuthorizationLink, err error) {
	err = database.Get(&userLink,
		`SELECT ual.*
FROM user_authorization_link ual
         JOIN user_authorization_link_actor ua_la ON ual.Id = ua_la.user_authorization_link_id
         JOIN authorization a ON ua_la.authorization_id = a.id
WHERE ual.link_type = ?
  AND a.user_id = ?`, linkType, userId)
	if err != nil {
		return userLink, err
	}
	return userLink, nil
}

func GetUserWithEmail(email string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetLink(authId int, authorizationLevel int, linkType int) (link []model.User, err error) {
	err = database.Select(&link,
		`SELECT u.*
			FROM user  u
					 JOIN authorization  a ON u.Id = a.user_id
					 JOIN user_authorization_link_actor  ua_la ON a.Id = ua_la.authorization_id
					 JOIN user_authorization_link  ua ON ua_la.user_authorization_link_id = ua.id
					 JOIN user_authorization_link  ua2 ON ua.id = ua2.id
					JOIN user_authorization_link_actor  ua_la2 ON ua2.id = ua_la2.user_authorization_link_id
			WHERE ua_la2.authorization_id = ?
			  AND ua2.link_type = ?
			  AND a.level = ?`, authId, linkType, authorizationLevel)
	if err != nil {
		return link, err
	}

	return link, nil
}

func SetUserAuthorizationLink(linkType int, userId int, userLevel int) (userAuthorizationLinkId int, err error) {
	var (
		userAuthorizationLink model.UserAuthorizationLink
	)

	userAuthorizationLink.LinkType = linkType

	userAuthorizationLinkId, err = database.InsertOne(userAuthorizationLink)
	if err != nil {
		return userAuthorizationLinkId, err
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, userId, userLevel)
	if err != nil {
		return userAuthorizationLinkId, err
	}

	return userAuthorizationLinkId, nil
}

func GetUserByUserName(name, familyName string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE name = ? and family_name = ?`, name, familyName)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserLink(linkType int, authorizationId int) (linkId int, err error) {
	var userLink model.UserAuthorizationLink

	err = database.Get(&userLink,
		`SELECT ual.* FROM user_authorization_link ual
                                  JOIN user_authorization_link_actor ua_la ON ual.Id = ua_la.user_authorization_link_id
                                  WHERE ual.link_type = ? AND ua_la.authorization_id = ?`, linkType, authorizationId)
	if err != nil {
		return state.ZERO, err
	}

	return userLink.Id, nil
}

func SetUserAuthorizationLinkActor(linkId int, userId int, level int) (err error) {
	var userAuthorizationLinkActor model.UserAuthorizationLinkActor

	auth, err := authorization.GetUserAuthorization(userId, level)
	if err != nil {
		return err
	}

	userAuthorizationLinkActor.AuthorizationId = auth.Id
	userAuthorizationLinkActor.UserAuthorizationLinkId = linkId

	_, err = database.InsertOne(userAuthorizationLinkActor)
	if err != nil {
		return err
	}

	return nil
}

func GetUserWithId(userId int) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return user, err
	}
	return user, nil
}
