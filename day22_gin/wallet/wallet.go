package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance = 0

func main() {
	router := gin.Default()
	router.GET("/deposit/:input", deposit)
	router.GET("/withdraw/:input", withdraw)
	router.GET("/balance/", getBalance)

	router.Run(":80")
}

func getBalance(context *gin.Context) {
	var msg = "您的錢包裡有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})
}

func deposit(context *gin.Context) {
	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	status := "status init"
	msg := "msg init"
	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "操作失敗，存款金額需大於0元！"
		} else {
			balance += amount
			status = "ok"
			msg = "已成功存款" + strconv.Itoa(amount) + "元"
		}
	} else {
		amount = 0
		status = "failed"
		msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}

func withdraw(context *gin.Context) {
	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	status := "status init"
	msg := "msg init"
	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "操作失敗，提款金額需大於0元！"
		} else {
			if balance-amount < 0 {
				amount = 0
				status = "failed"
				msg = "操作失敗，餘額不足！"
			} else {
				balance -= amount
				status = "ok"
				msg = "成功提款" + strconv.Itoa(amount) + "元"
			}
		}
	} else {
		amount = 0
		status = "failed"
		msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}
