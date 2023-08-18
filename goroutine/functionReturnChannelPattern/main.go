package main

import (
	"fmt"
	"log"
)

func printData(msg string) chan string { // nhan vao 1 message va tra ve 1 chan
	//Nhan vien order tao the dk cho khach hang
	result := make(chan string)
	// Giao cho nhan vien pha che
	go func() { // tao anonymous goroutine
		// tinh hieu day vao trong channel -> khach hang nhan duoc thong bao
		for i := 0; i <= 5; i++ {
			// moi lan chay thi day du lieu vao chan
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	// dua cho khach hang the
	return result
}

func main() {
	//khach hang order va nhan duoc the cho
	bridge := printData("hello")
	// khach hang nhan duoc thong bao  lay hang
	for i := 0; i < len(bridge); i++ {
		log.Println("Receive from", <-bridge) // doc du lieu tu channel
	}
	log.Println("\nFinished all.")
}
