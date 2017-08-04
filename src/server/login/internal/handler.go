package internal

import (
	"reflect"
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

func handleRoleLogin(args []interface{}) {

	//查询用户信息

}
