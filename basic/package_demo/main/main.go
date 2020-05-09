package main

import (
	"fmt"
	"go/basic/package_demo/model"
	"go/basic/package_demo/utils"
	// 包需要从src目录下开始引用
)

func main() {
	fmt.Println("测试包的使用")
	// model.Sayhello()
	utils.Say()
	// 引用model包中的内容。并使用构造函数初始化数据
	u := model.NewUser("hah", 12)
	u.GetUser()
	u.ModUser("wang", 22)
	u.GetUser()
}
