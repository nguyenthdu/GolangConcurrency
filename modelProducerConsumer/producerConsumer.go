package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// producer: liên tục tạo ra một chuỗi số nguyên dựa trên bội số factor và đưa vào channel
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
} // consumer: liên tục lấy các số từ channel ra để print
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	// hàng đợi
	ch := make(chan int, 64)
	// tạo một chuỗi số với bội số 3
	go Producer(3, ch)
	// tạo một chuỗi số với bội số 5
	go Producer(5, ch)
	// tạo consumer
	go Consumer(ch)
	// thoát ra sau khi chạy trong một khoảng thời gian nhất định
	time.Sleep(5 * time.Second)
}

//Chúng ta có thể để hàm main giữ trạng thái block mà không thoát và chỉ thoát khỏi chương trình khi người dùng gõ
//Ctrl-C :

func main2() {
	// hàng đợi
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)
	// Ctrl+C để thoát
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
