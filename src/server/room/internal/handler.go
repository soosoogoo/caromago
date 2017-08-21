package internal

import (
	"fmt"
	"reflect"
	"server/base"
	"server/conf"
	"server/model/room"
	"server/msg"

	"github.com/name5566/leaf/gate"
	//	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.CreateRoom{}, handleCreateRoom)
	handler(&msg.RoomList{}, handleRoomList)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/**
创建房间
*/
func handleCreateRoom(args []interface{}) {

	//收到的消息,消息的发送者
	m := args[0].(*msg.CreateRoom)
	a := args[1].(gate.Agent)

	//查询用户当前状态

	//UserData
	//var userInfo = a.UserData()
	//查询用户是否已经有房间

	//创建房间
	var ur = new(room.RoomModel)
	var RoomInfo = ur.GreateRoom(1, m)

	//fmt.Println(a.UserData().(base.UserInfo).UserId)

	//将创建的room_id 加入  user_info
	//a.SetUserData(userInfo)

	var userInfo = a.UserData()
	base.AddToGroup(userInfo.(base.UserInfo).UserId, conf.ROOMGROUP, RoomInfo.Room_id)

	a.WriteMsg(&msg.ReturnMsg{
		Ststus: "asdasd",
		Info:   RoomInfo,
	})
	//返回房间信息

}

/**
房间列表
*/
func handleRoomList(args []interface{}) {
	//收到的消息,消息的发送者
	m := args[0].(*msg.RoomList) //消息
	a := args[1].(gate.Agent)    //链接

	fmt.Println(m)
	fmt.Println(a)
	//查询房间信息
	var ur = new(room.RoomModel)
	var RoomList = ur.RoomList(m)
	fmt.Println(RoomList)
	//返回房间信息
	a.WriteMsg(&msg.ReturnMsg{
		Ststus: "asdasd",
		Info:   RoomList,
	})
}
