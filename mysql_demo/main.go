package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "Jwt@1234"
	ipAddrees string = "172.18.63.145"
	port      int    = 3306
	dbName    string = "crawler"
	charset   string = "utf8"
)

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func updateRecord(Db *sqlx.DB) {
	//更新uid=1的username
	result, err := Db.Exec("update userinfo set username = 'anson' where uid = 1")
	if err != nil {
		fmt.Printf("update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

func addRecord(Db *sqlx.DB) {
	result, err := Db.Exec("insert into proxy(proxy,checkDate)  values(?,?)", "172.18.63.111:8080", time.Now())
	if err != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("insert success, last id:[%d]\n", id)
}

func ping(Db *sqlx.DB) {
	err := Db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	} else {
		fmt.Println("ping success")
	}
}

func main() {
	var Db *sqlx.DB = connectMysql()
	defer Db.Close()
	addRecord(Db)
	// ping(Db)
}
