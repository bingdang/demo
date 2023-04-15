package main

import (
	"fmt"
	"net/http"
)

// 回调
func admin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, r.Method, r.URL.Path, r.Header.Values("user-agent"), "连接成功")
	w.Write([]byte("弄啥"))
}

func main() {
	//单独写回调函数，根据不同的路由走不同的方法
	http.HandleFunc("/admin", admin)

	//add : 监听的地址
	//handler : 回调函数
	http.ListenAndServe("0.0.0.0:80", nil)
}
