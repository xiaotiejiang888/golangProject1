package main

import "fmt"

/**
https://studygolang.com/articles/2394
因为 default 特性， 我们可以使用 select 语句来检测 chan 是否已经满了
因为 ch 插入 1 的时候已经满了， 当 ch 要插入 2 的时候，发现 ch 已经满了（case1 阻塞住），
则 select 执行 default 语句。 这样就可以实现对 channel 是否已满的检测， 而不是一直等待。

比如我们有一个服务， 当请求进来的时候我们会生成一个 job 扔进 channel，
由其他协程从 channel 中获取 job 去执行。 但是我们希望当 channel 瞒了的时候，
将该 job 抛弃并回复 【服务繁忙，请稍微再试。】 就可以用 select 实现该需求。
 */
func main(){
	ch := make (chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
}
