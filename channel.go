package main

import "fmt"

//channel 是golang在语言级别提供的goroutine间的通信方式
//我们可以使用channel在两个或多个goroutine之间传递消息
//如果要跨进程通信，建议采用分布式系统的方法来解决  比如socket或者http等通信协议
//一个channel只能传递一种类型的值  这个类型需要在声明channel时指定

/**
定义了一个包含10个channel的数组（名为chs），并把数组中的每个
channel分配给10个不同的goroutine。在每个goroutine的Add()函数完成后，我们通过ch <- 1语
句向对应的channel中写入一个数据。在这个channel被读取前，这个操作是阻塞的。在所有的
goroutine启动完成后，我们通过<-ch语句从10个channel中依次读取数据。在对应的channel写入
数据前，这个操作也是阻塞的。这样，我们就用channel实现了类似锁的功能，进而保证了所有
goroutine完成后主函数才返回
 */
func Count2(ch chan int){
	fmt.Println("Counting")
	//向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据
	ch <- 1 //在这个channel被读取前，这个操作是阻塞的
}

func main(){
	chs := make([]chan int, 10)
	for i:=0;i<10;i++  {
		chs[i]= make(chan int)
		go Count2(chs[i])
	}
	//如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。
	for _,ch:=range(chs){
		<-ch //在对应的channel写入数据前，这个操作也是阻塞的
	}
}