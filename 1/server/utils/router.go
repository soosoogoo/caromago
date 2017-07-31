package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	"../controller"
)

/*
	in this part, we try to decouple the whole code by a route-controller structure;
	before this server running, all the controller would be written in the router by function init();
	when the client send a json, this server decode this json and decide which controller to process this message;

	我在Server的内部加入一层Router,通过Router对通过Socket发来的信息，通过我们设定的规则进行解析判断后，调用相关的Controller进行任务的分发处理。
	在这个过程中不仅Controller彼此独立，匹配规则和Controller之间也是相互独立的。

*/

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

type Controller interface {
}

var routers [][2]interface{}

var countryCapitalMap map[string]interface{}

func Route(pred interface{}, controller Controller) {
	Log("pred is:", pred)
	switch pred.(type) {
	case func(entry Msg) bool:
		{
			Log("1111")
			var arr [2]interface{}
			arr[0] = pred
			arr[1] = controller
			routers = append(routers, arr)
		}
	case map[string]interface{}:
		{
			Log("2222")
			defaultPred := func(entry Msg) bool {
				for keyPred, valPred := range pred.(map[string]interface{}) {
					val, ok := entry.Meta[keyPred]
					if !ok {
						return false
					}
					if val != valPred {
						return false
					}
				}
				return true
			}
			var arr [2]interface{}
			arr[0] = defaultPred
			arr[1] = controller
			routers = append(routers, arr)
			fmt.Println(routers)
		}
	default:
		Log("3333")
		Log("didn't find requested controller")
	}
}

func TaskDeliver(postdata []byte, conn net.Conn) {
	for _, v := range routers {
		pred := v[0]
		act := v[1]
		var entermsg Msg
		err := json.Unmarshal(postdata, &entermsg)
		if err != nil {
			Log("4444")
			Log(err)
		}
		if pred.(func(entermsg Msg) bool)(entermsg) {

			Log("555")
			Log(entermsg.Meta["meta"])

			Log("asdasdasdasd")
			reflect.ValueOf(act.(Controller)).MethodByName("Test").Call([]reflect.Value{})

			//reflect.ValueOf(countryCapitalMap[className]).MethodByName("Test").Call([]reflect.Value{})

			//conn.Write(result)
			return
		}
	}
}

//定义控制器函数Map类型，便于后续快捷使用

func init() {

	/* 创建集合 */
	countryCapitalMap = make(map[string]interface{})

	var className = "TestController"
	countryCapitalMap[className] = &scontroller.TestController{}

	//	var aa = new(a)
	//reflect.ValueOf(countryCapitalMap[className]).MethodByName("Test").Call([]reflect.Value{})

	routers = make([][2]interface{}, 0, 20)
	Route(func(entry Msg) bool {
		return true
	}, countryCapitalMap[className])
}
