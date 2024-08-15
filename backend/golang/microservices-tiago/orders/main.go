package main

import "context"

func main() {
	store := NewStore()
	srv := NewService(store)

	srv.CreateOrder(context.Background())
}
