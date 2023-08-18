package main

import "fmt"

/*
Hàm ping nhận một channel pings chỉ để gửi (chan<- string) và một chuỗi msg.
Nó đơn giản là gửi msg vào channel pings.
*/
func ping(pings chan<- string, msg string) {
	pings <- msg
}

/*
Hàm pong nhận một channel pings chỉ để nhận (<-chan string) và một channel pongs chỉ để gửi (chan<- string).
Hàm này nhận thông điệp từ channel pings,
sau đó gửi thông điệp đó vào channel pongs.
*/
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg

}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message") // Gọi hàm ping để gửi thông điệp "passed message" vào channel pings.
	pong(pings, pongs)            //Gọi hàm pong với channel pings để nhận thông điệp và channel pongs để gửi thông điệp.
	fmt.Println(<-pongs)
	/*
		Với sự kết hợp của hai goroutine ping và pong, thông điệp được truyền từ ping qua pings đến pong, sau đó thông điệp đó lại được gửi từ pong qua pongs.
		Kết quả là thông điệp "passed message" được in ra từ channel pongs.
	*/
}
