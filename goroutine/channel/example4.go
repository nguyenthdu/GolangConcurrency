package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) { // sending data to ch channel
	ch <- "Ho Chi Minh"
	ch <- "Ha Noi"
	ch <- "Tokyo"
}
func getData(ch chan string) {
	var input string
	for {
		input = <-ch // receiving data sent to ch channel
		fmt.Printf("%s", input)
	}
	close(ch)
}
func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}
