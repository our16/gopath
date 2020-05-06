package dbserver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
type DbServer struct {

}
var DbConn *sql.DB

func init() {
	fmt.Println("加载mysql 驱动及其连接")
	user := "root"
	pwd :="000000"
	ip := "localhost"
	port :="3306"
	database := "test"
	addtress := user+":"+pwd+"@tcp("+ip+":"+port+")/"+database
	db, err := sql.Open("mysql", addtress)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(40)
	if err != nil {
		log.Println(err)
	}
	//在这里进行一些数据库操作
	//defer db.Close()
	DbConn = db
}

