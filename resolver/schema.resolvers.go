package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	models1 "github.com/PandaGoL/gql-server-booking/models"
	server1 "github.com/PandaGoL/gql-server-booking/server"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user models1.NewUser) (*models1.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - CreateUser"))
}

// GetUsers is the resolver for the GetUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) (*models1.GetUsersResponse, error) {
	panic(fmt.Errorf("not implemented: GetUsers - GetUsers"))
}

// Mutation returns server1.MutationResolver implementation.
func (r *Resolver) Mutation() server1.MutationResolver { return &mutationResolver{r} }

// Query returns server1.QueryResolver implementation.
func (r *Resolver) Query() server1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
