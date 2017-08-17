package base

import (
	"encoding/json"
)

func StringToMap(a []byte) map[int]int {
	var d map[int]int
	json.Unmarshal(a, &d)
	return d
}

func MapToString(a map[int]int) []byte {
	b, _ := json.Marshal(a)
	return b
}
