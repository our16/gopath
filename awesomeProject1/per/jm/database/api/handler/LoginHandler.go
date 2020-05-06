package handler

import (
	"awesomeProject1/per/jm/database/api/container"
	"awesomeProject1/per/jm/database/api/dbserver"
	"awesomeProject1/per/jm/database/api/model"
	"awesomeProject1/per/jm/database/api/socket"
	"awesomeProject1/per/jm/database/api/status"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type LoginHandler struct {

}

func (this LoginHandler) Login(data []byte, con net.Conn, len int, token string) (bool,int){
	var m model.Message
	err := json.Unmarshal(data[:len], &m)
	if err != nil {
		fmt.Println(err.Error())
		return false,status.NOLOGIN
	}
	var flag = false
	if m.Code == status.LOGIN {
		fmt.Println("请求登录")
		loginInfo := model.LoginInfo{}
		json.Unmarshal([]byte(m.Data), &loginInfo)
		uid,_ := strconv.ParseInt(loginInfo.Uid, 10, 64)
		var userDb dbserver.UserDb
		name,pwd,_ := userDb.SelectUserInfoById(uid);
		m2 := &model.Message{
			"完成登录", status.LOGIN, token,
		}

		if pwd != loginInfo.Pwd {
			m2.Data = ""
			m2.Msg = "密码错误"
			flag = false
		}else {
			container.UidMToken[loginInfo.Uid] = token;
			container.Conns[token] = con
			userContent := model.NewUserVo(name)
			container.UserContent[token] = userContent
			flag = true
		}
		da, _ := json.Marshal(m2)
		socket.OutputSocket{}.OutPut(con, da)
	}
	fmt.Println("返回登录结果")
	return flag,status.LOGIN
}