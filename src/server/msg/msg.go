package msg

import (
	"github.com/name5566/leaf/network/json"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {

	//login
	Processor.Register(&Login{})

	Processor.Register(&RoleLogin{})

	//通用
	Processor.Register(&ReturnMsg{})
}

// 一个结构体定义了一个 JSON 消息的格式
type ReturnMsg struct {
	Ststus interface{}
	Info   interface{}
}

type Login struct {
	//UserName string
	//Password string
	RoleId int
}

type RoleLogin struct {
	RoleId int
}

type S2C_Left struct {
	NumUsers int
	UserName string
}
