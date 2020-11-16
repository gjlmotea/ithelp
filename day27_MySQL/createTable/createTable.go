package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 與DB連線。 init() 初始化，時間點比 main() 更早。
func init() {
	dbConnect, err := sql.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/",
	)

	if err != nil {
		log.Fatalln(err)
	}

	err = dbConnect.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db = dbConnect // 用全域變數接

	db.SetMaxOpenConns(10) // 可設置最大DB連線數，設<=0則無上限
	db.SetMaxIdleConns(10) // 設置最大閒置連線數。
	// 更多用法可以 進 sql.DBStats{}、sql.DB{} 裡面看
}

func main() {
	// createDb("`school`") // 比較好的習慣會使用 `` 反引號（重音號）來區隔
	// deleteDb("`school`")
	// createTable1()
	// createTable2()
	// deleteTable1()
	// insertTeacher("Peter", 3)
	// queryTeachers()
	// alterStudent()
	// insertStudent("Ricky")
	// queryTeacherByAge(100)
	// deleteTeacher("QQ")
	// updateTeacherAge("Amy", 70)
	// oldestTeacher()
	// oldestAge()
	// numOfTeacher()
	SumOfTeachersAge()
}

func createDb(dbName string) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + ";")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(err)
}

func deleteDb(dbName string) {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbName + ";")
	if err != nil {
		log.Fatalln(err)
	}
}

func createTable1() {
	// SQL: CREATE TABLE IF NOT EXISTS `school.student`
	// 初始化Table 沒給任何欄位時，會出現 `A table must have at least 1 column` 的錯誤

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `school`.`student`(`name` VARCHAR (10))")
	if err != nil {
		log.Fatalln(err)
	}

	// 或者用 USE 來指定該資料庫
	// db.Exec("USE `school`")
	// _, err := db.Exec("CREATE TABLE IF NOT EXISTS `student`(`name` VARCHAR(10))")
}

func createTable2() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `school`.`teacher`(`name` VARCHAR (10) PRIMARY KEY, `age` INT (1))")
	if err != nil {
		log.Fatalln(err)
	}
}

func deleteTable1() {
	_, err := db.Exec("DROP TABLE IF EXISTS `school`.`student`")
	if err != nil {
		log.Fatalln(err)
	}
}

func insertTeacher(teacherName string, teacherAge int) {
	_, err := db.Exec("INSERT INTO `school`.`teacher`(`name` , `age`) VALUES (? , ?)", teacherName, teacherAge)
	if err != nil {
		log.Println(err)
	}
}

func insertStudent(studentName string) {
	rs, err := db.Exec("INSERT INTO `school`.`student`(`name`) VALUES (?)", studentName)
	if err != nil {
		log.Println(err)
	}

	rowCount, err := rs.RowsAffected() // 受影響的資料筆數
	rowId, err := rs.LastInsertId()    // 資料表中有Auto_Increment欄位才起作用，回傳剛剛新增的那筆資料ID

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("新增 %d 筆資料，id = %d \n", rowCount, rowId)
}

func alterStudent() {
	_, err := db.Exec("ALTER TABLE `school`.`student` ADD `id` INT AUTO_INCREMENT PRIMARY KEY;")
	if err != nil {
		log.Println(err)
	}
}

func queryTeachers() {
	rows, err := db.Query("SELECT `name`, `age` FROM `school`.`teacher`")
	// rows, err := db.Query("SELECT * FROM `school`.`teacher`") // 也可以使用`Select *`選取全部欄位。
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() { // rows.Next() 前往下一個項目。如果成功（還有下一項的話）返回True、失敗（沒有下一項可讀）則返回False
		var tName string //兩個欄位，依SELECT的順序用兩個變數來接
		var tAge int
		err = rows.Scan(&tName, &tAge) // 掃描後存進變數中
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%q %d\n", tName, tAge) // %q:quoted 用引號包起字串
	}
	defer rows.Close() // 當完整迭代rows.Next()完後會自動關閉rows，但以防萬一 最後記得要關閉rows 。
}

func queryTeacherByAge(age int) {
	rows, err := db.Query("SELECT `name`, `age` FROM `school`.`teacher` WHERE `age` > ?;", age)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var tName string
		var tAge int
		err = rows.Scan(&tName, &tAge)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%q %d\n", tName, tAge)
	}
	defer rows.Close()
}

func deleteTeacher(teacherName string) {
	_, err := db.Exec("DELETE FROM `school`.`teacher` WHERE `name` = ?;", teacherName)
	if err != nil {
		log.Println(err)
	}
}

func updateTeacherAge(teacherName string, age int) {
	_, err := db.Exec("UPDATE `school`.`teacher` SET `age`= ? WHERE `name` = ?;", age, teacherName)
	if err != nil {
		log.Println(err)
	}
}

func numOfTeacher() {
	row := db.QueryRow("SELECT COUNT(`age`) FROM `school`.`teacher`;")
	var count int
	row.Scan(&count)
	fmt.Println(count)
}

func oldestAge() {
	row := db.QueryRow("SELECT MAX(`age`) FROM `school`.`teacher`;")

	var tAge int
	row.Scan(&tAge)
	fmt.Println("老師年紀最大為:", tAge)
}

func oldestTeacher() {
	row := db.QueryRow("SELECT `name`, `age` FROM `school`.`teacher` ORDER BY `age` DESC Limit 1 ;")

	var tName string
	var tAge int
	row.Scan(&tName, &tAge)

	fmt.Println(tName, tAge)
}

func SumOfTeachersAge() {
	row := db.QueryRow("SELECT SUM(`age`) FROM `school`.`teacher`;")
	var sum int
	row.Scan(&sum)
	fmt.Println(sum)
}
