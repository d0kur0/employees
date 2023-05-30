package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/d0kur0/employees/graph/model"
)

// CreateDepartment is the resolver for the createDepartment field.
func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.NewDepartment) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: CreateDepartment - createDepartment"))
}

// UpdateDepartment is the resolver for the updateDepartment field.
func (r *mutationResolver) UpdateDepartment(ctx context.Context, input model.NewDepartment) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: UpdateDepartment - updateDepartment"))
}

// DeleteDepartment is the resolver for the deleteDepartment field.
func (r *mutationResolver) DeleteDepartment(ctx context.Context, input string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteDepartment - deleteDepartment"))
}

// Departments is the resolver for the departments field.
func (r *queryResolver) Departments(ctx context.Context, limit *int, offset *int) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented: Departments - departments"))
}
