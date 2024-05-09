package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	AccNumber int64     `json:"account_number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_At"`
}

func NewAccount(firstName, lastName string) *Account {

	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		AccNumber: int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
