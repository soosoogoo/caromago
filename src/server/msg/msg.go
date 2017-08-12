package msg

import (
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

var (
	JSONProcessor     = json.NewProcessor()
	ProtobufProcessor = protobuf.NewProcessor()
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {

	//login
	Processor.Register(&Login{})

	Processor.Register(&RoleLogin{})

	//room
	Processor.Register(&CreateRoom{})

	//通用
	Processor.Register(&ReturnMsg{})
}

// 返回值结构
type ReturnMsg struct {
	Ststus interface{}
	Info   interface{}
}

// room信息
type CreateRoom struct {
	Stage_id int
	Type     int
	Ut_id    int
	Is_open  int
}

type Login struct {
	RoleId int
}

type RoleLogin struct {
	RoleId int
}
