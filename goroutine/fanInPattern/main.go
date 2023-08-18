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
func fanIn(chan1, chan2 chan string) chan string { // channel nao gui truoc thi se nhan truoc
	c := make(chan string)
	go func() {
		for {
			select { // ham nay se lien tuc lang nghe tu 2 channel
			case <-chan1:
				c <- <-chan1 // cai nao den truoc thi no se day vao
			case <-chan2:
				c <- <-chan2
			}
		}
	}()
	return c
}
func main() {
	//khach hang order va nhan duoc the cho
	bridge := printData("hello")
	coffe := printData("HighLande")
	services := fanIn(bridge, coffe)
	// khach hang nhan duoc thong bao  lay hang
	for i := 0; i < len(services); i++ {
		//log.Println("Receive from", <-bridge) // thay vi lang nghe truc tiep tu 2 channel thi chi can lang nghe tu channel services
		log.Println("Receive from", <-services)

	}
	log.Println("\nFinished all.")
}
