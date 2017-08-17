package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()

	//	c.Send("HMSET", "album:1", "title", "Red", "rating", 5)
	//	c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1)
	//	c.Send("HMSET", "album:3", "title", "Beat")
	//	c.Send("LPUSH", "albums", "1")
	//	c.Send("LPUSH", "albums", "2")
	//	c.Send("LPUSH", "albums", "3")
	//	values, err := redis.Values(c.Do("SORT", "albums",
	//		"BY", "album:*->rating",
	//		"GET", "album:*->title",
	//		"GET", "album:*->rating"))
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	//	for len(values) > 0 {
	//		var title string
	//		rating := -1 // initialize to illegal value to detect nil.
	//		values, err = redis.Scan(values, &title, &rating)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		if rating == -1 {
	//			fmt.Println(title, "not-rated")
	//		} else {
	//			fmt.Println(title, rating)
	//		}
	//	}

	//	//切片
	arr := []int{1, 2}
	lang, err := json.Marshal(arr)
	fmt.Println(reflect.TypeOf(lang))

	var wo []int
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []int==")
		fmt.Println(wo)

		for key, val := range wo {
			fmt.Println(key, '-', val)
		}
	}

	//	_, err = c.Do("HSET", "username", "1", a)
	//	if err != nil {
	//		fmt.Println("redis set failed:", err)
	//	}
	//	username, err := c.Do("HGET", "username", "1")

	//	if err != nil {
	//		fmt.Println("redis set failed:", err)
	//	}

	//	fmt.Println(reflect.TypeOf(username)) //uint8
	//	fmt.Println(username)
	//	append(username, 1)
	//	fmt.Println(username)
	//	for val := range username {
	//		fmt.Println(val)
	//	}

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
	//	var b = make(map[string]int)
	//	b["aaa"] = 1
	//	b["bbbb"] = 2

	//	arg := make([]interface{}, 0)
	//	arg = append(arg, "username")
	//	for key, value := range b {
	//		arg = append(arg, key, value)
	//	}
	//	fmt.Println(arg)
	//	aa, err := c.Do("HMSET", arg...)
	//	fmt.Println(aa)
	//	fmt.Println(err)

}
