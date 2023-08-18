package main

import "log"

func inputDatafirst(nums ...int) chan int { // nhan vao cac gia tri kieu int
	result := make(chan int)
	go func() {
		for i := 0; i < len(nums); i++ {
			result <- nums[i]
		}
		close(result) // doc xong nen dong lai neu khong, thi se gap deadlock~~
	}()
	return result
}
func inputDatasecond(dataChan chan int) chan int { // doc du lieu tu channel duoc truyen vao va xu ly no
	result := make(chan int)
	go func() {
		for item := range dataChan {
			result <- item //day du lieu tu chanData va chan result
		}
		close(result)
	}()
	return result
}
func main() {
	firstChan := inputDatafirst(1, 2, 3, 4)
	secondChan := inputDatasecond(firstChan)
	for item := range secondChan {
		log.Println("Receive,", item)
	}
	log.Println("Main finished")
}
