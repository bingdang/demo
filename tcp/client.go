package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "192.168.31.253:88")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dial.Close()

	fmt.Println(dial.RemoteAddr(), "Connect")

	for {
		content := make([]byte, 1024)
		fmt.Println("请输入你要发送的内容:")
		fmt.Scanln(&content)
		dial.Write(content)

		n, err := dial.Read(content)
		if err != nil {
			fmt.Println(err)
			return
		}
		SContentBuf := content[:n]
		fmt.Println("接收到的数据:", string(SContentBuf))
	}

}
