package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {

	msg.Processor.SetRouter(&msg.RoleLogin{}, game.ChanRPC)

	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)

}
