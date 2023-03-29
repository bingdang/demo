package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
)

/*
解决中文乱码
go get github.com/axgle/mahonia
*/
var ps []string

/*
ps = []string{"北京市11", "天津市12", "河北省13",
"山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
"黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
"福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
"湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
"陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
"香港特别行政区81", "澳门特别行政区82", "台湾省83",}
*/

/* UTF8转GBK
iconv -f UTF-8 -t GBK ./kaifang.txt > ./kaifang-gbk.txt
*/

func HandleError(err error, where string) {
	if err != nil {
		fmt.Println(where, err)
	}
}

// 字符集转换
// 需要处理的数据、目前编码、返回正常数据
func CharacterSet(srcStr string, Encoder string) (dstStr string) {
	// 创建编码处理器
	encoder := mahonia.NewEncoder(Encoder)
	// 转换为UTF8的字符串
	dstStr = encoder.ConvertString(srcStr)
	return
}

func ReadFile(filename string) {
	open, err := os.Open(filename)
	HandleError(err, "OpenFile")
	defer open.Close()

	//建缓冲区
	reader := bufio.NewReader(open)
	for {
		Line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("读取完成")
			break
		}
		HandleError(err, "ReadLine")
		fmt.Println(CharacterSet(string(Line), "gbk"))
	}
}

func main() {
	ReadFile("./kaifang-gbk.txt")
}
