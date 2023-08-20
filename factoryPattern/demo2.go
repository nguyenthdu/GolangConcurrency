package main

import (
	"fmt"
	"time"
)

func main() {
	suck1(pump1())
	time.Sleep(1e9)
}
func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func suck1(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

}
