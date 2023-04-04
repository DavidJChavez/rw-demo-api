package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"DavidJChavez/rw-demo-api/configs"
	"DavidJChavez/rw-demo-api/graph/model"
	"context"
)

// `CreateUser` is the resolver for the `createUser` field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := configs.Db.CreateUser(&input)
	return user, err
}

// `Users` is the resolver for the `users` field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := configs.Db.ListUsers()
	return users, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
