package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/json", returnJson)
	router.GET("/json2", returnJson2)
	router.GET("/json3", returnJson3)

	router.Run(":80")
}

func returnJson(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func returnJson2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"狀態": "ok",
	})
}

func returnJson3(c *gin.Context) {
	type Result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	var result = Result{
		Status:  "OK",
		Message: "This is Json",
	}

	c.JSON(http.StatusOK, result)
}
