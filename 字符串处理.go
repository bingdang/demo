package main

import (
	"fmt"
)

/*
字符串操作
可以通过Go标准库中的strings和strconv两个包中的函数进行相应的操作
--长度
func len(v Type) int
求字符串长度
func main() {
   fmt.Println(len("hello"))
}
--包含
func Contains(s, substr string) bool
字符串s中是否包含substr，返回bool值
func main() {
   fmt.Println(strings.Contains("hello", "llo"))
}
--开头结尾
func HasPrefix(s, prefix string) bool
判断字符串s是否以prefix为开头、HasSuffix是判断结尾
func main() {
   fmt.Println(strings.HasPrefix("hello","he"))
}
--连接
func Join(a []string, sep string) string
字符串链接，把slice a通过sep链接起来
func main() {
   s := []string{"abc", "456", "999"}
   fmt.Println(strings.Join(s, "** "))
}
--定位
func Index(s, sep string) int
在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
LastIndex是从后往前查找
func main() {
   fmt.Println(strings.Index("chicken", "ken"))
}
--复制
func Repeat(s string, count int) string
重复s字符串count次，最后返回重复的字符串
func main() {
   fmt.Println("ba" + strings.Repeat("na", 2))
}
--替换
func Replace(s, old, new string, n int) string
在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
func main() {
   fmt.Println(strings.Replace("ok ok ok", "k", "ky", 2))
}
--分割
func Split(s, sep string) []string
把s字符串按照sep分割，返回slice
func main() {
   fmt.Printf("%q\n", strings.Split("a,b,c", ","))
}
--掐头去尾
func Trim(s string, cutset string) string
在s字符串的头部和尾部去除cutset指定的字符串
func main() {
   fmt.Printf("[%q]", strings.Trim(" !哈!哈! ", "! "))
}
--去除空格分割
func Fields(s string) []string
去除s字符串的空格符，并且按照空格分割返回slice
func main() {
   fmt.Println( strings.Fields("  a b  c   "))
}
*/

/*
Append函数系列：将整数等转换为字符串后，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	//以10进制方式追加
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
    //4567false"abcdefg"'单'

Format系列函数：把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatInt(-1234, 10)
	//Uint无符号
	c := strconv.FormatUint(1234, 10)
	//与FormatInt一样，简写
	d := strconv.Itoa(-2234)
	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(a))
	//false -1234 1234 -2234
	//string
*/

func main() {
	//写一个程序，对英文字符串进行逆序
	str := "hello"
	str2 := make([]byte, 0, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		str2 = append(str2, str[i])
	}
	fmt.Println(string(str2))

	//写一个程序，对包含中文的字符串进行逆序
	strcn := "电饼铛A"
	str1cn := []rune(strcn)
	str2cn := make([]rune, 0, len(str1cn))
	for i := len(str1cn) - 1; i >= 0; i-- {
		str2cn = append(str2cn, str1cn[i])
	}
	fmt.Println(string(str2cn))

	//判断是否是回文"上海自来水来自海上"
	strRing := "上海自来水来自海上"
	strRingarray := []rune(strRing)
	strRingreverse := make([]rune, 0, len(strRingarray))
	for i := len(strRingarray) - 1; i >= 0; i-- {
		strRingreverse = append(strRingreverse, strRingarray[i])
	}
	if string(strRingarray) == string(strRingreverse) {
		fmt.Printf("%s，这句话是回文，切片长度为：%d", strRing, len(strRingreverse))
	} else {
		fmt.Printf("%s，这句话不是回文", strRing)
	}

}

/*
输出
olleh
A铛饼电
上海自来水来自海上，这句话是回文，切片长度为：9
*/
