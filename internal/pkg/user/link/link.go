package link

import (
	"context"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
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

// Parent Handler```

func AddParentToUser(ctx *context.Context, input *model.UpdateUser) (*model.User, error) {
	var (
		tok                     *authentication.Token
		parent                  model.User
		currentParent           model.User
		userAuthorizationLinkId uint
		err                     error
	)

	// Select User
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &currentParent, errx.UnAuthorizedError
	}

	//check if user is a student
	if !authorization.IsUserStudent(tok.UserId) {
		return &currentParent, errx.Lambda(errors.New("user is not a student"))
	}

	parent = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	// Check if parent  doesn't exist in the database based on name and family name then  create a user named parent if not
	currentParent, err = GetUserByUserName(parent)
	if currentParent.Id == state.ZERO {
		//	Create parent with email parent+1@cend.intra
		currentParent, err = CreateNewUser(parent, ParentAuthorizationLevel)
		if err != nil {
			return &currentParent, errx.DbInsertError
		}

	}

	//Check if link already exist if not then create new link and add creator into link actor by default
	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		return &currentParent, errx.DbGetError
	}

	userAuthorizationLinkId, err = GetUserLink(StudentParent, auth.Id)
	if userAuthorizationLinkId != state.ZERO {
		//Check if parent is already added to the user
		currentParentAuth, err := authorization.GetUserAuthorization(currentParent.Id, ParentAuthorizationLevel)
		if err != nil {
			return &currentParent, errx.DbGetError
		}

		parents, err := GetLink(currentParentAuth.Id, ParentAuthorizationLevel, StudentParent)
		if len(parents) > 0 {
			return &currentParent, errx.DuplicateUserError
		}

	}

	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(StudentParent, tok.UserId, tok.UserLevel)
		if err != nil {
			return &currentParent, errx.DbInsertError
		}
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, currentParent.Id, ParentAuthorizationLevel)
	if err != nil {
		return &currentParent, errx.DbInsertError
	}

	return &currentParent, nil
}

func GetUserParent(ctx *context.Context) ([]*model.User, error) {
	var (
		tok        *authentication.Token
		parents    []model.User
		gqlParents []*model.User
		err        error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return gqlParents, errx.ParseError
	}

	if tok.UserId == state.ZERO {
		if err != nil {
			return gqlParents, errx.UnAuthorizedError
		}
	}
	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			return gqlParents, errx.DbGetError
		}
	}

	parents, err = GetLink(auth.Id, ParentAuthorizationLevel, StudentParent)
	if err != nil {
		return gqlParents, errx.DbGetError
	}

	for _, parent := range parents {
		gqlParents = append(gqlParents, &model.User{
			Id:                    parent.Id,
			CreatedAt:             parent.CreatedAt,
			UpdatedAt:             parent.UpdatedAt,
			DeletedAt:             parent.DeletedAt,
			Name:                  parent.Name,
			FamilyName:            parent.FamilyName,
			NickName:              parent.NickName,
			Email:                 parent.Email,
			Matricule:             parent.Matricule,
			Age:                   parent.Age,
			BirthDate:             parent.BirthDate,
			Sex:                   parent.Sex,
			Lang:                  parent.Lang,
			Status:                parent.Status,
			ProfileImageXid:       parent.ProfileImageXid,
			Description:           parent.Description,
			CoverText:             parent.CoverText,
			Profile:               parent.Profile,
			ExperienceDetail:      parent.ExperienceDetail,
			AdditionalDescription: parent.AdditionalDescription,
			AddOnTitle:            parent.AddOnTitle,
		})
	}

	return gqlParents, nil

}

func RemoveUserParent(ctx *context.Context, input *model.UpdateUser) (*string, error) {
	var (
		parent model.User
		actor  model.UserAuthorizationLinkActor
		tok    *authentication.Token
		err    error
		status string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}

	parent = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	actor, err = GetSelectedUserLinkActor(parent, StudentParent)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = DeleteUserLinkActor(actor)
	if err != nil {
		return &status, errx.DbDeleteError
	}
	status = "success"
	return &status, nil
}

// Tutor Handler

func AddTutorToUser(ctx *context.Context, input *model.UpdateUser) (*model.User, error) {
	var (
		tok                     *authentication.Token
		tutor                   model.User
		currentTutor            model.User
		userAuthorizationLinkId uint
		err                     error
	)

	// Select User
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &currentTutor, errx.UnAuthorizedError

	}

	if !authorization.IsUserStudent(tok.UserId) {
		return &currentTutor, errx.Lambda(errors.New("user is not a student"))

	}

	tutor = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	currentTutor, err = GetUserByUserName(tutor)
	if currentTutor.Id == state.ZERO {
		//	Create tutor with email tutor+1@cend.intra
		currentTutor, err = CreateNewUser(tutor, TutorAuthorizationLevel)
		if err != nil {
			return &currentTutor, errx.DbInsertError

		}

	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		return &currentTutor, errx.DbGetError

	}

	userAuthorizationLinkId, err = GetUserLink(StudentTutor, auth.Id)
	if userAuthorizationLinkId != state.ZERO {
		//Check if tutor is already added to the user
		currentTutorAuth, err := authorization.GetUserAuthorization(currentTutor.Id, TutorAuthorizationLevel)
		if err != nil {
			return &currentTutor, errx.DbGetError

		}
		tutors, err := GetLink(currentTutorAuth.Id, TutorAuthorizationLevel, StudentTutor)
		if len(tutors) > state.ZERO {
			return &currentTutor, errx.DuplicateUserError

		}
	}

	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(StudentTutor, tok.UserId, tok.UserLevel)
		if err != nil {
			return &currentTutor, errx.DbInsertError

		}
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, currentTutor.Id, TutorAuthorizationLevel)
	if err != nil {
		return &currentTutor, errx.DbInsertError

	}

	return &currentTutor, nil
}

func GetUserTutor(ctx *context.Context) ([]*model.User, error) {
	var (
		tok       *authentication.Token
		tutors    []model.User
		gqlTutors []*model.User
		err       error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return gqlTutors, errx.ParseError
	}

	if tok.UserId == state.ZERO {
		if err != nil {
			return gqlTutors, errx.UnAuthorizedError
		}
	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			return gqlTutors, errx.DbGetError

		}
	}

	tutors, err = GetLink(auth.Id, TutorAuthorizationLevel, StudentTutor)
	if err != nil {
		return gqlTutors, errx.DbGetError

	}

	for _, tutor := range tutors {
		gqlTutors = append(gqlTutors, &model.User{
			Id:                    tutor.Id,
			CreatedAt:             tutor.CreatedAt,
			UpdatedAt:             tutor.UpdatedAt,
			DeletedAt:             tutor.DeletedAt,
			Name:                  tutor.Name,
			FamilyName:            tutor.FamilyName,
			NickName:              tutor.NickName,
			Email:                 tutor.Email,
			Matricule:             tutor.Matricule,
			Age:                   tutor.Age,
			BirthDate:             tutor.BirthDate,
			Sex:                   tutor.Sex,
			Lang:                  tutor.Lang,
			Status:                tutor.Status,
			ProfileImageXid:       tutor.ProfileImageXid,
			Description:           tutor.Description,
			CoverText:             tutor.CoverText,
			Profile:               tutor.Profile,
			ExperienceDetail:      tutor.ExperienceDetail,
			AdditionalDescription: tutor.AdditionalDescription,
			AddOnTitle:            tutor.AddOnTitle,
		})
	}

	return gqlTutors, nil
}

func RemoveUserTutor(ctx *context.Context, input *model.UpdateUser) (*string, error) {
	var (
		tutor  model.User
		actor  model.UserAuthorizationLinkActor
		tok    *authentication.Token
		err    error
		status string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError

	}

	tutor = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	actor, err = GetSelectedUserLinkActor(tutor, StudentTutor)
	if err != nil {
		return &status, errx.DbGetError

	}

	err = DeleteUserLinkActor(actor)
	if err != nil {
		return &status, errx.DbDeleteError

	}
	status = "success"
	return &status, nil

}

// Professor Handler

func AddProfessorToUser(ctx *context.Context, input *model.UpdateUser) (*model.User, error) {
	var (
		tok                     *authentication.Token
		professor               model.User
		currentProfessor        model.User
		userAuthorizationLinkId uint
		err                     error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &currentProfessor, errx.UnAuthorizedError

	}

	if !authorization.IsUserStudent(tok.UserId) {
		return &currentProfessor, errx.Lambda(errors.New("user is not a student"))

	}

	professor = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}
	currentProfessor, err = GetUserByUserName(professor)
	if currentProfessor.Id == state.ZERO {
		//	Create professor with email professor+1@cend.intra
		currentProfessor, err = CreateNewUser(professor, ProfessorAuthorizationLevel)
		if err != nil {
			return &currentProfessor, errx.DbInsertError

		}

	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		return &currentProfessor, errx.DbGetError

	}

	userAuthorizationLinkId, err = GetUserLink(StudentProfessor, auth.Id)
	if userAuthorizationLinkId != state.ZERO {
		//Check if tutor is already added to the user
		currentTutorAuth, err := authorization.GetUserAuthorization(currentProfessor.Id, ProfessorAuthorizationLevel)
		if err != nil {
			return &currentProfessor, errx.DbGetError

		}
		professors, err := GetLink(currentTutorAuth.Id, ProfessorAuthorizationLevel, StudentProfessor)
		if len(professors) > state.ZERO {
			return &currentProfessor, errx.DuplicateUserError

		}
	}
	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(StudentProfessor, tok.UserId, tok.UserLevel)
		if err != nil {
			return &currentProfessor, errx.DbInsertError

		}
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, currentProfessor.Id, ProfessorAuthorizationLevel)
	if err != nil {
		return &currentProfessor, errx.DbInsertError
	}

	return &currentProfessor, nil
}

func GetUserProfessor(ctx *context.Context) ([]*model.User, error) {
	var (
		tok           *authentication.Token
		professors    []model.User
		gqlProfessors []*model.User
		err           error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return gqlProfessors, errx.ParseError

	}

	if tok.UserId == state.ZERO {
		if err != nil {
			return gqlProfessors, errx.UnAuthorizedError

		}
	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			return gqlProfessors, errx.DbGetError

		}
	}

	professors, err = GetLink(auth.Id, ProfessorAuthorizationLevel, StudentProfessor)
	if err != nil {

		return gqlProfessors, errx.DbGetError

	}

	for _, professor := range professors {
		gqlProfessors = append(gqlProfessors, &model.User{
			Id:                    professor.Id,
			CreatedAt:             professor.CreatedAt,
			UpdatedAt:             professor.UpdatedAt,
			DeletedAt:             professor.DeletedAt,
			Name:                  professor.Name,
			FamilyName:            professor.FamilyName,
			NickName:              professor.NickName,
			Email:                 professor.Email,
			Matricule:             professor.Matricule,
			Age:                   professor.Age,
			BirthDate:             professor.BirthDate,
			Sex:                   professor.Sex,
			Lang:                  professor.Lang,
			Status:                professor.Status,
			ProfileImageXid:       professor.ProfileImageXid,
			Description:           professor.Description,
			CoverText:             professor.CoverText,
			Profile:               professor.Profile,
			ExperienceDetail:      professor.ExperienceDetail,
			AdditionalDescription: professor.AdditionalDescription,
			AddOnTitle:            professor.AddOnTitle,
		})
	}

	return gqlProfessors, nil
}

func RemoveUserProfessor(ctx *context.Context, input *model.UpdateUser) (*string, error) {
	var (
		professor model.User
		actor     model.UserAuthorizationLinkActor
		tok       *authentication.Token
		err       error
		status    string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError

	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError

	}

	professor = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	actor, err = GetSelectedUserLinkActor(professor, StudentProfessor)
	if err != nil {
		return &status, errx.DbGetError

	}

	err = DeleteUserLinkActor(actor)
	if err != nil {
		return &status, errx.DbDeleteError

	}
	status = "success"
	return &status, nil

}

//Student Handler

func AddStudentToLink(ctx *context.Context, input *model.UpdateUser) (*model.User, error) {
	var (
		tok            *authentication.Token
		err            error
		student        model.User
		currentStudent model.User
		linkType       uint
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &currentStudent, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserId) {
		return &currentStudent, errx.Lambda(errors.New("current user is a student"))
	}

	if authorization.IsUserParent(tok.UserId) {
		linkType = StudentParent
	}

	if authorization.IsUserTutor(tok.UserId) {
		linkType = StudentTutor
	}

	if authorization.IsUserProfessor(tok.UserId) {
		linkType = StudentProfessor
	}

	student = model.User{
		Id:         input.Id,
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	currentStudent, err = GetUserByUserName(student)
	if err != nil {
		return &currentStudent, errx.DbGetError
	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		return &currentStudent, errx.UnAuthorizedError
	}

	userAuthorizationLinkId, err := GetUserLink(linkType, auth.Id)
	if userAuthorizationLinkId != state.ZERO {
		//Check if student is already added to the user
		currentStudentAuth, err := authorization.GetUserAuthorization(currentStudent.Id, StudentAuthorizationLevel)
		if err != nil {
			return &currentStudent, errx.DbGetError
		}

		links, err := GetLink(currentStudentAuth.Id, StudentAuthorizationLevel, linkType)
		if len(links) > 0 {
			return &currentStudent, errx.DuplicateUserError
		}

	}
	if userAuthorizationLinkId == state.ZERO {
		userAuthorizationLinkId, err = SetUserAuthorizationLink(linkType, tok.UserId, tok.UserLevel)
		if err != nil {
			return &currentStudent, errx.DbInsertError
		}
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, currentStudent.Id, StudentAuthorizationLevel)
	if err != nil {
		return &currentStudent, errx.DbInsertError
	}

	return &currentStudent, nil
}

func GetStudent(ctx *context.Context) ([]*model.User, error) {
	var (
		tok         *authentication.Token
		students    []model.User
		gqlStudents []*model.User
		err         error
		linkType    uint
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return gqlStudents, errx.ParseError
	}

	if tok.UserId == state.ZERO {
		if err != nil {
			return gqlStudents, errx.UnAuthorizedError
		}
	}

	auth, err := authorization.GetUserAuthorization(tok.UserId, tok.UserLevel)
	if err != nil {
		if err != nil {
			return gqlStudents, errx.DbGetError
		}
	}

	if authorization.IsUserParent(tok.UserId) {
		linkType = StudentParent
	}

	if authorization.IsUserTutor(tok.UserId) {
		linkType = StudentTutor
	}

	if authorization.IsUserProfessor(tok.UserId) {
		linkType = StudentProfessor
	}

	students, err = GetLink(auth.Id, StudentAuthorizationLevel, linkType)
	if err != nil {
		return gqlStudents, errx.DbGetError
	}

	for _, student := range students {
		gqlStudents = append(gqlStudents, &model.User{
			Id:                    student.Id,
			CreatedAt:             student.CreatedAt,
			UpdatedAt:             student.UpdatedAt,
			DeletedAt:             student.DeletedAt,
			Name:                  student.Name,
			FamilyName:            student.FamilyName,
			NickName:              student.NickName,
			Email:                 student.Email,
			Matricule:             student.Matricule,
			Age:                   student.Age,
			BirthDate:             student.BirthDate,
			Sex:                   student.Sex,
			Lang:                  student.Lang,
			Status:                student.Status,
			ProfileImageXid:       student.ProfileImageXid,
			Description:           student.Description,
			CoverText:             student.CoverText,
			Profile:               student.Profile,
			ExperienceDetail:      student.ExperienceDetail,
			AdditionalDescription: student.AdditionalDescription,
			AddOnTitle:            student.AddOnTitle,
		})
	}

	return gqlStudents, nil
}

func RemoveStudent(ctx *context.Context, input *model.UpdateUser) (*string, error) {
	var (
		student  model.User
		actor    model.UserAuthorizationLinkActor
		tok      *authentication.Token
		err      error
		linkType uint
		status   string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}

	student = model.User{
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		Email:      input.Email,
	}

	if authorization.IsUserParent(tok.UserId) {
		linkType = StudentParent
	}

	if authorization.IsUserTutor(tok.UserId) {
		linkType = StudentTutor
	}

	if authorization.IsUserProfessor(tok.UserId) {
		linkType = StudentProfessor
	}

	actor, err = GetSelectedUserLinkActor(student, linkType)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = DeleteUserLinkActor(actor)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	status = "success"
	return &status, nil
}

/*
	UTILS
*/

func SetUserAuthorizationLink(linkType uint, userId uint, userLevel uint) (userAuthorizationLinkId uint, err error) {
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

func SetUserAuthorizationLinkActor(linkId uint, userId uint, level uint) (err error) {
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

func CreateNewUser(user model.User, authLevel uint) (currentUser model.User, err error) {
	if authLevel == ParentAuthorizationLevel {
		user.Email = "parent+1@cend.intern"
	}

	if authLevel == TutorAuthorizationLevel {
		user.Email = "tutor+1@cend.intern"
	}

	if authLevel == ProfessorAuthorizationLevel {
		user.Email = "professor+1@cend.intern"
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return user, err
	}

	if user.Name == state.EMPTY {
		user.Name = user.Matricule
	}

	if user.NickName == state.EMPTY {
		user.NickName = user.Matricule
	}

	user.Id, err = database.InsertOne(user)
	if err != nil {
		return user, err
	}

	err = authorization.NewUserAuthorization(user.Id, authLevel)
	if err != nil {
		return user, err
	}

	_, err = authentication.GetTokenString(user.Id)
	if err != nil {
		return user, err
	}

	currentUser = user
	return currentUser, nil
}

func GetLink(authId uint, authorizationLevel uint, linkType uint) (link []model.User, err error) {
	err = database.GetMany(&link, `SELECT user.* FROM user
                       JOIN authorization ON user.id = authorization.user_id
                       JOIN user_authorization_link_actor ON authorization.id = user_authorization_link_actor.authorization_id
                       JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
WHERE user_authorization_link.id =  (
    SELECT user_authorization_link_actor.user_authorization_link_id
    FROM user_authorization_link_actor
             JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
    WHERE user_authorization_link_actor.authorization_id = ? AND user_authorization_link.link_type = ?
    )
AND authorization.level = ?`, authId, linkType, authorizationLevel)
	if err != nil {
		return link, err
	}

	return link, nil
}

func GetUserByUserName(currentUser model.User) (user model.User, err error) {
	err = database.Get(&user, `SELECT user.* FROM user WHERE user.name = ? and user.family_name = ?`, currentUser.Name, currentUser.FamilyName)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserLink(linkType uint, authorizationId uint) (linkId uint, err error) {
	var userLink model.UserAuthorizationLink

	err = database.Get(&userLink,
		`SELECT user_authorization_link.* FROM user_authorization_link
                                  JOIN user_authorization_link_actor ON user_authorization_link.id = user_authorization_link_actor.user_authorization_link_id
                                  WHERE user_authorization_link.link_type = ? AND user_authorization_link_actor.authorization_id = ?`, linkType, authorizationId)
	if err != nil {
		return 0, err
	}

	return userLink.Id, nil
}

func GetSelectedUserLinkActor(user model.User, linkType uint) (actor model.UserAuthorizationLinkActor, err error) {
	err = database.Get(&actor,
		`SELECT user_authorization_link_actor.*
FROM user_authorization_link_actor
JOIN user_authorization_link ON user_authorization_link_actor.user_authorization_link_id = user_authorization_link.id
JOIN authorization ON user_authorization_link_actor.authorization_id = authorization.id
JOIN user ON authorization.user_id = user.id
WHERE user.family_name = ? AND  user.name = ? AND user_authorization_link.link_type = ?`, user.FamilyName, user.Name, linkType)
	if err != nil {
		return actor, err
	}

	return actor, nil
}

func DeleteUserLinkActor(userAuthorizationLinkActor model.UserAuthorizationLinkActor) (err error) {
	err = database.Delete(userAuthorizationLinkActor)
	if err != nil {
		return err
	}

	return nil
}
