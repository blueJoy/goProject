package net

import "net"




func createConnection(){

	/*
		目前Dail支持如下网络协议：
		tcp,tcp4,tcp6,udp,udp4,udp6,ip,ip4,ip6

		建立连接后，使用conn的Write()发送数据，使用conn的Read()接收数据
	 */

	//创建TCP链接
	conn1,err := net.Dial("tcp","192.168.0.10:2100")

	//创建UDP链接
	conn2,err := net.Dial("udp","192.168.0.12:975")

	//创建ICMP链接(使用协议名称)
	conn3,err := net.Dial("ip4.icmp","www.baidu.com")

	//创建ICMP链接（使用协议编号）
	conn4,err := net.Dial("ip4:1","10.0.0.3")


}
