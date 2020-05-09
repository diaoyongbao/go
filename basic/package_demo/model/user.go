package model

import "fmt"

type User struct {
	Name string
	Age  int
}

// 构造函数,新建
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

// 修改User的值，需要修改他的指针对应的值
func (u *User) ModUser(name string, age int) {
	u.Name = name
	u.Age = age
}

// 获取user信息
func (u User) GetUser() {
	fmt.Println(u.Name)
	fmt.Println(u.Age)
}
