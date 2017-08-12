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
	AllConn  map[string]*gorm.DB
}

/**
初始化数据库连接
*/
func (c *MysqlDriver) Connent(databae string) *gorm.DB {

	fmt.Println(databae)
	var db *gorm.DB
	if databae == conf.MAINDATABASE {
		db, _ = gorm.Open("mysql", conf.Mysql.Username+":"+conf.Mysql.Password+"@tcp("+conf.Mysql.Hostname+":3306)/"+databae+"?charset=utf8&parseTime=False&loc=Local")
	} else {
		m := c.GetMysqlDevice(databae)
		db, _ = gorm.Open("mysql", m.Username+":"+m.Password+"@tcp("+m.Hostname+":3306)/"+databae+"?charset=utf8&parseTime=False&loc=Local")
	}

	//优先使用
	db.SingularTable(true)

	// 打印LOG
	db.LogMode(true)

	return db
}

/**
获取mysql配置信息
*/
type MysqlDevice struct {
	server_id int `gorm:"primary_key"`

	Dbdriver string //必须要大些才能在外部调用
	Hostname string
	Username string
	Password string
	Database string
	Dbprefix string
	Pconnect int
}

func (d MysqlDevice) TableName() string {
	return "soosoogoo_db_device"
}

//获取配置的mysql设置
func (c *MysqlDriver) GetMysqlDevice(server_id string) MysqlDevice {

	db, err := gorm.Open("mysql", conf.Mysql.Username+":"+conf.Mysql.Password+"@tcp("+conf.Mysql.Hostname+":3306)/"+conf.Mysql.Database+"?charset=utf8&parseTime=False&loc=Local")

	if err != nil {
		println(err.Error())
	}

	db.LogMode(true)

	var mysqlDevice MysqlDevice

	row := db.First(&mysqlDevice, "server_id = ?", server_id)

	if row != nil {
		//fmt.Println(mysqlDevice.server_id)
	}
	return mysqlDevice
}
