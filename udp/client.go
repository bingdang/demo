package main

import (
	"fmt"
	"net"
)

func main() {
	//1.连接服务端
	dialUDP, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 80,
	})
	if err != nil {
		return
	}
	defer dialUDP.Close()

	//1.写数据到服务端
	_, err = dialUDP.Write([]byte("驰！"))
	if err != nil {
		fmt.Println(err)
	}

	//3.接收数据
	buf := make([]byte, 16)
	n, addr, err := dialUDP.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("data:%s,addr:%v", string(buf[:n]), addr)

}
