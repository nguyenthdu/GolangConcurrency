package main

import (
	"log"
)

// TODO: Cơ chế locking đảm bảo khi gửi xong và nhận xong thì các tác vụ tiếp theo mới được thực thị
// goi la goroutine
func printNumbers(numberChan chan int) { // truyen channel vao goruotine tuong ung
	var result int // khi bao bien de luu du lieu vao channel
	for i := 0; i < 20; i++ {
		//fmt.Printf("%d", i)
		result += i // luu du lieu
	}
	numberChan <- result // gui du lieu vao channel, khi gửi dữ liệu xong thì các tác vụ sau mới được thực thi

}
func printChar(charChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		//fmt.Printf("%c", i)
		result += string(i)
	}
	charChan <- result

}
func main() {
	//Tao channel chua du lieu dung de trao doi du lieu
	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)
	//var wg sync.WaitGroup
	//wg.Add(2)
	go printNumbers(chanPrintNumber) //de con tro de truy cap truc tiep
	go printChar(chanPrintChar)
	//Lay du lieu tu channel ra, khi lấy dữ liệu xong thì   các taác vụ sau mới được thực thi
	log.Println("Du lieu trong printNumber la: ", <-chanPrintNumber)
	log.Println("Du lieu trong printChar la: ", <-chanPrintChar)
	// wg.Wait()
	log.Println("\nFinished all.")
}
