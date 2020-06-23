package main

import "fmt"

func main( ){
	var ch1 chan bool
	ch1 = make(chan bool)

	go func(){
		for i:=0;i<10;i++ {
			fmt.Println("子goroutine 中，i=",i)
		}
		ch1 <- true
		fmt.Println("结束")

	}()
	data := <- ch1
	fmt.Println("main data->",data)
	fmt.Println("main over")
}

/*
channel 传值方向
发送数据 chan <- data
接受数据  data <- chan
*/