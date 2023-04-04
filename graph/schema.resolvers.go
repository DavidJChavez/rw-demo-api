package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"DavidJChavez/rw-demo-api/graph/model"
	"context"
	"fmt"
	"time"
)

// `CreateUser` is the resolver for the `createUser` field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:        fmt.Sprintf("T%d", len(r.users)+1),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		Phone:     input.Phone,
		Email:     input.Email,
	}
	return user, nil
}

// `Users` is the resolver for the `users` field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }