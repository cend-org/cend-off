package model

import (
	"github.com/cend-org/duval/internal/database"
	"time"
)

type UserCollector struct{}

func (c *UserCollector) Name(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE name = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) FamilyName(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE family_name = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) NickName(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE nick_name = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) Email(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE email = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) BirthDate(a *time.Time) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE birth_date = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) Sex(a *int) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE sex = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) Lang(a *int) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE lang = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) Description(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE description = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) CoverText(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE cover_text = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) Profile(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE profile = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) ExperienceDetail(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE experience_detail = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) AdditionalDescription(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE additional_description = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) AddOnTitle(a *string) (r User, err error) {
	err = database.Get(&r, `SELECT * FROM user WHERE add_on_title = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserCollector) _b() {
	_ = time.Now()
}

func MapUserInputToUser(a UserInput, e User) User {
	if a.Name != nil {
		e.Name = *a.Name
	}
	if a.FamilyName != nil {
		e.FamilyName = *a.FamilyName
	}
	if a.NickName != nil {
		e.NickName = *a.NickName
	}
	if a.Email != nil {
		e.Email = *a.Email
	}
	if a.BirthDate != nil {
		e.BirthDate = *a.BirthDate
	}
	if a.Sex != nil {
		e.Sex = *a.Sex
	}
	if a.Lang != nil {
		e.Lang = *a.Lang
	}
	if a.Description != nil {
		e.Description = *a.Description
	}
	if a.CoverText != nil {
		e.CoverText = *a.CoverText
	}
	if a.Profile != nil {
		e.Profile = *a.Profile
	}
	if a.ExperienceDetail != nil {
		e.ExperienceDetail = *a.ExperienceDetail
	}
	if a.AdditionalDescription != nil {
		e.AdditionalDescription = *a.AdditionalDescription
	}
	if a.AddOnTitle != nil {
		e.AddOnTitle = *a.AddOnTitle
	}
	return e
}
