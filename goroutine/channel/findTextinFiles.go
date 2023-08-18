package main

import (
	"log"
	"os"
	"strings"
)

func readFileA(filePath string, text string, sumWord chan int) {
	var numberWord int

	//doc file
	fileContent, err := os.ReadFile(filePath) // tra ve noi dung file la mang kieu byte
	if err != nil {                           // neu co loi
		log.Println(err) // in ra loi
		sumWord <- 0     //tra ve khong co tu
		return
	}
	//dem so tu
	numberWord = strings.Count(string(fileContent), text) // do tra ve mang kieu byte nen can chuyen ve kieu string
	//gui vao channel
	sumWord <- numberWord
	defer close(sumWord)

}
func readFileB(filePath string, text string, sumWord chan int) {
	var numberWord int
	//doc file
	fileContent, err := os.ReadFile(filePath) // tra ve noi dung file la mang kieu byte
	if err != nil {                           // neu co loi
		log.Println(err) // in ra loi
		sumWord <- 0     //tra ve khong co tu
		return
	}
	//dem so tu
	numberWord = strings.Count(string(fileContent), text) // do tra ve mang kieu byte nen can chuyen ve kieu string
	//gui vao channel
	sumWord <- numberWord
	defer close(sumWord)

}
func main() {
	countWordChanA := make(chan int)
	countWordChanB := make(chan int)

	go readFileA("\\LearnGolang\\tutorial\\lessions\\goroutine\\channel\\a.txt", "ab", countWordChanA)
	go readFileB("\\LearnGolang\\tutorial\\lessions\\goroutine\\channel\\b.txt", "ab", countWordChanB)
	log.Println("Tong so lan xuat hien cua tu trong file: ", <-countWordChanA+<-countWordChanB)
	log.Println("\nFinished all.")
}
