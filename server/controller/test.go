package scontroller

import (
	"log"
)

type testController struct {
	i    int
	name string
}

func (ac *testController) Test() {
	log.Println("asd")
}
