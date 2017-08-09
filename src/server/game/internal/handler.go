package internal

import (
	"reflect"
	"server/model/user"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.RoleLogin{}, handleRoleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/**
用户登录
*/
func handleRoleLogin(args []interface{}) {
	//获取用户信息

	// 收到的 消息
	m := args[0].(*msg.RoleLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)

	//查询用户信息
	var md = new(user.UserModel)
	var userInfo = md.GetUserInfo(m.RoleId)

	// 输出收到的消息的内容
	log.Debug("role_id %v", m.RoleId)

	//添加进chan

	a.SetUserData(userInfo)
	BroadcastMsg(&msg.ReturnMsg{
		Ststus: "asdasd",
		Info:   userInfo,
	}, a)
	//skeleton.AsynCall(game.ChanRPC, "NewAgent", args)

}
