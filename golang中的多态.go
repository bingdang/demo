package main

import "fmt"

type people interface {
	process()
}

type leader struct {
	name string
	age  int
}

type staff struct {
	name string
	age  int
}

func (leader *leader) process() {
	fmt.Printf("我叫:%s,年龄:%d。我压榨员工\n", leader.name, leader.age)
}

func (staff *staff) process() {
	fmt.Printf("我叫:%s,年龄:%d。我被老板压榨\n", staff.name, staff.age)
}

func WhomakeMoney(p people) {
	p.process()
}

func main() {
	evan := &staff{name: "王大东", age: 18}
	hank := &staff{name: "陈韵小涵", age: 18}
	vincent := &leader{name: "颜辉", age: 38}
	wikifx := make([]people, 3)
	wikifx[0], wikifx[1], wikifx[2] = vincent, hank, evan
	for _, v := range wikifx {
		WhomakeMoney(v)
	}
}

/*
felix@MacBook-Pro project02 % go run main.go
我叫:颜辉,年龄:38。我压榨员工
我叫:陈韵小涵,年龄:18。我被老板压榨
我叫:王大东,年龄:18。我被老板压榨
*/
