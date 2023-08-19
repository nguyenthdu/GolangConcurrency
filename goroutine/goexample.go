package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/goroutines
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") // cach thong thuong, chay dong bo - synchronously

	go f("goroutine") // chay voi goroutine

	go func(msg string) { // anonymous function
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
