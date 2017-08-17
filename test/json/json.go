package main

import (
	"encoding/json"
	"fmt"
	//"os"
	//"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()
	var a = make(map[int]int)
	a[1] = 2
	a[2] = 3

	b := MapToString(a)
	c.Do("HSET", "testMapToString", "1", b)

	e, err := redis.Bytes(c.Do("HGET", "testMapToString", "1"))
	d := StringToMap(e)
	fmt.Println(d[1])

}
func StringToMap(a []byte) map[int]int {
	var d map[int]int
	json.Unmarshal(a, &d)
	return d
}

func MapToString(a map[int]int) []byte {
	b, _ := json.Marshal(a)
	return b
}
