package main

import "fmt"

func sendDataChan(chanData chan int) {
	chanData <- 10

}
func main() {
	chanInt := make(chan int)
	chanInt <- 5 // sẽ xảy ra deadlock  vì  từ hàm main gửi dữ liệu vào channel nhưng không có goroutine nào đọc dữ liệu từ channel đó  ra
	<-chanInt    // tương tự việc đọc dữ liệu cũng vậy
	//Cach khac phuc la gui du lieu tu goroutine vao channel
	go sendDataChan(chanInt)     // gui du lieu vao
	fmt.Println("%d", <-chanInt) // doc du lieu ra
	// Neu thieu 1 trong 2 thi se xay ra deadlock
}
