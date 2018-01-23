package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method,params string) *Response {
	return &Response{"200","hello EchoServer:"+method +"--"+params}
}

func (server *EchoServer) Name() string{
	return "EchoServer"
}

func TestIpc(t *testing.T){

	server := NewIpcServer(&EchoServer{})

	client1 := newIpcClient(server)
	client2 := newIpcClient(server)

	resp1,err1 := client1.Call("test","From Client1")
	resp2,err2 := client2.Call("test","From Client2")
	if err1 != nil || resp1.Body != "hello EchoServer:test--From Client1"{
		t.Error("IpcClient.Call failed. resp1:")
	}
	if err2 != nil || resp2.Body != "hello EchoServer:test--From Client2"{
		t.Error("IpcClient.Call failed. resp2:")
	}

	client1.Close()
	client2.Close()
}