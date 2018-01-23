package main

import (
	"fmt"
	"time"
)

var counter int =0

func Count(ch chan int){
	ch <- 1   //向channel写数据
	fmt.Println("Counting")
}

func main(){
	chs := make([]chan int ,10)
	for i := 0;i<10 ;i++  {
		chs[i]=make(chan int)
		go Count(chs[i])
	}

	for _,ch := range chs{
		value := <-ch   //从channel读数据
		counter += value
	}

	fmt.Println("counter = ",counter)


	//死循环：随机向ch中写入0or1
/*	ch := make(chan  int, 1)
	for{
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <- ch
		fmt.Println("Value received:",i)
	}*/

	/** ----------处理timeout问题    go没有提供直接的超时处理机制，可以利用select  ------*/

	//1024为缓冲区大小
	ch := make(chan int,1024)
	timeout := make(chan bool,1)

	 go func() {
	 	time.Sleep(1e9)  //等待1秒钟
	 	timeout <- true
	 }()

	 //因为从timeout的channel中读取到了数据，从而达成1秒超时的效果
	 select{
	 case <-ch:
	 	//从ch中读取到数据
	 case <-timeout:
	 	//一直没有从ch中读取到数据，但从timeout中读取到了数据
	 }


	/**   ---------      单向channel           ------------*/
	//声明
	var ch1 chan int
	var ch2 chan <- float64   //单向写float64数据的channel
	var ch3 <-chan int       //单向读int数据的channel

	//初始化
	ch4 := make(chan int)
	ch5 := <-chan int(ch4)   //转换为单向读
	ch6 := chan <- int(ch4)  //转为单向写

	/*** -------- channel 关闭  ------*/
	close(ch1)
	x,ok := <-ch1
	if ok{
		//判断channel是否已经关闭
	}
}

/** --------------   Pipe特性     ---------------- */

//利用channel的传递实现pipe(管道)特性
type PipeData struct {
	value int
	handler func(int) int
	next chan int
}

//只要定义一些了PipeData的数据结构并一起传递个这个函数，就可以达到流式处理数据的目的
func handle(queue chan *PipeData){
	for data := range queue{
		data.next <- data.handler(data.value)
	}
}

//单向函数用法。。。单向读
func Parse(ch <-chan int){
	for value := range ch{
		fmt.Println("Parseing value",value)
	}
}
