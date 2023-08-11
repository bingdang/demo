package main

import "Bank/Account"

func main() {
	MyAccount := Account.CreateSession(9623503, "root591740@", 100000.5)
	//存钱
	MyAccount.Save_Money("root591740@", 20)
	//查余额
	MyAccount.Select_Money("root591740@")
	//取钱
	MyAccount.Delete_Money("root591740@", 1000)
}

/*
登陆成功
存款成功,当前余额100020.5块钱
登陆成功
当前余额100020.5块钱
登陆成功
取款成功,当前余额99020.5块钱
*/
