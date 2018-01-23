package ipc

import "encoding/json"

type IpcClient struct{
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient{
	c := server.Connect()
	return &IpcClient{c}
}

/*
 会将调用信息封装成json格式发送到对于的channel，并等待获取反馈
 */
func (client *IpcClient) Call(method,params string) (resp *Response,err error)  {

	req := &Request{method,params}

	var b []byte
	//把对象转换为json
	b,err = json.Marshal(req)

	if err != nil{
		return
	}

	//发送到channel
	client.conn <- string(b)

	str := <- client.conn   //等待返回值

	var resp1 Response
	//json转换为对象
	err = json.Unmarshal([]byte(str),&resp1)
	resp = &resp1

	return
}

func (client *IpcClient) Close()  {
	client.conn <- "CLOSE"
}