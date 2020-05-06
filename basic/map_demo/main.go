// go语言中的map是引用类型，需要初始化使用，未初始化时为nil
// map[keyType]valueType
// make(map[keyType]valueType,cap)
package main

import (
	"fmt"
	"sort"
)

func main() {
	// map的两种构造方式
	scoreMap := make(map[string]int, 8)
	scoreMap["小米"] = 88
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小米"])

	a := map[string]int{
		"xc": 11,
		"xx": 66,
	}
	fmt.Printf("%#v\n", a)
	// 判断某个键存不存在
	// 如果存在，v是默认类型的0值，如scoreMap v的类型为int，默认0，
	// ok，如果存在 返回 true，如果不存在，返回false
	v, ok := scoreMap["小米"]
	fmt.Println(v, ok)
	if ok {
		fmt.Println("小米 is in socreMap", v)
	} else {
		fmt.Println("小米 is not int scoreMap")
	}
	// 遍历map
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
	// 只遍历map中的key
	for key := range scoreMap {
		fmt.Println(key)
	}
	// 只遍历map中的value
	for _, value := range scoreMap {
		fmt.Println(value)
	}
	// 按照某个顺序遍历map
	// 按照key的大小遍历scoremap
	keys := make([]string, 0, 10)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	// 删除某个键值对
	delete(scoreMap, "小米")
	fmt.Println(scoreMap)
}
