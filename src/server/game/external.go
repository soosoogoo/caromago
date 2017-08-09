package game

import (
	"server/game/internal"

	"github.com/name5566/leaf/gate"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

//调用
func PBroadcastMsg(msg interface{}, _a gate.Agent) {
	internal.BroadcastMsg(msg, _a)
}
