package internal

import (
	"reflect"
	"server/base"
	"server/model/user"
	"server/msg"

	"server/game"

	"github.com/name5566/leaf/gate"
	//"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.Login{}, handleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/**
用户登录
*/
func handleLogin(args []interface{}) {
	//获取用户信息

	// 收到的 消息
	m := args[0].(*msg.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)

	//查询用户信息
	var md = new(user.UserModel)
	var userInfo = md.GetUserInfo(m.RoleId)

	// 输出收到的消息的内容
	//log.Debug("role_id %v", m.RoleId)

	//添加进chan

	var ud base.UserInfo
	ud.UserId = m.RoleId
	a.SetUserData(ud)

	//base.BindUid(m.RoleId, a)
	game.PBroadcastMsg(&msg.ReturnMsg{
		Ststus: "asdasd",
		Info:   userInfo,
	}, a)

	//game.ChanRPC.Go("CNewAgent", m, a)
}
