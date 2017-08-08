package mysql

import (
	"fmt"
	"server/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	user_id int `gorm:"primary_key"`

	Username string //必须要大些才能在外部调用
	sex      int
}

func (d User) TableName() string {
	return "soosoogoo_users"
}

//定义链接
type MysqlDriver struct {
	database string
	DbConn   *gorm.DB
	//AllConn  map[string]*gorm.DB
}

/**
初始化数据库连接
*/
func (c *MysqlDriver) Connent(databae string) (um *gorm.DB) {

	db, err := gorm.Open("mysql", conf.Mysql.Username+":WeiXin123@tcp(127.0.0.1:3306)/"+databae+"?charset=utf8&parseTime=False&loc=Local")
	//defer db.Close()

	//c.AllConn = make(map[string]*gorm.DB)
	//if _, ok := c.AllConn[databae]; !ok {
	//	c.DbConn = db
	//	c.AllConn[databae] = db
	//	c.database = databae
	//}

	if err != nil {
		println(err.Error())
	}
	c.DbConn = db
	um = db
	//优先使用
	um.SingularTable(true)

	// 打印LOG
	um.LogMode(true)

	return
}

func (c *MysqlDriver) Test() User {

	c.Connent("soosoogoo_main")
	var user User

	row := c.DbConn.First(&user, "user_id = ?", "1")

	if row != nil {
		fmt.Println(user.Username)
	}
	return user
}
