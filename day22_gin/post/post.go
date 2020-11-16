package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.RedirectFixedPath = true
	router.POST("/post", post)
	router.Any("/any", any)

	router.Run(":80")
}

func post(c *gin.Context) {
	//msg := c.PostForm("input")
	msg := c.DefaultPostForm("input", "表單沒有input。") // 沒有輸入參數時 可設定預設值

	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)
}

func any(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
