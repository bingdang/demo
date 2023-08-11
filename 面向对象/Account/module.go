package Account

import "fmt"

type account struct {
	id      int
	pwd     string
	balance float64
}

// 构造函数
func CreateSession(c_id int, c_pwd string, c_balance float64) *account {
	return &account{
		id:      c_id,
		pwd:     c_pwd,
		balance: c_balance,
	}
}

func (a account) Login(pwd string) (bool, string) {
	if a.pwd != pwd {
		return false, "密码不正确"
	}
	return true, "登陆成功"
}

// 增
func (a *account) Save_Money(pwd string, moneys float64) {
	loginStatus, mass := a.Login(pwd)
	fmt.Println(mass)
	if loginStatus == true {
		a.balance += moneys
		fmt.Printf("存款成功,当前余额%v块钱\n", a.balance)
	}
}

// 查
func (a *account) Select_Money(pwd string) {
	loginStatus, mass := a.Login(pwd)
	fmt.Println(mass)
	if loginStatus {
		fmt.Printf("当前余额%v块钱\n", a.balance)
	}
}

// 删
func (a *account) Delete_Money(pwd string, moneys float64) {
	loginStatus, mass := a.Login(pwd)
	fmt.Println(mass)
	if loginStatus {
		a.balance -= moneys
		fmt.Printf("取款成功,当前余额%v块钱\n", a.balance)
	}
}
