package main

import (
	"context"
	"fmt"
	"time"
)

func heavyWorkFailResult() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go heavyWork(ctx, 4)

	time.Sleep(6 * time.Second)

}
func heavyWorkPassResult() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go heavyWork(ctx, 2)
	time.Sleep(6 * time.Second)

}

func heavyWork(ctx context.Context, seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Println("Heavy work is done")
	case <-ctx.Done():
		fmt.Println("Work cancelled")
	}
}
