# Xử lý Đồng Thời và Song Song trong Golang
![modelparra](https://github.com/nguyenthdu/GolangConcurrency/assets/110290495/ba56d35c-7ba2-4d96-9fed-98532c8975f5)

## Concurrency và Parallelism - Mô hình lập trình đồng thời và lập trình song song

**Concurrency - Xử lý Đồng Thời**

Mô hình lập trình đồng thời (concurrency) đề cập đến khả năng phân chia và quản lý nhiều tác vụ khác nhau trong cùng một khoảng thời gian. Tại một thời điểm, chỉ có thể xử lý một tác vụ. Khái niệm này tương phản với xử lý tuần tự, trong đó chỉ xử lý một tác vụ tại mỗi thời điểm. Ví dụ, trình duyệt web cho phép mở nhiều tab song song, nhưng chỉ một tab được hiển thị tại một thời điểm.

Tất cả chương trình trên máy tính chạy thông qua quản lý của hệ điều hành dưới dạng các tiến trình, được gọi là process, mỗi process có một process ID (PID) riêng để hệ điều hành quản lý. Các tác vụ của tiến trình được xử lý bởi CPU core (nhân CPU). Mặc dù máy tính có CPU với một nhân, xử lý đồng thời nhiều tác vụ từ các tiến trình khác nhau vẫn có thể xảy ra. CPU không đợi cho đến khi một tác vụ hoàn thành trước khi thực hiện tác vụ khác, mà thay vào đó, nó chia nhỏ các tác vụ lớn và xử lý xen kẽ chúng để tận dụng thời gian rảnh rỗi.
![concurrency](https://github.com/nguyenthdu/GolangConcurrency/assets/110290495/60c21dbb-1c44-4669-abf3-ac3a5eccb77b)

**Parallelism - Xử lý Song Song**

Xử lý song song là khả năng xử lý nhiều tác vụ độc lập trong cùng một thời điểm. Điều này chỉ khả thi trên máy tính có nhiều nhân CPU. Thay vì chỉ có thể xử lý một tác vụ nhỏ tại mỗi thời điểm, với nhiều nhân CPU, chúng ta có thể xử lý các tác vụ song song trên các nhân CPU khác nhau.

Ví dụ, bạn có thể nghe nhạc, đọc tài liệu và tải tài liệu cùng một lúc. Trong xử lý song song, các tác vụ này có thể được xử lý bởi các nhân CPU riêng biệt.
![image](https://github.com/nguyenthdu/GolangConcurrency/assets/110290495/6d6427a8-dcac-4633-86fb-299dd8dd0c42)


## Xử lý Đồng Thời trong Golang

Trong Golang, mô hình lập trình đồng thời và song song là một trong những đặc điểm nổi bật. Golang hỗ trợ cả hai mô hình này thông qua goroutine và channel.

- **Goroutine**: Là các đơn vị thực thi nhẹ, tương tự như luồng ảo, được quản lý bởi máy ảo Golang. Một chương trình Golang có thể tạo hàng nghìn goroutine và chạy chúng đồng thời trên một số lượng hạn chế các luồng vật lý. Goroutine cho phép chuyển đổi linh hoạt giữa các luồng vật lý, giúp tận dụng tối đa các bộ xử lý.

- **Channel**: Là cơ chế để giao tiếp an toàn và đồng bộ giữa các goroutine. Channel cho phép truyền dữ liệu từ một goroutine đến một goroutine khác. Chúng có thể được sử dụng để xây dựng các mẫu lập trình đồng thời như pipeline, fan-in, fan-out, worker pool và nhiều mẫu khác.

Golang cho phép lập trình song song bằng cách sử dụng hàm `runtime.GOMAXPROCS` để định số lượng luồng vật lý sử dụng cho việc lên lịch các goroutine. Khi số luồng vật lý bằng hoặc lớn hơn số bộ xử lý, các goroutine có thể chạy song song trên các bộ xử lý khác nhau.

## Goroutines và System Threads

Goroutines là khái niệm quan trọng về đồng thời trong ngôn ngữ Go. Đơn giản, trong Go, chúng ta sử dụng goroutine để thực hiện nhiều tác vụ cùng một lúc. So với các ngôn ngữ khác, việc tạo ra goroutine ít tốn kém hơn việc tạo thread. Sự khác biệt cơ bản giữa goroutine và thread đã dẫn đến những lợi ích đáng kể.

Khi tạo ra một goroutine, chúng ta chỉ cần sử dụng từ khóa go, điều này rất đơn giản và linh hoạt. Trong thực tế, goroutine và thread có sự khác biệt đáng kể:

Mỗi thread trong hệ thống sẽ có một vùng nhớ stack cố định (thông thường khoảng 2MB). Vùng nhớ này lưu trữ các thông tin như tham số, biến cục bộ và địa chỉ trả về khi gọi hàm.
Tuy nhiên, vùng nhớ stack cố định có thể gây ra các vấn đề như stack overflow cho những chương trình đệ quy sâu và lãng phí vùng nhớ cho các chương trình đơn giản.
Goroutines đã giải quyết vấn đề này một cách linh hoạt:

Mỗi goroutine bắt đầu với một vùng nhớ stack nhỏ (khoảng 2KB hoặc 4KB).
Khi cần, goroutine có thể tự động tăng kích thước stack để tránh việc vượt quá giới hạn, và kích thước tối đa của stack có thể lên đến 1GB.
Với chi phí thấp của việc tạo goroutine, chúng ta có thể linh hoạt tạo và giải phóng hàng nghìn goroutines mà không gặp vấn đề về tài nguyên.
So với các ngôn ngữ như Java, nơi các thread được quản lý bởi hệ điều hành, Go sử dụng cơ chế định thời của riêng mình cho goroutines. Cơ chế này tương tự với cách hệ điều hành quản lý thread, nhưng điều khiển ở mức chương trình. Biến runtime.GOMAXPROCS quy định số lượng thread được sử dụng cho goroutines và đảm bảo chúng hoạt động song song một cách hiệu quả.

![image](https://github.com/nguyenthdu/GolangConcurrency/assets/110290495/5dd3d5cb-46e8-44f9-aa13-05cf8b0e79ac)

## Mô hình Thực Thi Đồng Thời

Một ưu điểm quan trọng của Golang là tích hợp sẵn khái niệm xử lý đồng thời. Golang áp dụng lý thuyết tương tranh CSP (Communicating Sequential Process), đề xuất bởi Hoare vào năm 1978. Khái niệm này được áp dụng vào lập trình concurrency trong Golang.

Trong việc xử lý đồng thời trong Golang, ta cần quan tâm đến các khái niệm liên quan đến xử lý xung đột và khóa tài nguyên. Một số khái niệm quan trọng bao gồm:
- **Race condition**: Xảy ra khi nhiều goroutines chạy song song mà kết quả không nhất quán khi so sánh với chạy tuần tự. Điều này có thể xảy ra khi các goroutines truy cập cùng một biến mà không có sự đồng bộ hóa.
- **Locking**: Kỹ thuật để đảm bảo chỉ có một goroutine được truy cập vào một tài nguyên tại một thời điểm. Tuy nhiên, locking có thể gây ra deadlock nếu không sử dụng cẩn thận.
- **Deadlock**: Xảy ra khi nhiều goroutines bị treo do chờ nhau giải phóng tài nguyên. Deadlock có thể xảy ra khi sử dụng locking hoặc channel không đúng cách.

# Cơ chế Synchronized

Trong Golang, cơ chế đồng bộ hóa (synchronization) được sử dụng để đảm bảo tính nhất quán và an toàn cho các tài nguyên chia sẻ giữa các goroutine. Cơ chế này bao gồm sử dụng mutex (mutual exclusion), semaphore và channel. Mutex là một loại khóa dùng để giới hạn quyền truy cập của một hoặc nhiều goroutine vào một tài nguyên. Semaphore cung cấp khả năng giới hạn số lượng goroutine có thể truy cập tài nguyên cùng một lúc. Channel cung cấp một cách để các goroutine truyền dữ liệu an toàn qua lại.

# Wait Group và Fan-out, Fan-in

Trong trường hợp các goroutine cần thực hiện xong trước khi chương trình kết thúc, ta có thể sử dụng Wait Group. Wait Group là một cơ chế đồng bộ hóa cho phép chương trình chờ cho đến khi một tập hợp các goroutine kết thúc.

Các mẫu lập trình đồng thời như Fan-out và Fan-in cũng đáng để nhắc đến. Fan-out đề cập đến việc chia tác vụ thành nhiều goroutine để tận dụng nhiều bộ xử lý. Fan-in đề cập đến việc tổng hợp kết quả từ nhiều goroutine vào một kênh duy nhất.

# Cơ chế Atomic

Package sync/atomic cung cấp cơ chế để thực hiện các thao tác atomic trên biến số nguyên và con trỏ. Các hàm như Add, Load, Store, Swap, CompareAndSwap giúp đảm bảo tính nhất quán khi nhiều goroutine truy cập cùng một biến.




# Select và Non-blocking Communication

Trong Golang, cơ chế select cho phép chọn lựa giữa nhiều trường hợp giao tiếp. Điều này giúp tránh việc chương trình bị treo khi gửi hoặc nhận dữ liệu từ các goroutine khác.

**Tài liệu tham khảo**


https://zalopay-oss.github.io/go-advanced/?fbclid=IwAR2VZA12geYELXbeIWoDiNj54zdNNvXNEqE4-CIBSidwMoXZRwu_cu12r8U

https://go.dev/blog/waza-talk

The Go Programming Language
