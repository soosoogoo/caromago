package sqlModule

import (
	"fmt"

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

type UserModel struct {
}

func (ac *UserModel) GetUserInfo() {
	db, err := gorm.Open("mysql", "root:WeiXin123@tcp(127.0.0.1:3306)/soosoogoo_main?charset=utf8&parseTime=False&loc=Local")
	defer db.Close()

	if err != nil {
		println(err.Error())
		return
	}

	// テーブル名が複数系でない場合、これを指定すること
	db.SingularTable(true)

	// 打印LOG
	db.LogMode(true)

	var user User
	row := db.First(&user, "user_id = ?", "1")

	//db.Model(&user).Update("sex", 1)
	//row.Scan(&user)

	if row != nil {
		fmt.Println("id:" + user.Username)
	}

	fmt.Println("a")
}
