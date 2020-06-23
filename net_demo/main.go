package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

// 代理池定义
type Proxy []struct{
	Proxy string `json:"proxy"`
	FailCount int `json:"fail_count"`
	Region  string `json:"region"`
	Type  string `json:"type "`
	Source  string `json:"source"`
	CheckCount int `json:"check_count"`
	LastStatus int `json:"last_status"`
	LastTime string `json:"last_time"`
}


func main(){
	reponse,err := http.Get("http://172.18.63.111:5010/get_all/")
	if err !=nil {
		fmt.Println("get faild")
	}
	fmt.Println("data")
	b ,err := ioutil.ReadAll(reponse.Body)
	if err !=nil {
		panic(err.Error())
	}
	// fmt.Println(string(b))
	// for i :=0 ;i< len(b); i++ {
	// 	fmt.Println(i)
	// }

	var  p Proxy
	err1 := json.Unmarshal([]byte(b), &p)
	if err1 != nil {
		fmt.Println("ERROR", err1)
		return
	}
	// fmt.Println(p)
	fmt.Println(len(p))
	for  i :=1;i< len(p); i++ {
		fmt.Println(p[i].Proxy)
	}
}