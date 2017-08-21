package user

import (
	//"fmt"
	"server/conf"
	"server/db/mysql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserData struct {
	UserId int
}

type User struct {
	UserId int `gorm:"primary_key"`

	Username string //必须要大些才能在外部调用
	sex      int
}

func (d User) TableName() string {
	return "soosoogoo_users"
}

/**
user info 查询结构体
*/
type UserModel struct {

	//继承mysql.MysqlDriver
	mysql.MysqlDriver
}

type UserInfo struct {

	//继承mysql.MysqlDriver
	UserId int
}

//#TODO::构造函数怎么写

//查询用户信息
func (u *UserModel) GetUserInfo(RoleId int) User {
	var user User
	var uDB = u.Connent(conf.MAINDATABASE)
	uDB.First(&user, "user_id = ?", RoleId)

	return user
}
