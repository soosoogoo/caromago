package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

var Mysql struct {
	Hostname string
	Username string
	Password string
	Database string
	Dbdriver string
	Dbprefix string
	Pconnect int
	Db_debug int
	Cache_on int
	Cachedir string
	Char_set string
	Dbcollat string
	Swap_pre string
	Autoinit int
	Stricton int
}

var RedisConfig struct {
	Host     string
	Port     int
	Prefix   string
	Username string
	Pwd      string
}

//初始化
func init() {
	loadConfig("conf/server.json", &Server)
	loadConfig("conf/mysql.json", &Mysql)
	loadConfig("conf/redis.json", &RedisConfig)
}

func loadConfig(configPath string, structConfig interface{}) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &structConfig)
	if err != nil {
		log.Fatal("%v", err)
	}
}
