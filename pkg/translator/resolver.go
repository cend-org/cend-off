package translator

import (
	"context"
	"github.com/cend-org/duval/graph/model"
)

type TranslationQuery struct{}
type TranslationMutation struct{}

func (m *TranslationMutation) NewLanguageResource(ctx context.Context, languageResource model.LanguageResourceInput) (*model.LanguageResource, error) {
	return AddLanguageResource(languageResource)
}

func (m *TranslationMutation) RemoveLanguageResource(ctx context.Context, language int, resourceRef string) (*bool, error) {
	return DeleteLanguageResource(language, resourceRef)
}

func (q *TranslationQuery) LanguageResources(ctx context.Context, language int) ([]model.LanguageResource, error) {
	return GetAllLanguageResources(language)
}

func (q *TranslationQuery) LanguageResource(ctx context.Context, language int, resourceRef string) (*model.LanguageResource, error) {
	return GetLanguageResourceByLangAndRef(language, resourceRef)
}

func (m *TranslationMutation) RemoveLanguageResources(ctx context.Context, resourceRef string) (*bool, error) {
	return RemoveLanguageResourcesByRef(resourceRef)
}

func (q *TranslationQuery) AllReferencedLanguageResources(ctx context.Context, resourceRef string) ([]model.LanguageResource, error) {
	return GetLanguagesByRef(resourceRef)
}
