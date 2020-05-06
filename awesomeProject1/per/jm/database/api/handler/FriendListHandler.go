package handler

import (
	"awesomeProject1/per/jm/database/api/container"
	"awesomeProject1/per/jm/database/api/model"
	"awesomeProject1/per/jm/database/api/socket"
	"awesomeProject1/per/jm/database/api/status"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type FriendListHandler struct {

}

func (f FriendListHandler) Handler(body string,m model.Message){
	var fList model.Friend
	json.Unmarshal([]byte(body), &fList)
	if(fList.Type == status.FRID_LIST){
		//获取好友列表
	}else if(fList.Type == status.ADD_FRI){
		//添加好友
		uid := fList.Uid
		tokenInfo := container.UidMToken[string(uid)]
		addF(tokenInfo,m)
	}

}

func addF(tokenInfo string ,m model.Message){
	if tokenInfo != "" {
		m.Msg = "发起请求添加的请求"
		var listInfo []model.Friend
		listInfo = make([]model.Friend, 10)
		jbyte, _ := json.Marshal(listInfo)
		m.Data = string(jbyte)
		content, _ := json.Marshal(m)
		conn := container.Conns["uid"]
		//给自己发
		socket.OutputSocket{}.OutPut(conn, content)

	} else {
		m.Code = status.ADD_FRI
		m.Msg = "好友不存在"
		m.Data = ""
		content, _ := json.Marshal(m)
		conn := container.Conns["uid"]
		//给自己发
		socket.OutputSocket{}.OutPut(conn, content)
	}
}
var friendList map[string][]map[string]string //保存好友列表

func init() {
	//加载本地记录如何实现？
	friendList = make(map[string][]map[string]string, 10) //初始化大小为　１０
	loadLocalFriendInfo()
}

func (f FriendListHandler) Get(uid string) []map[string]string {
	return friendList[uid]
}

func (f *FriendListHandler) Add(name string, uid string, owner string) {
	friendList[owner] = append(friendList[owner],map[string]string{
		"name": name,
		"uid":  uid,
	})
}

func loadLocalFriendInfo() {
	listPath := "list.txt"
	// 文件不存在则返回error
	_, err := os.Stat(listPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(listPath)
		}
	}
	fi, err := os.Open(listPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	count := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if(count > 0) {
			strarray := strings.Fields(strings.TrimSpace(string(a)))
			fmt.Println(strarray)
			friendList[strarray[0]] = append(friendList[strarray[0]], map[string]string{
				"uid":strarray[1],
				"name":strarray[2],
			})
		}
		count++;
	}
}