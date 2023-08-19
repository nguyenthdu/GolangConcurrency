package main

//Thay vì dùng mutex, chúng ta cũng có thể dùng package sync/atomic, đây là giải pháp hiệu quả hơn đối với một biến
//số học.
import (
	"fmt"
	"sync"
	// khai báo biến gói sync/atomic
	"sync/atomic"
)

// biến total được truy cập đồng thời
var total2 uint64

func worker2(wg *sync.WaitGroup) {
	// wg thông báo hoàn thành khi ra khỏi hàm
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		// lệnh cộng atomic.AddUint64 total được đảm bảo là atomic (đơn nguyên)
		atomic.AddUint64(&total2, i)
	}
}
func main() {
	// wg được dùng để dừng hàm main đợi các Goroutines khác
	var wg sync.WaitGroup
	// wg cần đợi hai Goroutines gọi lệnh Done() mới thực thi tiếp
	wg.Add(2)
	// tạo Goroutines thứ nhất
	go worker2(&wg)
	// tạo Goroutines thứ hai
	go worker2(&wg)
	// bắt đầu việc đợi
	wg.Wait()
	// in ra kết quả
	fmt.Println(total2)
}
