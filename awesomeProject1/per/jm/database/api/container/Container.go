package container

import (
	"awesomeProject1/per/jm/database/api/model"
	"net"
)

type Container struct {

}
var Conns = make(map[string]net.Conn)
var UidMToken = make(map[string]string)
//指针可以读写，非指针只能读
var UserContent = make(map[string] *model.UserVo)