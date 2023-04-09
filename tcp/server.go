package main

import (
	"fmt"
	"net"
	"strings"
)

func Process_connections(acceptconn net.Conn) {
	defer acceptconn.Close()
	ipAdd := acceptconn.RemoteAddr().String()
	fmt.Println(ipAdd + "连接成功")

	for {
		buf := make([]byte, 1024)
		n, err := acceptconn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		CContentBuf := buf[:n]
		fmt.Printf("客户端:%s,发送数据:%s\n", ipAdd, string(CContentBuf))
		if string(CContentBuf) == "exit" {
			fmt.Printf("客户端:%s,退出连接\n", ipAdd)
			continue
		}

		acceptconn.Write([]byte(strings.ToUpper(string(CContentBuf))))
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:88")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	fmt.Println(listen.Addr(), "START")

	for {
		//阻塞等待连接
		acceptconn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go Process_connections(acceptconn)
	}

}
