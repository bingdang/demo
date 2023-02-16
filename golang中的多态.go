package main

import "fmt"

type people interface {
	process()
}

type party interface {
	//相当于写了process()，接口的继承
	people
	sing(lyric string)
}

type leader struct {
	name string
	age  int
}

type staff struct {
	name string
	age  int
}

func (staff *staff) sing(lyric string) {
	fmt.Printf("我叫:%s,年龄:%d。我聚会只能唱:%s\n", staff.name, staff.age, lyric)
}

func (leader *leader) sing(lyric string) {
	fmt.Printf("我叫:%s,年龄:%d。我聚会想唱:%s\n", leader.name, leader.age, lyric)
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
	evan.sing("互撸娃")
	vincent.sing("敢问路在何方")

}

/*
felix@MacBook-Pro project02 % go run main.go
我叫:颜辉,年龄:38。我压榨员工
我叫:陈韵小涵,年龄:18。我被老板压榨
我叫:王大东,年龄:18。我被老板压榨
我叫:王大东,年龄:18。我只能唱:互撸娃
我叫:颜辉,年龄:38。我想唱:敢问路在何方
*/
