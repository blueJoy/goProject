package main

import (
	"errors"
	"net/rpc"
	"net"
	"log"
	"net/http"
	"fmt"
)

type Args struct {
	A,B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args,reply *int) error{
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args,quo *Quotient) error{
	if args.B == 0{
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main(){

	//开启rpc服务
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp","127.0.0.1:1234")
	if e != nil{
		log.Fatal("listen error:",e)
	}
	go http.Serve(l,nil)

	//建立连接
	client, err := rpc.Dial("tcp","127.0.0.1:1234")
	if err != nil{
		log.Fatal("dialing:",err)
	}

	//同步调用
	args  := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply",args,&reply)
	if err != nil{
		log.Fatal("arith error:",err)
	}
	fmt.Printf("Arith:%d*%d=%d",args.A,args.B,reply)

	//异步调用
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide",args,&quotient,nil)
	replyCall := <- divCall.Done
	fmt.Printf("over",replyCall)
}

