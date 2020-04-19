/*
map与slice的使用
*/
package main

import "fmt"

func main() {
	// map类型的切片
	var mapSlice = make([]map[string]int, 8, 8) //只完成初始化一个map类型的切片
	mapSlice[0] = make(map[string]int, 8)       //完成了mapslice[1]的map初始化
	mapSlice[0]["xx"] = 100
	fmt.Println(mapSlice)
	// map的值为slice
	var sliceMap = make(map[string][]int, 8) //只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["中国"] = make([]int, 8)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
	}
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
