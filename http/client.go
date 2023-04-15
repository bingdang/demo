package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1/admin")
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	buf := make([]byte, 100)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			fmt.Println(string(buf[:n]))
			break
		}
	}
}
