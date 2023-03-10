// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Customer struct {
	ID        string    `json:"id"`
	BrandName string    `json:"brand_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerInput struct {
	BrandName string `json:"brand_name"`
}

type User struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	Email      string    `json:"email"`
	Passwd     string    `json:"passwd"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserInput struct {
	Email      string `json:"email"`
	CustomerID string `json:"customer_id"`
	Passwd     string `json:"passwd"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

type UserUpdate struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	CustomerID string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}
