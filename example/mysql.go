package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "red:123@tcp(127.0.0.1:3306)/po?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接错误:",err)
	}
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(7*time.Hour)
	err= db.Ping()
	if err != nil {
		fmt.Println("数据库ping错误:",err)
	}
	fmt.Println(db.Query("select  now()"))
	defer db.Close()
}
