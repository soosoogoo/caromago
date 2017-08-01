package gosql

import (
	"fmt"

	"github.com/hfdend/gosql"
)

func main() {
	// 建立连接
	m := gosql.NewDbMysql("192.168.0.197", 3306, "root", "WeiXin123", "test")
	// 设置最大连接数
	//m.SetMaxOpenConns(30)
	// SetMaxIdleConns sets the maximum number of connections in the idle
	//m.SetMaxIdleConns(10)
	// 设置空闲连接池的生存时间
	//m.SetAutoCloseTime(100)
	// 设置表名
	m.SetTableName("sc")

	// 数据插入
	data := map[string]interface{}{
		"S#":    "1",
		"C#":    "1",
		"score": 11,
	}
	id, e := m.Insert(data)
	fmt.Println(e)
	fmt.Println(id)
	fmt.Println(m.LastSql)
}
