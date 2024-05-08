package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AccNumber int64  `json:"account_number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {

	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
		AccNumber: int64(rand.Intn(100000)),
	}
}
