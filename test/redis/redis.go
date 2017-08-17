package main

import (
	"fmt"
	//"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	//切片
	//	var a = []int32{999999}

	//	_, err = c.Do("HSET", "username", "1", a)
	//	if err != nil {
	//		fmt.Println("redis set failed:", err)
	//	}
	//	username, err := c.Do("HGET", "username", "1")

	//	if err != nil {
	//		fmt.Println("redis set failed:", err)
	//	}

	//	fmt.Println(reflect.TypeOf(username)) //uint8

	//	username1, err := redis.String(c.Do("HGET", "username", "1"))
	//	if err != nil {
	//		fmt.Println("redis get failed:", err)
	//	} else {
	//		fmt.Printf("Got username %v \n", username1)
	//	}

	//map
	//	var b = make(map[string]int)
	//	b["aaa"] = 1
	//	b["bbbb"] = 2
	//	_, err = c.Do("HSET", "username", 2, b)
	//	if err != nil {
	//		fmt.Println("redis set failed:", err)
	//	}
	//	username2, err := redis.StringMap(c.Do("HGET", "username", 2))

	//	if err != nil {
	//		fmt.Println("redis get failed:", err)
	//	}

	//	fmt.Println(reflect.TypeOf(username2)) //map
	//	fmt.Println(username2)
	//	//fmt.Println(username2["aaa"])

	//	username3, err := redis.String(c.Do("HGET", "username", "1"))
	//	if err != nil {
	//		fmt.Println("redis get failed:", err)
	//	} else {
	//		fmt.Printf("Got username %v \n", username3)
	//	}

	//HMSET
	var b = make(map[string]int)
	b["aaa"] = 1
	b["bbbb"] = 2

	arg := make([]interface{}, 0)
	arg = append(arg, "username")
	for key, value := range b {
		arg = append(arg, key, value)
	}
	fmt.Println(arg)
	aa, err := c.Do("HMSET", arg...)
	fmt.Println(aa)
	fmt.Println(err)

}
