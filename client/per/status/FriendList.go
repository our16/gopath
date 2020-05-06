package status

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type FriendList struct {
	Uid string
}

var friendList map[string][]map[string]string //保存好友列表

func init() {
	//加载本地记录如何实现？
	friendList = make(map[string][]map[string]string, 10) //初始化大小为　１０
	loadLocalFriendInfo()
}

func (f FriendList) Get(uid string) []map[string]string {
	return friendList[uid]
}

func (f *FriendList) Add(name string, uid string, owner string) {
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
