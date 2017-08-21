package game

import (
	"server/base"
	"server/game/internal"

	"github.com/name5566/leaf/gate"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

//调用消息群发
func PBroadcastMsg(msg interface{}, _a gate.Agent) {
	internal.BroadcastMsg(msg, _a)
}

//用户登录
func BindUid(uid int, _a gate.Agent) {
	base.BindUid(uid, _a)
}
