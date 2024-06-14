package translator

import (
	"database/sql"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
)

func AddLanguageResource(new model.LanguageResourceInput) (*string, error) {
	var (
		lang model.LanguageResource
		err  error
	)

	if new.ResourceMessage == nil && new.ResourceRef != nil {
		new.ResourceMessage = new.ResourceRef
	}

	if new.ResourceMessage == nil && new.ResourceRef == nil {
		return &lang.ResourceMessage, errx.LangError
	}

	oldLanguage, err := GetLanguageInfo(*new.ResourceLanguage, *new.ResourceRef)
	if err == nil && oldLanguage.Id > 0 {
		return &lang.ResourceMessage, errx.DuplicateError
	}

	lang = model.MapLanguageResourceInputToLanguageResource(new, lang)

	lang.Id, err = SetLang(lang)
	if err != nil && errx.IsDuplicate(err) {
		return &lang.ResourceMessage, errx.DuplicateError
	} else if err != nil && !errx.IsDuplicate(err) {
		return &lang.ResourceMessage, errx.SupportError
	}

	return &lang.ResourceMessage, nil
}

func AddOrGetResource(new model.LanguageResourceInput) (*string, error) {
	var (
		lang model.LanguageResource
		err  error
	)

	oldLanguage, err := GetLanguageInfo(*new.ResourceLanguage, *new.ResourceRef)
	if err == nil && oldLanguage.Id > 0 {
		return &oldLanguage.ResourceMessage, nil
	}

	if new.ResourceMessage == nil && new.ResourceRef == nil {
		return &lang.ResourceMessage, errx.LangError
	}
	new.ResourceMessage = new.ResourceRef
	lang = model.MapLanguageResourceInputToLanguageResource(new, lang)

	lang.Id, err = SetLang(lang)
	if err != nil && errx.IsDuplicate(err) {
		return &lang.ResourceMessage, errx.DuplicateError
	} else if err != nil && !errx.IsDuplicate(err) {
		return &lang.ResourceMessage, errx.SupportError
	}

	return &lang.ResourceMessage, nil
}

func UpdateLanguageResource(new model.LanguageResourceInput) (*string, error) {
	var (
		lang model.LanguageResource
		err  error
	)

	lang, err = GetLanguageInfo(*new.ResourceLanguage, *new.ResourceRef)
	if err != nil {
		return &lang.ResourceMessage, errx.SupportError
	}

	if new.ResourceMessage == nil {
		new.ResourceMessage = new.ResourceRef
	}

	lang.ResourceMessage = *new.ResourceMessage

	err = database.Update(lang)
	if err != nil {
		return nil, errx.SupportError
	}
	return &lang.ResourceMessage, nil
}

func DeleteLanguageResource(language int, resourceRef string) (*bool, error) {
	var (
		lang   model.LanguageResource
		err    error
		status bool
	)

	lang, err = GetLanguageInfo(language, resourceRef)

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

func GetLanguageResourceByLangAndRef(language int, ref string) (*string, error) {
	var (
		lang model.LanguageResource
		err  error
	)

	lang, err = GetLanguageInfo(language, ref)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return &lang.ResourceMessage, errx.MLangError
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &lang.ResourceMessage, errx.SupportError
	}

	return &lang.ResourceMessage, nil
}

func GetAllLanguageResources(language int) ([]string, error) {
	var (
		lang []string
		err  error
	)
	err = database.Select(&lang, `SELECT language_resource.resource_message FROM language_resource WHERE resource_language = ?`, language)
	if err != nil {
		return lang, errx.SupportError
	}

	return lang, nil
}

func RemoveLanguageResourcesByRef(ref string) (*bool, error) {
	var (
		languages []model.LanguageResource
		err       error
		status    bool
	)
	languages, err = GetLanguagesByRef(ref)
	if err != nil {
		return &status, errx.SupportError
	}

	for _, language := range languages {
		err = RemLang(language)
		if err != nil {
			return &status, errx.SupportError
		}
	}

	status = true
	return &status, nil

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

func GetMessageByRef(ref string) (lang []string, err error) {
	err = database.Select(&lang, `SELECT language_resource.resource_message FROM language_resource WHERE resource_ref = ?`, ref)
	if err != nil {
		return lang, err
	}

	return lang, nil
}

func GetLanguagesByRef(ref string) (lang []model.LanguageResource, err error) {
	err = database.Select(&lang, `SELECT * FROM language_resource WHERE resource_ref = ?`, ref)
	if err != nil {
		return lang, err
	}

	return lang, nil
}

func GetLanguageInfo(language int, ref string) (languageResource model.LanguageResource, err error) {
	err = database.Get(&languageResource, `SELECT *  FROM language_resource WHERE resource_ref = ? AND resource_language = ? `, ref, language)
	if err != nil {
		return languageResource, err
	}
	return languageResource, nil
}

func RemLang(lang model.LanguageResource) (err error) {
	err = database.Delete(lang)
	if err != nil {
		return err
	}
	return nil
}
