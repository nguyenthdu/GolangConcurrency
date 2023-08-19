package main

import (
	"GolangConcurrency/ModelPubSub/pubsub"
	"fmt"
	"strings"
	"time"
)

func main() {
	// khoi tao 1 publisher
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	// de dam bao p duoc dong truoc khi exit
	defer p.Close()
	//'all' subscribe het tat ca topic
	all := p.Subscriber()
	//subscribe cac topic co "golang"
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("hello, world!")
	p.Publish("hello, golang!")
	// print nhung fi subscribe 'all' nhan duoc
	go func() {
		for msg := range all {
			fmt.Println("golang", msg)
		}
	}()

	// print nhung gi subscribe 'golang nhan duoc
	go func() {
		for msg := range golang {
			fmt.Println("golang", msg)
		}
	}()
	// thoat ra sau khi chay 3s
	time.Sleep(3 * time.Second)
}

/*
Trong mô hình pub/sub, mỗi thông điệp được gửi tới nhiều subscriber. Publisher thường không biết hoặc không quan
tâm subscriber nào nhận được thông điệp. Subscriber và publisher có thể được thêm vào động ở thời điểm thực thi,
cho phép các hệ thống phức tạp có thể phát triển theo thời gian. Trong thực tế, những ứng dụng như dự báo thời tiết
có thể áp dụng mô hình concurrency này.
*/
