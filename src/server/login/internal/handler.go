package internal

import (
	"reflect"
	"server/db/mysql"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.RoleLogin{}, handleRoleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/**
通过role_id获取用户信息
	// 收到的 Hello 消息
	m := args[0].(*msg.RoleLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.RoleLogin{
		Name: "client",
	})
*/
func handleRoleLogin(args []interface{}) {
	//获取用户信息

	// 收到的 消息
	m := args[0].(*msg.RoleLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)

	//查询用户信息
	var md = new(mysql.MysqlDriver)
	var username = md.Test()

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Role_id)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.ReturnMsg{
		Ststus: "aaa",
		Info:   username,
	})

}
