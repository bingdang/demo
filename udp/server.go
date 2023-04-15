package main

import (
	"fmt"
	"net"
)

func main() {
	listenUDP, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 80,
	})
	if err != nil {
		return
	}

	fmt.Println(listenUDP.LocalAddr(), "UDPSTART")

	defer listenUDP.Close()

	for {
		// 缓冲区
		var buf [1024]byte

		//接收udp的传输
		n, addr, err := listenUDP.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println(err)
			continue
			//报错跳出循环
		}

		//打印读到的数据
		fmt.Printf("data:%s, addr:%v\n", string(buf[0:n]), addr)

		//返回信息
		_, err = listenUDP.WriteToUDP([]byte("666"), addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
