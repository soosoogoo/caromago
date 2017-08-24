package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"

	"github.com/name5566/leaf/log"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.115:3563")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go writeFromServer(conn)
	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
		"Login": {
			"RoleId": 1
		}
	}`)

	// 发送消息
	msg := truncatMsg(data)
	conn.Write(msg)

	//创建房间
	data2 := []byte(`{
			"CreateRoom": {
				"Stage_id": 1,
				"Type": 1,
				"Ut_id": 1,
				"Is_open": 1
			}
		}`)

	// 发送消息
	msg2 := truncatMsg(data2)
	conn.Write(msg2)

	for {
		var talkContent string
		fmt.Scanln(&talkContent)
		data := []byte(`{
			"RoomList": {
				"RoomList": {
					"Stage_id": 1,
					"Room_id": 1,
					"Ul_id": 1,
					"Type": 1,
					"StageType": 1
				}
			}
		}`)

		if len(talkContent) > 0 {
			_, err = conn.Write(truncatMsg(data))
			if err != nil {
				fmt.Println("write to server error")
				return
			}
		}
	}

}

func writeFromServer(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		c, err := conn.Read(data)
		if err != nil {
			log.Debug("rand", rand.Intn(10), "have no server write", err)
			return
		}

		log.Debug(string(data[2:c]) + "\n ")

	}
}

func truncatMsg(data []byte) []byte {
	// len + data
	//data := []byte(`{"RoleLogin": {"RoleId": ` + talkContent + `}}`)
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	return m
}
