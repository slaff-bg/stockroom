package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"

	"github.com/slaff-bg/stockroom/facilities"
	"github.com/slaff-bg/stockroom/ports/graph/model"
	"gorm.io/gorm/clause"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	u := &model.User{
		CustomerID: input.CustomerID,
		Email:      input.Email,
		Passwd:     facilities.PasswdGen(input.Passwd),
		FirstName:  input.FirstName,
		LastName:   input.LastName,
	}

	if err := r.GDB.Clauses(clause.Returning{}).Omit("id", "created_at", "updated_at").Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserUpdate) (*model.User, error) {
	u := &model.User{
		ID:         input.ID,
		CustomerID: input.CustomerID,
		Email:      input.Email,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
	}

	// if err := r.GDB.Clauses(clause.Returning{}).Omit("id", "customer_id", "passwd", "created_at", "updated_at").
	if err := r.GDB.Clauses(clause.Returning{}).Select("email", "first_name", "last_name").
		Where("id = ?", input.ID).
		Where("customer_id = ?", input.CustomerID).
		Save(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.GDB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UserByID is the resolver for the UserById field.
func (r *queryResolver) UserByID(ctx context.Context, id *string) (*model.User, error) {
	var u model.User
	if err := r.GDB.WithContext(ctx).Where("id = ?", id).Take(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	var customers []*model.Customer
	if err := r.GDB.WithContext(ctx).Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

// CustmerByID is the resolver for the CustmerById field.
func (r *queryResolver) CustmerByID(ctx context.Context, id *string) (*model.Customer, error) {
	var cmr model.Customer
	if err := r.GDB.WithContext(ctx).Where("id = ?", id).Take(&cmr).Error; err != nil {
		return nil, err
	}
	return &cmr, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
