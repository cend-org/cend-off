package resolver

import (
	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/pkg/academic"
	usr "github.com/cend-org/duval/pkg/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

type mutationResolver struct {
	*usr.UserMutation
	*academic.AcademicMutation
}

type queryResolver struct {
	*usr.UserQuery          `json:"*Usr.Query,omitempty"`
	*academic.AcademicQuery `json:"*Academic.Query,omitempty"`
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{} }
