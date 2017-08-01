package main

import (
	"encoding/binary"
	"math/rand"
	"net"

	"github.com/name5566/leaf/log"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go writeFromServer(conn)
	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
		"Hello": {
			"Name": "hello,world"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)

	log.Debug("asdasd")

}

func writeFromServer(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		c, err := conn.Read(data)
		if err != nil {
			log.Debug("rand", rand.Intn(10), "have no server write", err)
			return
		}
		log.Debug(string(data[0:c]) + "\n ")

	}
}
