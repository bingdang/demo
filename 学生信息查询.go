package main

import (
	"flag"
	"fmt"
)

//	写一个程序，实现学生信息的存储，学生有id、年龄、分数等信息，需要非常方便地通过id查找到对应学生的信息
// id:1 name:zhang.mr age:18 score:98

var (
	students = map[int]map[string]interface{}{
		1: {"id": 1, "name": "zhang.mr", "age": 18, "score": 98},
		2: {"id": 2, "name": "yao.mr", "age": 17, "score": 100},
		3: {"id": 3, "name": "li.mr", "age": 17, "score": 100},
		4: {"id": 4, "name": "yu.mr", "age": 19, "score": 100},
	}
	id int
)

func Init() {
	flag.IntVar(&id, "n", -1, "学生id")
	flag.Parse()
}

func studentsQuery(id int) {
	m, ok := students[id]
	if !ok {
		fmt.Println("此学生不存在")
	} else {
		fmt.Printf("学生信息：%s", m)
	}
}

func main() {
	Init()
	studentsQuery(id)
}
