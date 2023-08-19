package main

import (
	"fmt"
	"log"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d", i)
	}
	wg.Done()
}
func printChar(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c", i)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg) //de con tro de truy cap truc tiep
	go printChar(&wg)
	wg.Wait() // dung waitGroup nghia la bat ham main phai cho cho den khi 2 goroutine thu hien xong
	log.Println("\nFinished all.")
}
