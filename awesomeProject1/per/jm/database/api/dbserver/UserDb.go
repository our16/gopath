package dbserver

import (
	"fmt"
	"log"
)

type UserDb struct {

}

func (d UserDb)SelectUserInfoById(id int64)(string,string,int64){
	var name string
	var pwd string
	rows, err := DbConn.Query("select name,pwd from user where id = ? ", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &pwd)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(name,pwd)
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	return name,pwd,id
}

func (d UserDb)Insert(name string,pwd string) int64{
	stmt, err := DbConn.Prepare("insert into user(name,pwd)values(?,?)")
	defer  stmt.Close()
	if err != nil {
		log.Println(err)
		fmt.Println("insert prepare err")
	}
	rs, err := stmt.Exec(name, pwd)
	if err != nil {
		log.Println(err)
	}
	//我们可以获得插入的id
	id, err := rs.LastInsertId()
	fmt.Println("id",id)
	return id
}

func (d UserDb)Update(){
	stmt, err := DbConn.Prepare("update user set user.name = ? , pwd = ? where id = ? ")
	defer  stmt.Close()
	if err != nil {
		fmt.Println("insert prepare err")
	}
	rs, err := stmt.Exec("petter", "1asfasf23",3)
	if err != nil {
		log.Println(err)
		return
	}
	num,_ := rs.RowsAffected()
	fmt.Println("影响行数:",num)

}

func (d UserDb)Delete(){
	stmt, _ := DbConn.Prepare("delete from user where id = ?")
	defer stmt.Close()
	stmt.Exec(2)
}
