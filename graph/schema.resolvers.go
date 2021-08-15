package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.com/kian00sh/rockbooks-be/graph/generated"
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	"gitlab.com/kian00sh/rockbooks-be/src/jwt"
	"gitlab.com/kian00sh/rockbooks-be/src/users"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (string, error) {
	var user users.User
	user.DisplayName = input.DisplayName
	user.Email = input.Email
	user.Password = input.Password
	err := user.Register()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	err := user.Authenticate()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	email, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", grapherrors.ReturnGQLError("Token is invalid!", err)
	}
	token, err := jwt.GenerateToken(email)
	if err != nil {
		return "", grapherrors.ReturnGQLError("Something went wrong with token creation!", err)
	}
	return token, nil
}

func (r *queryResolver) Todo(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
