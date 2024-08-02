package resolver

import (
	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/pkg/academic"
	mediafile "github.com/cend-org/duval/pkg/media"
	"github.com/cend-org/duval/pkg/translator"
	usr "github.com/cend-org/duval/pkg/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

type mutationResolver struct {
	*usr.UserMutation
	*academic.AcademicMutation
	*translator.TranslationMutation
}
type queryResolver struct {
	*usr.UserQuery               `json:"*Usr.Query,omitempty"`
	*academic.AcademicQuery      `json:"*Academic.Query,omitempty"`
	*mediafile.MediaQuery        `json:"*Media.Query,omitempty"`
	*translator.TranslationQuery `json:"*Translation.Query,omitempty"`
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{} }
