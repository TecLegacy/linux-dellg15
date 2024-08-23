package main

import "time"

func main() {

	go func() {
		heavyWorkFailResult()
	}()

	go func() {
		heavyWorkPassResult()
	}()

	time.Sleep(6 * time.Second)

}
