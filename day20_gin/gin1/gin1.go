package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/test", test)
	router.Run(":80")
}

func test(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}
