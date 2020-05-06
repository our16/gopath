package status

const (
	LOGIN = iota
	OK = 200
	FALSE = 500
	NOLOGIN = 40001
	PRIVATE = 301
	GROUP = 302
	//历史消息列表
	HISTORY = 257
	//好友列表
	FRID_LIST = 258
	//退出
	EXIT = -1
	//自己吗？
	SELF = true
	ADD_FRI = 259
)



type Status struct {
}
