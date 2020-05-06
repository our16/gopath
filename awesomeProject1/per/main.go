package main

import (
	"awesomeProject1/per/jm/database/api/dbserver"
	"awesomeProject1/per/jm/database/api/server"
	"fmt"
	"net"
)
//加载数据库
var initDb dbserver.DbServer

func main() {

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	conn := &server.Connection{}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go conn.HandlerCon(c)
	}
}