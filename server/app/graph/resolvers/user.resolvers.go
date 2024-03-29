package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/m-butterfield/heatmap-creator/server/app/data"
	"github.com/m-butterfield/heatmap-creator/server/app/graph/generated"
	"github.com/m-butterfield/heatmap-creator/server/app/graph/model"
	"github.com/m-butterfield/heatmap-creator/server/app/lib"
	"google.golang.org/api/idtoken"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserCreds) (*data.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	gctx, err := ginContextFromContext(ctx)
	if err != nil {
		return false, internalError(err)
	}
	cookie, err := getSessionCookie(gctx.Request)
	if err != nil {
		return false, internalError(err)
	}
	if cookie == nil {
		return true, nil
	}
	if err := r.DS.DeleteAccessToken(cookie.Value); err != nil {
		return false, internalError(err)
	}
	unsetSessionCookie(gctx.Writer)
	return true, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, credential string) (*model.LoginResponse, error) {
	result, err := idtoken.Validate(ctx, credential, lib.GoogleOAuthClientID)
	if err != nil {
		return nil, internalError(err)
	}
	email, ok := result.Claims["email"].(string)
	if !ok {
		return nil, internalError(err)
	}
	if email == "" {
		return nil, internalError(errors.New("no email present in token"))
	}

	user, err := r.DS.GetUser(email)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		user = &data.User{
			Username: email,
		}
		if err = r.DS.CreateUser(user); err != nil {
			return nil, internalError(err)
		}
	}

	token, err := cookieLogin(ctx, r.DS, user)
	if err != nil {
		return nil, internalError(err)
	}

	return &model.LoginResponse{User: user, QueryToken: token.QueryToken}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*data.User, error) {
	user, err := loggedInUser(ctx)
	if err != nil {
		return nil, internalError(err)
	}
	return user, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, username string) (*data.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetUserStats is the resolver for the getUserStats field.
func (r *queryResolver) GetUserStats(ctx context.Context) (*model.UserStats, error) {
	user, err := loggedInUser(ctx)
	if err != nil {
		return nil, internalError(err)
	}
	yesterday := time.Now().AddDate(0, 0, -1)
	queryCount, err := r.DS.GetQueryCountForUser(user.ID, &yesterday)
	if err != nil {
		return nil, internalError(err)
	}
	return &model.UserStats{
		NumQueries: queryCount,
		MaxQueries: user.DailyQueries,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
