package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe(":80", nil)
	// 傾聽 TCP 80 port及 處理服務
}

func test(writer http.ResponseWriter, request *http.Request) {
	// 把字串寫回writer
	str := []byte("ok")
	writer.Write(str)

	// 或者直接用以下這行
	fmt.Fprintf(writer, "123")
}
