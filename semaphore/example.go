package semaphore

/*

type Empty interface{} và var empty Empty:
Đoạn mã định nghĩa một kiểu Empty không có giá trị cụ thể.
Đây là một loại dữ liệu trống được sử dụng để gửi vào channel.
*/

type Empty interface{}

var empty Empty

/*
data := make([]float64, N) và res := make([]float64, N):
Tạo một slice data chứa N phần tử kiểu float64 và
một slice res cùng kích thước để lưu kết quả của các tính toán.
*/
func doSomething(i int, ix float64) float64 {

	return 0
}
func main() {
	var N int
	data := make([]float64, N)
	res := make([]float64, N)
	/*
	   sem := make(chan Empty, N): Tạo một channel kiểu Empty có kích thước N,
	   là một biến Semaphore. Kích thước của
	   channel này xác định số lượng tác vụ có thể được thực hiện đồng thời.
	*/

	sem := make(chan Empty, N) // semaphore

	/*
	   Trong vòng lặp, mỗi lần lặp tạo một goroutine để thực hiện hàm doSomething(i, xi)
	   với các tham số tương ứng từ slice data. Sau khi hoàn thành tính toán,
	   goroutine gửi một giá trị (đại diện cho một tác vụ đã hoàn thành) vào channel sem.
	*/
	for i, xi := range data {
		go func(i int, xi float64) {
			res[i] = doSomething(i, xi)
			sem <- empty
		}(i, xi)
	}
	// wait for goroutines to finish

	/*
	   Chờ cho tất cả các goroutine hoàn thành bằng cách lặp qua N lần và rút giá trị từ channel sem.
	   Mỗi khi một goroutine hoàn thành, một giá trị (không quan trọng) được rút ra khỏi channel,
	   cho đến khi tất cả goroutine đều hoàn thành.
	*/

	for i := 0; i < N; i++ {
		<-sem
	}

	/*
	   đoạn mã này sử dụng Semaphore (channel kiểu Empty) đaể thực hiện nhiều tính toán song song trên mảng dta,
	   và sau đó chờ cho tất cả các tính toán hoàn thành bằng cách sử dụng channel sem.
	   Cách này giúp đảm bảo rằng tất cả các tính toán đã hoàn thành trước khi chương trình tiếp tục thực hiện các công việc khác.
	*/
}
