package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend/graph/generated"
	"backend/graph/model"
	"backend/internal/user"
	"context"
)

func (r *mutationResolver) UserCreate(ctx context.Context, data model.UserInput) (*model.User, error) {
	var ret []*model.User
	var user user.User
	return user.Create(data)
}

func (r *queryResolver) UserGet(ctx context.Context, id *string) ([]*model.User, error) {
	var ret []*model.User
	var user user.User
	res, err := user.Get(id)
	if err != nil {
		return ret, err
	}
	ret = append(ret, res)
	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
