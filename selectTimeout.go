package main

import (
	"time"
	"fmt"
)
/**
当超时时间到的时候，case2 会操作成功。 所以 select 语句则会退出。
而不是一直阻塞在 ch 的读取操作上。 从而实现了对 ch 读取操作的超时设置。
 */
func main(){
	timeout := make (chan bool, 1)
	go func() {
		time.Sleep(1e10) // sleep ten second
		timeout <- true
	}()
	ch := make (chan int)
	select {
	case <- ch:
	case <- timeout:
		fmt.Println("timeout!")
	}
}
