package msg

import (
	"github.com/name5566/leaf/network/json"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Hello{})
	Processor.Register(&RoleLogin{})

	//test
	Processor.Register(&S2C_Left{})

	//通用
	Processor.Register(&ReturnMsg{})
}

// 一个结构体定义了一个 JSON 消息的格式
type ReturnMsg struct {
	Ststus interface{}
	Info   interface{}
}

// 消息名为 Hello
type Hello struct {
	Name string
}

type Login struct {
	Name string
}

type RoleLogin struct {
	Role_id string
}

type S2C_Left struct {
	NumUsers int
	UserName string
}
