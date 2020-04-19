// 切片的学习
// 切片是一个拥有相同类型元素的可变长度的序列，它是基于数据类型做的一层封装
// 类型声明  var name []T  name:变量名  T 类型声明
package main

import "fmt"

func main() {
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个bool切片，并初始化

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 基于数组获取切片
	l := [5]int{1, 12, 13, 51, 44}
	l2 := l[1:4]
	fmt.Println(l2) //[12 13 51]
	fmt.Printf("%T\n", l2)
	// 切片后再次切片
	l3 := l2[0:2]
	fmt.Println(l3) //[12 13]
	// make函数构造切片,make([]T,size,cap),T 切片元素类型，size 切片元素数量，cap 切片容量
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Println(len(d), cap(d))
	fmt.Printf("%T\n", d)
	// 切片的判断
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}

	// 切片的复制拷贝
	s1 := make([]int, 3)
	s2 := s1
	s2[0] = 100
	fmt.Println(s1, s2) //[100 0 0] 验证了切片复制公用一个底层数组

	// 切片的遍历
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	for index, value := range l {
		fmt.Println(index, value)

	}
	// 切片的扩容
	var x []int //未申请内存地址
	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", x, len(x), cap(x), x)
	}
	// 切片拷贝
	area := []string{"北京", "上海", "广州", "深圳"}
	// fmt.Printf("%d", len(area))
	areaCopy := make([]string, 4)
	copy(areaCopy, area)
	fmt.Println(areaCopy)
	// 切片删除元素
	area = append(area[0:2], area[3:]...)
	fmt.Println(area)
}
