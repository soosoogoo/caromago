package main

import (
	"fmt"
)

type B struct {
	userInfo interface{}
}

func (b *B) Set(info interface{}) {
	b.userInfo = info
}

func (b *B) GetInfo() interface{} {
	return b.userInfo
}

func main() {

	type C struct {
		c_1 string
	}

	type D struct {
		c_1 string
	}

	var c C
	c.c_1 = "aaa"

	//	var c = make(map[string]string)
	//	c["a"] = "s"

	var b = &B{userInfo: "a"}
	fmt.Println(b.userInfo)

	b.Set(c)

	var d = b.GetInfo()

	fmt.Println(d.(C).c_1)

}
