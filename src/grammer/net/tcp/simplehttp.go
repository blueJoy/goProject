package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

const(
	TCP="tcp"
	TCP4="tcp4"
)

func main(){

	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	/**     创建tcp连接      ***/
	//方案1
	conn,err := net.Dial(TCP,service)
	checkError(err)

	//方案二
	tcpAddr,err := net.ResolveTCPAddr(TCP4,service)
	conn1,err1 := net.DialTCP(TCP,nil,tcpAddr)


	_,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	
	result,err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)

	/*****    more func   *********/
	//校验IP的有效性
	net.ParseIP("ipaddress")
	//创建子网掩码   DefaultMask()获取默认子网掩码
	net.IPv4Mask(255,255,255,255)
	//根据域名查找IP
	net.ResolveTCPAddr("newowrk","address")
	net.LookupHost("name")


}
func readFully(conn net.Conn) ([]byte,error) {

	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for{
		n,err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil{
			if err == io.EOF{
				break
			}
			return nil,err
		}
	}

	return result.Bytes(),nil
}
func checkError(err error) {
	if err != nil{
		fmt.Fprint(os.Stderr,"Fail error: %s",err.Error())
		os.Exit(1)
	}
}

