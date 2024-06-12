package translator

import (
	"database/sql"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
)

func AddLanguageResource(new model.LanguageResourceInput) (*model.LanguageResource, error) {
	var (
		lang model.LanguageResource
		err  error
	)

	if new.ResourceMessage == nil && new.ResourceRef != nil {
		new.ResourceMessage = new.ResourceRef
	}

	if new.ResourceMessage == nil && new.ResourceRef == nil {
		return &lang, errx.LangError
	}

	lang = model.MapLanguageResourceInputToLanguageResource(new, lang)

	lang.Id, err = SetLang(lang)
	if err != nil && errx.IsDuplicate(err) {
		return &lang, errx.DuplicateError
	} else if err != nil && !errx.IsDuplicate(err) {
		return &lang, errx.SupportError
	}

	return &lang, nil
}

func DeleteLanguageResource(id int) (*bool, error) {
	var (
		lang   model.LanguageResource
		err    error
		status bool
	)

	lang, err = GetLang(id)
	if err != nil {
		return &status, err
	}

	err = RemLang(lang)
	if err != nil {
		return &status, err
	}
	status = true
	return &status, nil
}

func GetLanguageResourceByIdAndRef(ref string, id int) (*model.LanguageResource, error) {
	var (
		lang model.LanguageResource
		err  error
	)

	err = database.Get(&lang, `SELECT * FROM language_resource WHERE resource_ref = ? AND id = ?`, ref, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return &lang, errx.MLangError
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &lang, errx.SupportError
	}

	return &lang, nil
}

func GetAllLanguageResources() ([]model.LanguageResource, error) {
	var (
		lang []model.LanguageResource
		err  error
	)

	err = database.Select(&lang, `SELECT * FROM language_resource`)
	if err != nil {
		return lang, errx.SupportError
	}

	return lang, nil
}

/*

	Utils

*/

func SetLang(lang model.LanguageResource) (id int, err error) {
	id, err = database.InsertOne(lang)
	if err != nil {
		return id, err
	}

	return id, nil
}

func GetLang(id int) (lang model.LanguageResource, err error) {
	err = database.Get(&lang, `SELECT * FROM language_resource WHERE id = ?`, id)
	if err != nil {
		return lang, err
	}

	return lang, nil
}

func RemLang(lang model.LanguageResource) (err error) {
	err = database.Delete(lang)
	if err != nil {
		return err
	}
	return nil
}
