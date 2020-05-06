package dbserver

import (
	"awesomeProject1/per/jm/database/api/model"
	"awesomeProject1/per/jm/database/api/status"
	"strconv"
)

type FriendDb struct {

}

func(f FriendDb)Insert(uid string,fUid string){
	stmt,_ := DbConn.Prepare("insert into friend(uid,fuid) values(?,?) ")
	uid64,_ := strconv.ParseInt(uid,10,64)
	fUid64,_ := strconv.ParseInt(fUid,10,64)
	stmt.Exec(uid64,fUid64)
	defer stmt.Close()
}

func(f FriendDb)SelectByUid(uid string) []model.Friend{
	uid64,_ := strconv.ParseInt(uid,10,64)
	re,_ :=DbConn.Query("select name,fuid from friend where uid = ?",uid64)
	defer re.Close()
	friendLs := []model.Friend{}
	friendLs = make([]model.Friend,1)
	var name string
	var uidint int64
	for re.Next(){
	 fls := &model.Friend{}
	 re.Scan(&name,&uidint)
	 fls.Uid = uidint
	 fls.Name = name;
	 fls.Type = status.FRID_LIST
	 friendLs = append(friendLs,*fls)
	}
	return friendLs
}
