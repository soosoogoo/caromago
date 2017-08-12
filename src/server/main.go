package main

import (
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/room"

	//	"server/model/user"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {

	//var md = new(mysql.MysqlDriver)
	//var userInfo = md.Test()

	//fmt.Println(conf.Server.LogLevel)
	//fmt.Println(conf.MAINDATABASE)

	//var md = new(user.UserModel)
	//var username = md.GetUserInfo(1)
	//fmt.Println(username)

	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
		room.Module,
	)

}
