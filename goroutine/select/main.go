package main

import (
	"log"
	"time"
)

//TODO: chon goroutine

// goole search
func googleSearch(search chan string) {
	time.Sleep(3 * time.Second)
	search <- "Found from Google!"
}
func bingSearch(search chan string) {
	time.Sleep(5 * time.Second)
	search <- "Found from Bing!"
}
func main() {
	googleS := make(chan string)
	bingS := make(chan string)
	go googleSearch(googleS)
	go bingSearch(bingS)
	//locking- khi chay den khoi nay se dung lai khi co du lieu gui vao thi moi chay tiep
	select { // nhan duoc ket qua tu channel truoc thi in ra truoc
	case result := <-googleS: // neu nhu 2 channel co thoi gian bang nhau thi se tu random
		log.Println(result)
	case result := <-bingS:
		log.Println(result)
	}
	log.Println("Main finished")

}
