package count

/*
	go的标准包testing提供了自动化测试相关的框架
	支持单元测试和压力测试
	测试的代码文件必须以 _test.go 结尾
	单元测试的函数名必须以Test开头，并且只有一个参数，类型是 *testing.T
	压力测试的函数名必须以Benchmark开头，并且只有一个参数，类型是 *testing.B
*/

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 2
	b := 1
	t.Logf("a=%d,b=%d", a, b)
	add := Add(a, b)
	fmt.Println(add)
}

/*
=== RUN   TestAdd
    count_test.go:11: a=2,b=1
3
--- PASS: TestAdd (0.00s)
PASS

进程 已完成，退出代码为 0
*/

func BenchmarkSub(b *testing.B) {
  //N由测试框架自己决定
	for i := 0; i <= b.N; i++ {
		a := 1000
		b := 890
		Sub(a, b)
	}
}

/*
goos: darwin
goarch: amd64
pkg: 0402
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSub
                    执行次数               平均耗时
BenchmarkSub-12    	1000000000	         0.3318 ns/op
PASS

进程 已完成，退出代码为 0
*/
