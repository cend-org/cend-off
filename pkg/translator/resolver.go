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

func (m *TranslationMutation) RemoveLanguageResource(ctx context.Context, resourceId int) (*bool, error) {
	return DeleteLanguageResource(resourceId)
}

func (q *TranslationQuery) LanguageResources(ctx context.Context) ([]model.LanguageResource, error) {
	return GetAllLanguageResources()
}

func (q *TranslationQuery) LanguageResource(ctx context.Context, resourceRef string, resourceId int) (*model.LanguageResource, error) {
	return GetLanguageResourceByIdAndRef(resourceRef, resourceId)
}
