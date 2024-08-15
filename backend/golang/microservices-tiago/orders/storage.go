package main

import "context"

type Store struct {
	// mongodb instance
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(context.Context) error {
	return nil
}
