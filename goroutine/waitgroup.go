package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d stating\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done!\n", id)
}
func main() {
	var wg sync.WaitGroup // doi tat ca ca goroutine duoc chay ket thuc
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
