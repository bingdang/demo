package main

import (
	"faceinter/app"
	"faceinter/data"
)

// 依赖定义的接口而不依赖实现
func callAppFunc(ap app.App) {
	ap.Sum(1, 2)
}

func main() {
	newsql := data.Newmgdb("127.0.0.1:27017")
	appins := app.NewAppIns(newsql) //app层依赖data层，通过data层私有构造函数做依赖注入

	//将实现传递进去
	callAppFunc(appins)
}
