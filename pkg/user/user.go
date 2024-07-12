package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
)

const (
	StatusNeedPassword = 0
	StatusNeedProfile  = 1
	StatusActive       = 2
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

func UpdateProfileAndPassword(userId int, new model.UserInput, newPass model.PasswordInput) (usr *model.User, err error) {
	var user model.User

	if !utils.PasswordHasValidLength(*newPass.Hash) {
		return nil, errx.PasswordLengthError
	}

	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return nil, errx.DbGetError
	}

	user = model.MapUserInputToUser(new, user)

	_, err = NewPassword(user.Id, newPass)
	if err != nil {
		return nil, err
	}

	err = database.Update(user)
	if err != nil {
		return nil, errx.SupportError
	}

	return &user, err
}

func UpdMyProfile(userId int, new model.UserInput) (usr *model.User, err error) {
	var user model.User

	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return nil, err
	}

	user = model.MapUserInputToUser(new, user)

	err = database.Update(user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func MyProfile(userId int) (usr *model.User, err error) {
	var user model.User
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return nil, errx.UnknownUserError
	}
	return &user, nil
}

func getUserByEmail(email string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}
	return user, err
}

func CreateStudentPassword(studentId int) (string, error) {
	var (
		password string
		student  model.User
		err      error
	)

	student, err = GetUserWithId(studentId)
	if err != nil {
		return password, err
	}

	password = createPassword(student.Name, student.FamilyName)
	if err != nil {
		return password, errx.PasswordInsertError
	}

	passwordInput := model.PasswordInput{
		Hash: &password,
	}

	_, err = NewPassword(studentId, passwordInput)
	if err != nil {
		return password, errx.PasswordInsertError
	}
	return password, nil
}

func createPassword(name, familyName string) string {
	label := fmt.Sprintf("%s%s", familyName, name)

	hashed := make([]byte, len(label))
	for i, char := range label {
		switch {
		case 'a' <= char && char <= 'z':
			hashed[i] = byte((char-'a'+13)%26 + 'a')
		case 'A' <= char && char <= 'Z':
			hashed[i] = byte((char-'A'+13)%26 + 'A')
		default:
			hashed[i] = byte(char)
		}
	}
	return string(hashed)
}

func GetUserWithId(userId int) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return user, err
	}
	return user, nil
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
