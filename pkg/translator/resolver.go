package translator

import (
	"context"
	"github.com/cend-org/duval/graph/model"
)

type TranslationQuery struct{}
type TranslationMutation struct{}

func (m *TranslationMutation) NewLanguageResource(ctx context.Context, languageResource model.LanguageResourceInput) (*string, error) {

	return AddLanguageResource(languageResource)
}

func (m *TranslationMutation) RemoveLanguageResource(ctx context.Context, language int, resourceRef string) (*bool, error) {
	return DeleteLanguageResource(language, resourceRef)
}

func (q *TranslationQuery) LanguageResources(ctx context.Context, language int) ([]string, error) {
	return GetAllLanguageResources(language)
}

func (q *TranslationQuery) LanguageResource(ctx context.Context, language int, resourceRef string) (*string, error) {
	return GetLanguageResourceByLangAndRef(language, resourceRef)
}

func (m *TranslationMutation) RemoveLanguageResources(ctx context.Context, resourceRef string) (*bool, error) {
	return RemoveLanguageResourcesByRef(resourceRef)
}

func (q *TranslationQuery) AllReferencedLanguageResources(ctx context.Context, resourceRef string) ([]string, error) {
	return GetMessageByRef(resourceRef)
}

func (m *TranslationMutation) AddOrGetLanguageResource(ctx context.Context, language int, resourceRef string) (*string, error) {
	var (
		languageResource model.LanguageResourceInput
	)

	languageResource = model.LanguageResourceInput{
		ResourceLanguage: &language,
		ResourceRef:      &resourceRef,
	}

	return AddOrGetResource(languageResource)
}

func (m *TranslationMutation) UpdLanguageResource(ctx context.Context, languageResource model.LanguageResourceInput) (*string, error) {
	return UpdateLanguageResource(languageResource)
}
