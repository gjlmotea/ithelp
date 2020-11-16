package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //為什麼要用底線_來占著位子？
)

func main() {
	dbConnect, err := sql.Open( // 建立與MySQL Server端的連線
		"mysql", // 因為只有這裡才用到這個`引數`，並沒有直接使用到了mysql.XXX 相關的函式或物件，會被認為沒有用到mysql這個依賴而 被go編譯器省略import
		"root:root@tcp(127.0.0.1:3306)/",
	)

	if err != nil {
		log.Fatalln(err)
	}

	err = dbConnect.Ping() //Ping() 這裡才開始建立連線。上面 sql.Open 只建立物件、容器，並未進行連線，無法連線並不造成err。
	if err != nil {
		log.Fatalln(err)
	}
}
