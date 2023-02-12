package main

import (
	"fmt"
	"strings"
)

func Wordscount(Words string) (map1 map[string]int) {
	words := strings.Split(Words, " ") //将单词切割成切片
	map1 = make(map[string]int, 16)
	for _, v := range words { //遍历切片，并检查map中是否存在该键
		count, ok := map1[v]
		if !ok {
			map1[v] = 1 //不存在则赋值
		} else {
			map1[v] = count + 1 //存在并赋值
		}
	}
	return map1
}

//单词统计
func main() {
	Words := "to be or not to be"
	chr := Wordscount(Words)
	fmt.Println(chr)
	//map[be:2 not:1 or:1 to:2]
}
