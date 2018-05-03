package main

import "fmt"
import "sync"
import "runtime"
//共享内存的方式  实现并发编程   这种方式是golang不推荐的。
//Go语言提供的是另一种通信模型，即以消息机制而非共享内存作为通信方式。
var counter int = 0
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}