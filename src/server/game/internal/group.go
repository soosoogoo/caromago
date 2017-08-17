package internal

import (
	"fmt"
	"server/db/sredis"
	//"server/msg"

	"server/conf"

	"github.com/name5566/leaf/gate"
)

//用户信息
type UserInfo struct {
	userinfo  map[string]string
	group     map[string]string
	room_id   int
	battle_id int
	status    string
}

var UsersList = make(map[int]gate.Agent)

//用户登录 绑定用户信息
func BindUid(uid int, a gate.Agent) {
	//	rdb, err := sredis.Connent("main")
	//	if err != nil {
	//		fmt.Println("Connect to redis error", err)
	//	}
	UsersList[uid] = a

	//设置用户信息
	var userinfo UserInfo
	userinfo.status = conf.FREEGROUP
	a.SetUserData(userinfo)

	fmt.Println(a.UserData())
}

//解除绑定
func UnBindUid(uid int, a gate.Agent) {
	//	rdb, err := sredis.Connent("main")
	//	if err != nil {
	//		fmt.Println("Connect to redis error", err)
	//	}
	delete(UsersList, uid)

	userInfo, ok := a.UserData().(string)
	if !ok {
		return
	}
	fmt.Println(userInfo)

}

//将用户添加到group
func AddToGroup(uid int, group string, indexKey int) {

	var rd = sredis.RedisDriver{Database: conf.MAINREDISDATABASE}

	userList, err = rd.Hset(group, indexKey, userList)

	if userList == nil {
		var userList = []int32{uid}
	} else {
		append(userList, uid)
	}
	rd.Hset(group, indexKey, userList)

	//groupUser, _ := rd.Hget(channel, indexKey)

	//获取人数
	//	_, err = c.Do("HGET", channel, indexKey)
	//	_, err = c.Do("HSET", channel, indexKey, uid)
}

//从group移除uid
func RemoveFromGroup(uid int, channel string) {

}

//删除group
func DelGroup(uid int, channel string) {

}
