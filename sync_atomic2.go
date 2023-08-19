package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}
func requests() chan int {
	return make(chan int)

}
func main() {
	// nam giua thong tin cau hinh cua server
	var config atomic.Value
	// khoi tao gia tri ban dau
	config.Store(loadConfig())
	go func() {
		// cap nhat thong tin sau moi 10s
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()
	// tao nhieu worker xu ly request
	// dung thong tin cau hinh gan nhat
	for i := 0; i < 10; i++ {
		go func() {
			for req := range requests() {
				c := config.Load()
				// xu ly request voi cau hinh c
				_, _ = req, c
			}

		}()

	}
}
