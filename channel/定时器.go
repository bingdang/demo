package main

import (
	"fmt"
	"time"
)

func timerF() {
	//1.timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()

	fmt.Printf("t1:%v\n", t1)
	fmt.Printf("timer1:%v\n\n", <-timer1.C)
	/*
		t1:2023-03-22 23:41:50.593029 +0800 CST m=+0.000107477
		timer1:2023-03-22 23:41:52.594116 +0800 CST m=+2.001213259%
	*/

	//2.验证timer只能执行一次
	/*
		timer2 := time.NewTimer(2 * time.Second)
		for {
			<-timer2.C
			fmt.Println("取出一次")
		}
	*/

	/*
		取出一次
		fatal error: all goroutines are asleep - deadlock! //管道中只有一个时间值，取了一次就不能再取了
	*/

	//3.使用timer实现延迟功能
	//方法1:
	time.Sleep(2 * time.Second)
	fmt.Println("1 延迟时间到了", time.Now())
	//方法2:
	timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	fmt.Println("2 延迟时间到了", <-timer3.C)
	//方法3:
	//<-time.After(2 * time.Second)
	fmt.Println("3 延迟时间到了", <-time.After(2*time.Second))

	/*
		1 延迟时间到了 2023-03-22 23:52:57.027143 +0800 CST m=+4.001684033
		2 延迟时间到了 2023-03-22 23:52:59.028323 +0800 CST m=+6.002925748
		3 延迟时间到了 2023-03-22 23:53:01.029397 +0800 CST m=+8.004061694
	*/

	//4.停止定时器
	time4 := time.NewTimer(2 * time.Second)
	boolt := time4.Stop()
	if boolt {
		fmt.Println("定时器被停止了")
	}
	/*
		定时器被停止了
	*/

	//5.定时器重置
	time5 := time.NewTimer(2 * time.Second)
	time5.Reset(4 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-time5.C)
	/*
		2023-03-22 23:59:31.868701 +0800 CST m=+8.005341040
		2023-03-22 23:59:35.86889 +0800 CST m=+12.005552423
	*/
}

func tickerF() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		i++
		<-ticker.C              //取1秒
		fmt.Println(<-ticker.C) //再取1秒
		if i == 5 {
			ticker.Stop()
		}

	}
}

/*
2023-03-23 00:13:56.945535 +0800 CST m=+2.001219982
2023-03-23 00:13:58.945562 +0800 CST m=+4.001258048
2023-03-23 00:14:00.944661 +0800 CST m=+6.000367660
2023-03-23 00:14:02.945501 +0800 CST m=+8.001218576
2023-03-23 00:14:04.945487 +0800 CST m=+10.001216062
fatal error: all goroutines are asleep - deadlock!
*/

func main() {
	//只运行一次的定时器。只网管道延迟放入一个时间数据。只能取一次
	//timerF()
	tickerF()

}
