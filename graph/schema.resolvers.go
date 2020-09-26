package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/yuzuy/gqlgen-todos/graph/generated"
	"github.com/yuzuy/gqlgen-todos/graph/model"
)

func (r *mutationResolver) AddTodo(ctx context.Context, input model.AddTodoRequest) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprint(rand.Int()),
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "user_" + input.UserID,
		},
	}

	err := r.DB.AddTodo(todo)
	return todo, err
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoRequest) (*model.Todo, error) {
	todo, err := r.DB.FirstTodo(input.ID)
	if err != nil {
		return nil, err
	}
	todo.Text = input.Text

	err = r.DB.UpdateTodo(todo)
	return todo, err
}

func (r *mutationResolver) MarkAsDone(ctx context.Context, input model.MarkAsDoneRequest) (*model.Todo, error) {
	todo, err := r.DB.FirstTodo(input.ID)
	if err != nil {
		return nil, err
	}
	todo.Done = true

	err = r.DB.UpdateTodo(todo)
	return todo, err
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, input model.RemoveTodoRequest) (string, error) {
	if err := r.DB.RemoveTodo(input.ID); err != nil {
		return "", err
	}
	return input.ID, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.DB.FindTodos(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
