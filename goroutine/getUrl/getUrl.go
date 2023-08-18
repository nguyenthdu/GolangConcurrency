package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func getUrl(urlPath string, result chan string) {
	resp, errGetUrl := http.Get(urlPath)
	if errGetUrl != nil {
		log.Println(errGetUrl)
		return
	}
	urlContent, errReadBody := io.ReadAll(resp.Body) // doc trong body cua resp va tra ve 1 mang byte
	if errReadBody != nil {
		log.Println(errGetUrl)
		return
	}
	result <- string(urlContent)
	defer resp.Body.Close()
}

func main() {
	firstChanUrl := make(chan string)
	secondChanUrl := make(chan string)
	thirdChanUrl := make(chan string)
	f, err := os.OpenFile("tutorial/lessions/goroutine/getUrl/saveUrl.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return
	}
	listChanUrl := []chan string{firstChanUrl, secondChanUrl, thirdChanUrl}                   //channel dung de luu du lieu
	listUrlGet := []string{"http://youtube.com", "http://google.com", "http://vnexpress.net"} //danh sach url dung de doc
	for i := 0; i < len(listUrlGet); i++ {
		go getUrl(listUrlGet[i], listChanUrl[i]) //dung goroutine de lay du lieu tu url luu vao channel
	}
	for i := 0; i < len(listChanUrl); i++ {
		_, err := f.WriteString(<-listChanUrl[i]) // doc du lieu tu channel luu vao file
		if err != nil {
			log.Println(err)
			return
		}
	}
	log.Println("Main Finished")
}
