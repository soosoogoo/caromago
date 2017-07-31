package scontroller

import (
	"log"
)

func Test() {
	log.Println("asdasd")
}

type TestController struct {
}

func (ac *TestController) Test() string {

	log.Println("2222222222222222222")
	return "2222222222"
}

func (this *TestController) Excute(msg []byte) []byte {
	log.Println("echo the message:11111111111111111111")

	return msg
}
