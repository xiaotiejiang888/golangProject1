package main

import "fmt"

/**
当 select 语句带有 default 的时候
因为 ch1 和 ch2 都为空，所以 case1 和 case2 都不会读取成功。 则 select 执行 default 语句。
 */
func main(){
	ch1 := make (chan int, 1)
	ch2 := make (chan int, 1)

	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	case <-ch2:
		fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}
}