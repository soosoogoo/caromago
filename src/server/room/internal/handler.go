package internal

import (
	"reflect"
	"server/game"
	"server/model/room"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.CreateRoom{}, handleCreateRoom)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/**
创建房间
*/
func handleCreateRoom(args []interface{}) {

	//收到的消息
	m := args[0].(*msg.CreateRoom)
	// 消息的发送者
	a := args[1].(gate.Agent)

	//查询用户当前状态

	//查询用户是否已经有房间

	//创建房间
	var ur = new(room.RoomModel)
	var RoomInfo = ur.GreateRoom(m)

	// 输出收到的消息的内容
	log.Debug("stage_id %v", m.Stage_id)

	//添加进chan
	//a.SetUserData(userInfo)

	game.PBroadcastMsg(&msg.ReturnMsg{
		Ststus: "asdasd",
		Info:   RoomInfo,
	}, a)
	//返回房间信息

}
