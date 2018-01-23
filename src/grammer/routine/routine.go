package main

import (
	"fmt"
	"sync"
	"runtime"
)

func Add(x,y int){
	z := x+y
	fmt.Println("z = ",z)
}

var counter int = 0

func Count(lock *sync.Mutex){
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}


//并发通信模型：共享数据和消息
func main(){
	//普通引用,因为主函数启动了10个goroutine函数就结束了。。所有goroutine函数不一定执行完成
/*	for i := 0;i < 10 ; i++  {
		go Add(i,i)
	}*/


	/*
		共享数据   10个线程共享counter,所以每次操作都必须加锁
		协程：轻量级线程（也叫用户线程），无需系统调度

		ps:不要通过共享内存来通信，而应该通过通信来共享内存
	 */
	lock := &sync.Mutex{}

	for i := 0; i < 10 ; i++  {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c>= 10{
			break
		}
	}


}