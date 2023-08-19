package main

import (
	"fmt"
	"sync"
)

/*
Tác vụ atomic trên một vùng nhớ chia sẻ thì đảm bảo rằng vùng nhớ đó chỉ có thể được truy cập bởi một Goroutine
tại một thời điểm. Để đạt được điều này ta có thể dùng sync.Mutex.
*/
// / total là một atomic struct
var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	// thông báo hoàn thành khi ra khỏi hàm
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		// chặn các Goroutines khác vào
		total.Lock()
		// bây giờ, lệnh total.value += i được đảm bảo là atomic (đơn nguyên)
		total.value += i
		// bỏ chặn
		total.Unlock()
	}
}
func main() {
	// khai báo wg để main Goroutine dừng chờ các Goroutines khác trước khi kết thúc chương trình
	var wg sync.WaitGroup
	// wg cần chờ 2 Goroutines khác
	wg.Add(2)
	go worker(&wg) // thực thi Goroutines thứ nhất
	go worker(&wg) // thực thi Goroutines thứ hai
	wg.Wait()      // thực thi Goroutines thứ hai
	fmt.Println(total.value)
}

/*
Trong một chương trình đồng thời, ta cần có cơ chế để lock và unlock trước và sau khi truy cập vào vùng critical
section. Nếu không có sự bảo vệ biến total , kết quả cuối cùng có thể bị sai khác do sự truy cập đồng thời của
nhiều thread.
*/
