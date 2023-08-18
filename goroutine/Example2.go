package main

import (
	"fmt"
	"sync"
)

func HeavyFunction1(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
}

func HeavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
}

func main() {
	wg := new(sync.WaitGroup) // tao doi tuong theo doi cac goroutine
	wg.Add(2)                 // them so luong cong viec
	go HeavyFunction1(wg)     // chay goroutine song song truyen den con tro WaitGroup de biet khi nao hoan thanh
	go HeavyFunction2(wg)
	wg.Wait() //doi tat ca cac goroutine ket thuc
	fmt.Printf("All Finished!")
}

/*
Đoạn mã trình bày cách sử dụng goroutine và sync.WaitGroup trong Go để thực hiện hai hàm nặng và đồng thời đợi chúng hoàn thành trước khi tiếp tục thực hiện mã trong hàm main(). Hãy tìm hiểu từng phần của mã:

Import Packages:
Đầu tiên, chúng ta import các gói cần thiết: "fmt" để định dạng in ra và "sync" để sử dụng sync.WaitGroup.

HeavyFunction1 và HeavyFunction2:
Đây là hai hàm đại diện cho các tác vụ nặng (heavy tasks) mà bạn muốn thực hiện đồng thời. Mỗi hàm nhận một con trỏ đến một sync.WaitGroup (wg), cho biết chúng sẽ được hoàn thành. Trong mỗi hàm, chúng ta sử dụng defer wg.Done() để đánh dấu rằng hàm đã hoàn thành. Trong thực tế, "các công việc nặng" này sẽ được thay thế bằng các tác vụ thực sự mà bạn muốn thực hiện.

Hàm main():
Trong hàm main(), chúng ta tạo một đối tượng sync.WaitGroup mới (wg := new(sync.WaitGroup)) để theo dõi các goroutine. Sau đó, chúng ta thêm hai công việc vào WaitGroup bằng cách sử dụng wg.Add(2) - mỗi công việc tương ứng với một goroutine sẽ được chạy.

Tiếp theo, chúng ta sử dụng từ khóa go để khởi chạy hai goroutine song song. Mỗi goroutine chạy một hàm nặng khác nhau và truyền con trỏ đến WaitGroup vào hàm để biết khi nào chúng hoàn thành. Sau đó, chúng ta sử dụng wg.Wait() để đợi tất cả các goroutine kết thúc.

Cuối cùng, sau khi tất cả các goroutine hoàn thành, dòng fmt.Printf("All Finished!") được thực hiện, thông báo rằng tất cả công việc đã hoàn thành.

Tóm lại, đoạn mã này sử dụng sync.WaitGroup để đồng bộ hóa và đợi nhiều goroutine thực hiện các tác vụ nặng, sau đó thông báo khi tất cả chúng hoàn thành.*/
