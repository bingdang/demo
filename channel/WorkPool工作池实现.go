package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 2.定义两个对象分别是任务输入管道和右边任务输出管道
type Job struct {
	id           int
	randomNumber int
}

type Result struct {
	// 这里必须传对象实例即指针否则没值
	Job *Job
	sum int
}

// 2.创建工作池
func CountWorkPool(Concurrency int, job chan *Job, result chan *Result) {
	for i := 0; i < Concurrency; i++ {
		go func(Job chan *Job, result chan *Result) {
			//读取数据
			for job := range job {
				sNum := job.randomNumber

				var sum int
				for sNum != 0 {
					tmp := sNum % 10 //取模运算取最后一位
					sum += tmp
					sNum /= 10 //去掉最后一位
				}

				//将计算结果放入结果管道
				result <- &Result{
					job,
					sum,
				}
			}
		}(job, result)
	}
}

/*
	WorkPool

计算随机数所有位数之和
job --> WorkPool --> result
协程实现WorkPool
*/
func main() {
	//1.定义两个管道
	job := make(chan *Job, 10240)
	result := make(chan *Result, 10240)

	//2.创建工作池
	CountWorkPool(2560, job, result)

	//3.打印结果
	go func(Concurrency int, result chan *Result) {
		for resultPrice := range result {
			fmt.Printf("job id:%d randnum:%d result:%d\n", resultPrice.Job.id, resultPrice.Job.randomNumber,
				resultPrice.sum)
		}
	}(2560, result)

	//4.生成随机数并放入管道
	for id := 0; id <= 100000; id++ {
		r_num := rand.Int()
		job <- &Job{
			id:           id,
			randomNumber: r_num,
		}
	}

	//防止main退出 协程没跑完。这里等待10秒
	<-time.NewTimer(10 * time.Second).C

}

/*
job id:248357 randnum:7432583274265739194 result:91
job id:248358 randnum:495238005400424024 result:56
job id:248359 randnum:6562667944349411816 result:92
....
*/
