package server

import (
	"awesomeProject1/per/jm/database/api/handler"
	"awesomeProject1/per/jm/database/api/model"
	"awesomeProject1/per/jm/database/api/status"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net"
)

type Connection struct {
	priHander handler.PrivateChatHandler
}

func (c *Connection) HandlerCon(con net.Conn) {
	defer con.Close()

	buff := make([]byte, 1024)
	l, err := con.Read(buff)
	//随机token
	token, _ := uuid.NewV4()
	login := handler.LoginHandler{}
	//数据读取没问题
	if err == nil {
		re, _ := login.Login(buff, con, l, token.String())
		if !re {
			fmt.Println("登录失败")
			return
		}
	}
	var m model.Message
	//循环监听新消息
	for {
		len, err := con.Read(buff)
		if err != nil {
			fmt.Scanln(err)
			continue
		}
		er := json.Unmarshal(buff[:len], &m)
		fmt.Println("监听消息")
		if er != nil {
			fmt.Println(er)
			continue
		}
		switch m.Code {
		case status.PRIVATE:
			{
			fmt.Println("私聊")
				c.priHander.Chat(m.Data, m)
			}
		case status.GROUP:
			{

			}
		case status.FRID_LIST:
			{
				handler.FriendListHandler{}.Handler(m.Data,m)
			}
		case status.HISTORY:
			{

			}
			//switch over
		}

		//for over
	}
}
