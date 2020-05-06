package handler

import (
	"awesomeProject1/per/jm/database/api/container"
	"awesomeProject1/per/jm/database/api/model"
	"awesomeProject1/per/jm/database/api/socket"
	"encoding/json"
	"fmt"
)

type PrivateChatHandler struct {

}
/**
body 是josn 格式，
需要解析出接收对象
*/
func(this PrivateChatHandler) Chat(body string,m model.Message){
	connMaps := container.Conns
	tokenMaps := container.UidMToken
	pri := model.PrivateChat{}
	//解析数据body
	json.Unmarshal([]byte(body),&pri)
	//接收对象
	to := pri.To
	//获取好友链接
	token := tokenMaps[to]
	conn := connMaps[token]
	//不需要修改消息内容
	m.Data = body
	content, _ := json.Marshal(m)

	fmt.Println("好友连接:",conn)
	//发送方法发送，解决软循环引用
	out := socket.OutputSocket{}
	out.OutPut(conn,content)
}