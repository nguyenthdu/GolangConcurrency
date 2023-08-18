package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/rate-limiting
func main() {
	requests := make(chan int, 5) // tạo một channel requests với một buffer size là 5, giới hạn số lượng yêu cầu được chờ xử lý.
	for i := 1; i < 5; i++ {      // gửi yêu cầu vào channel
		requests <- i
	}
	close(requests)                              // đóng khi đã hoàn tất gửi
	limiter := time.Tick(200 * time.Millisecond) // giới hạn tốt độ xử lý yêu cầu, gửi giá trị sau mỗi 200 mili

	for req := range requests { // lặp qua các request
		<-limiter // chờ cho đến khi một gia trị thời gian từ limiter được nhận, giới hạn tốt độ xử lý
		fmt.Println("Requset", req, time.Now())
	}
	burstyLimit := make(chan time.Time, 3) //Tạo một channel burstyLimit với buffer size là 3 để lưu trữ thời gian.
	for i := 0; i < 3; i++ {               //gửi các thời gian hiện tại vào burstyLimit, tạo một số lượng ban đầu các giá trị để "nổ" yêu cầu.
		burstyLimit <- time.Now()
	}
	//Một goroutine được bắt đầu để gửi thời gian vào burstyLimit sau mỗi 200 milliseconds.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimit <- t
		}
	}()
	burstyRequests := make(chan int, 5) //Tạo một channel burstyRequests với buffer size là 5 để lưu trữ yêu cầu.
	for i := 0; i <= 5; i++ {           // lặp qua từng yêu cầu
		burstyRequests <- i //chờ cho đến khi có một thời gian từ burstyLimit, cho phép xử lý yêu cầu "nổ".
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimit
		fmt.Println("request", req, time.Now())

	}
}

/*
time.Tick trong Go được sử dụng để tạo một channel có khả năng gửi các giá trị thời gian liên tục sau một khoảng thời gian cố định. Nó là một cách tiện lợi để tạo ra một loạt các giá trị thời gian mà bạn có thể đợi và nhận trong các vòng lặp hoặc trong các cấu trúc kiểm soát xử lý thời gian.

Cú pháp time.Tick(duration) sẽ trả về một channel kiểu <-chan time.Time, trong đó duration là một khoảng thời gian. Khi sử dụng time.Tick, sau mỗi duration, một giá trị thời gian sẽ được gửi vào channel, và bạn có thể nhận giá trị này bằng cách sử dụng phép toán <- trên channel.
*/
