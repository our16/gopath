package main

import (
	"client/per/status"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"
)

var ip = flag.String("IP", "127.0.0.1", "ip 地址")
var uid = flag.String("u", "", "好友id")
var pwd = flag.String("p", "", "密码")

var uid2 = flag.String("u2", "", "好友id")

func main() {
	test()
}
func test() {
	flag.Parse()
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//catch exception
	defer func() {
		conn.Close()
		fmt.Println("b")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()

	// panic("a bug occur") throw exception
	_, token := loginHandler(conn, *uid, *pwd)
	go showInfo(conn)
	fmt.Println("to :", *uid2)
	whi := write(token, *uid2)
	start := true
	for start {
		whi(conn)
	}
}
func loginHandler(con net.Conn, uid string, pwd string) (bool, string) {

	loginInfo := status.LoginInfo{
		uid,pwd,
	}
	loginInfoJsonb,_ :=json.Marshal(loginInfo)
	//登录消息
	content := status.Message{
		"登录认证",
		status.LOGIN,
		string(loginInfoJsonb),
	}
	data, _ := json.Marshal(content)
	// 发送数据
	con.Write(data)
	buff := make([]byte, 1024)
	len, _ := con.Read(buff)
	var mes = &status.Message{}
	//json 解析这里必须指定切片长度
	json.Unmarshal(buff[:len], &mes)
	token := mes.Data
	return true, token
}
func showInfo(con net.Conn) {
	start := true
	var m status.Message
	for start {
		buff := make([]byte, 1024)
		len, _ := con.Read(buff)
		json.Unmarshal(buff[:len], &m)
		switch m.Code {
		case status.PRIVATE:
			{
			  priChat := status.PrivateChat{}
				json.Unmarshal([]byte(m.Data),&priChat)
				fmt.Println("----------------------------------------------")
				fmt.Println(priChat.From)
				fmt.Println(priChat.Msg)
				fmt.Println(priChat.Date)
				fmt.Println("-----------------------------------------------")
			}
		case status.GROUP:
			{

			}
		case status.FRID_LIST:
			{
				fmt.Println("获取到好友列表:")
				//ls := m.Data
				//list := status.FriendList{}
				//list.ShowList(ls)
			}
		case status.HISTORY:
			{

			}

		}

	}
}

//闭包,共用　token
func write(token string, to string) func(con net.Conn) {
	return func(con net.Conn) {
		pri := status.PrivateChat{}
			pri.Date = time.Now().String()
			pri.From = *uid
			pri.To = to
		mes := status.Message{
			"通信消息",
			status.PRIVATE, //私聊\
			"",
		}
		for {
			var content string
			fmt.Scanln(&content)
			 pri.Msg = content
			 da,_:=json.Marshal(pri)
			 mes.Data = string(da)
			 cont, _ := json.Marshal(mes)
			con.Write([]byte(cont))
		}
	}

}
