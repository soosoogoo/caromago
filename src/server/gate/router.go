package gate

import (
	"server/game"
	"server/login"
	"server/msg"
	"server/room"
)

func init() {

	//登录模块
	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)

	//游戏模块
	msg.Processor.SetRouter(&msg.RoleLogin{}, game.ChanRPC)

	/**房间模块*/
	//创建房间
	msg.Processor.SetRouter(&msg.CreateRoom{}, room.ChanRPC)

	//房间列表
	msg.Processor.SetRouter(&msg.RoomList{}, room.ChanRPC)

}
