package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func createPizza(pizza int) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Create pizza %d\n", pizza)
}
func timeTrack(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
}
func main() {
	defer timeTrack(time.Now(), "Build Pizza")
	number_of_ovens := 3 // so luong tao ra tren 1 core
	// 3 cai chay tren cung 1 thoi diem
	runtime.GOMAXPROCS(number_of_ovens)
	wg.Add(number_of_ovens)
	for i := 0; i < number_of_ovens; i++ {
		go createPizza(i)
	}
	wg.Wait()
}
