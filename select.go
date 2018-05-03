package main

import "fmt"

/**
由select开始一个新的选择块，每个选择条件由
case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的
限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，大致的结构如下：
select {
 case <-chan1:
 // 如果chan1成功读到数据，则进行该case处理语句
 case chan2 <- 1:
 // 如果成功向chan2写入数据，则进行该case处理语句
default:
 // 如果上面都没有成功，则进入default处理流程
}
 */
func main()  {
	ch := make(chan int, 1)
	//select的每个case语句里必须是一个IO操作
	//for {
	//随机向ch中写入一个0或者1的过程
	//select 多个 case 同时满足的话，会随机触发一个
		select {
			case ch<-0:
			case ch<-1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	//}
}
