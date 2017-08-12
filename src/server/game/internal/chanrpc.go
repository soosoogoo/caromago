package internal

import (
	"fmt"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	//skeleton.RegisterChanRPC("CNewAgent", rpcCNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var users = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	users[a] = struct{}{}
	fmt.Println("上线回调2 ", len(users))
	fmt.Println(args)
}

func rpcCloseAgent(args []interface{}) {
	fmt.Println("离线回调2 ")
	a := args[0].(gate.Agent)
	delete(users, a)

	userInfo, ok := a.UserData().(string)
	if !ok {
		return
	}

	BroadcastMsg(&msg.ReturnMsg{
		Ststus: "aaa",
		Info:   userInfo,
	}, a)
}

func BroadcastMsg(msg interface{}, _a gate.Agent) {
	for a := range users {
		//		if a == _a {
		//			fmt.Println("nimayoudu ")
		//			continue
		//		}
		a.WriteMsg(msg)
	}
}
