package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/tickers
func main() {
	//Tạo một ticker mới với khoảng thời gian 500 milliseconds. Ticker sẽ gửi một giá trị thời gian qua channel ticker.C sau mỗi khoảng thời gian này.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) // thong bao ket thuc goroutine khi chay ticker
	//xu ly ticker
	go func() {
		for { // lăng nghe tín hiệu tu cac channel
			select {
			case <-done: // neu channel được đọc kết thúc bằng return
				return
			case t := <-ticker.C: // khi có 1 giá trị thời gian được gửi qua channel ticker.C
				fmt.Println("Tick at: ", t) // in ra thông báo cùng với thời điểm t
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond) //chờ  1600 để ticker hoạt động đủ 1 số lần
	ticker.Stop()                       // Dừng ticker, ngăn việc gửi thêm giá trị thời gian vào channel  ticker.C
	done <- true
	fmt.Println("Ticker stoped")
	//Mã này cho thấy cách sử dụng time.NewTicker để tạo một ticker và sử dụng channels để kiểm soát hoạt động của ticker trong một goroutine.
}
