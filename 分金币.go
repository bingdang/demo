package main

import (
	"fmt"
)

// 	有100枚金币，需要分配给以下用户，分配规则如下：
// 	名字中包含a或A，分配1枚金币
// 	名字中包含e或E，分配1枚金币
// 	名字中包含i或I，分配2枚金币
// 	名字中包含o或O，分配3枚金币
// 	名字中包含u或U，分配5枚金币
// Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth

var (
	Users    = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	UsersMap = make(map[string]int, 10)
)

// 封装
func addCoin(user string, coins int) {
	_, ok := UsersMap[user]
	if !ok {
		UsersMap[user] = coins
	} else {
		UsersMap[user] += coins
	}
}

func divideCoin(user string) int {
	num := 0
	for _, v := range user {
		switch v {
		case 'a', 'A':
			num += 1
		case 'e', 'E':
			num += 1
		case 'i', 'I':
			num += 2
		case 'o', 'O':
			num += 3
		case 'u', 'U':
			num += 5
		}
	}
	return num
}

func lineup(users []string) int {
	num := 100
	for _, v := range users {
		x := divideCoin(v)
		addCoin(v, x)
		num -= x
	}
	return num
}

func main() {
	i := lineup(Users)
	for u, v := range UsersMap {
		fmt.Printf("用户: %s,分到了金币: %d\n", u, v)
	}
	fmt.Println("还剩:", i)
}
