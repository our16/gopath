package socket

import (
	"fmt"
	"net"
)

type OutputSocket struct {

}

func (out OutputSocket)OutPut(con net.Conn,data []byte){
	if nil == con {
		fmt.Println("to: ", con)
		return
	}
	if data != nil {
		con.Write(data)
	}
}
