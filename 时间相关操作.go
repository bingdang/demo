package main

import (
	"fmt"
	"time"
)

/*
golang生日:2006-01-02 15:04:05
const (
    //纳秒
	Nanosecond  Duration = 1
    //微秒
	Microsecond          = 1000 * Nanosecond
    //毫秒
	Millisecond          = 1000 * Microsecond
    //秒
	Second               = 1000 * Millisecond
    //分钟
	Minute               = 60 * Second
    //小时
	Hour                 = 60 * Minute
)

ns := now.UnixNano() // 获得当前单位为纳秒的时间戳
log.Println("时间戳（秒）：", ns/1e9)		// 输出：时间戳（秒） ： 1665807442
log.Println("时间戳（毫秒）：", ns/1e6)	// 输出：时间戳（毫秒）： 1665807442207
log.Println("时间戳（微秒）：", ns/1e3)	// 输出：时间戳（微秒）： 1665807442207974
log.Println("时间戳（纳秒）：", ns)		// 输出：时间戳（纳秒）： 1665807442207974500

*/

//获取当前时间纳秒/1e3 = 微秒
func getNowTimeMicro() int64 {
	return time.Now().UnixNano() / 1e3
}

func timediff(startT, endT int64) (runT int64) {
	return endT - startT
}

func test() {
	//Sleep中传的是时间类型 = 4毫秒
	time.Sleep(4 * time.Millisecond)
}

func main() {
	//写一个程序，获取当前时间，并格式化成2019-01-01 08:00:00格式
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	//写一个程序，统计一段代码的执行耗时，单位精确到微秒
	startT := getNowTimeMicro()
	test()
	endT := getNowTimeMicro()
	diffT := timediff(startT, endT)

	fmt.Println("运行test()函数花费了", diffT, "微秒")

}

/*
2023-03-06 01:26:06
运行test()函数花费了 4334 微秒
*/
