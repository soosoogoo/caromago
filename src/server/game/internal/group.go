package internal

import (
	"fmt"
	"server/db/sredis"

	"github.com/garyburd/redigo/redis"
	//"server/msg"

	"server/base"
	"server/conf"

	"github.com/name5566/leaf/gate"
)

//用户信息
type UserInfo struct {
	Info      map[string]string
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
	var info UserInfo
	info.status = conf.FREEGROUP
	a.SetUserData(info)

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

	//获取当前集合
	userList, _ := redis.Bytes(rd.Hget(group, indexKey))

	if userList == nil {
		var ul = make(map[int]int)
		ul[uid] = 1
		rd.Hset(group, indexKey, base.MapToString(ul))
	} else {
		var ul = make(map[int]int)
		ul = base.StringToMap(userList)
		ul[uid] = 1
		rd.Hset(group, indexKey, base.MapToString(ul))
	}

	//储存信息

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
