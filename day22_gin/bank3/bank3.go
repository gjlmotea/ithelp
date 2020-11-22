package main

import (
	"errors"
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

// getBalance 取得帳戶內餘額
func getBalance(context *gin.Context) {
	wrapResponse(context, balance, nil) // 操作成功，返回Amount為balance，返回err為nil
}

// deposit 儲值、存款
func deposit(context *gin.Context) {
	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			wrapResponse(context, 0, errors.New("操作失敗，存款金額需大於0元！"))
			return
			// 按理說如果程式執行中出現錯誤，為避免程式繼續往下執行 會在這裡加上 `return`
			// 但因為程式到這裡已經結束，底下沒有任何程式碼了，所以可不用寫 `return`
			// 以下同理
		} else {
			balance += amount
			wrapResponse(context, balance, nil) // 操作成功，返回Amount為balance，返回err為nil
		}
	} else {
		wrapResponse(context, 0, errors.New("操作失敗，輸入有誤！"))
	}
}

// withdraw 提款
func withdraw(context *gin.Context) {
	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			wrapResponse(context, 0, errors.New("操作失敗，提款金額需大於0元！"))
		} else {
			if balance-amount < 0 {
				wrapResponse(context, 0, errors.New("操作失敗，餘額不足！"))
			} else {
				balance -= amount
				wrapResponse(context, balance, nil) // 操作成功，返回Amount為balance，返回err為nil
			}
		}
	} else {
		wrapResponse(context, 0, errors.New("操作失敗，輸入有誤！"))
	}
}

func wrapResponse(context *gin.Context, amount int, err error) {
	var r = struct {
		Amount  int    `json:"amount"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Amount:  amount,
		Status:  "ok", // 預設狀態為ok
		Message: "",
	}

	if err != nil {
		r.Amount = 0
		r.Status = "failed"     // 若出現任何err，狀態改為failed
		r.Message = err.Error() // Message回傳錯誤訊息
	}

	context.JSON(http.StatusOK, r)
}
