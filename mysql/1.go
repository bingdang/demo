package main

import (
	"fmt"
	"time"

	//仅仅只引用驱动
	_ "github.com/go-sql-driver/mysql"

	//mysql方法库
	"github.com/jmoiron/sqlx"
)

/*
CREATE TABLE `person` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `money` double DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3
*/

// 查询抽象对象
type Person struct {
	Id       int       `db:"id"`
	Name     string    `db:"name"`
	Age      int       `db:"age"`
	Money    float32   `db:"money"`
	Birthday time.Time `db:"birthday"`
}

func main() {
	//1.连接数据库
	//parseTime=true 解析处理数据库事件
	opendb, _ := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tt?parseTime=true")
	defer opendb.Close()

	//2.添加数据
	result, _ := opendb.Exec("insert into person(name,age,money,birthday) values(?,?,?,?)",
		"张三", 23, 100.5, 20230606)
	//得到受影响行数
	row, _ := result.RowsAffected()
	//受影响的最后一位id
	lastid, _ := result.LastInsertId()

	fmt.Println("受影响行数", row, "受影响的最后一位id", lastid)

	//3.修改数据
	result, _ = opendb.Exec("update person set name=? where name=?", "李四", "张三")
	fmt.Println("受影响行数", row, "受影响的最后一位id", lastid)

	//4.删除数据
	result, _ = opendb.Exec("delete from person where id=?", 3)
	fmt.Println("受影响行数", row, "受影响的最后一位id", lastid)

	//5.事务操作
	tx := opendb.MustBegin()
	tx.MustExec("insert into person(name,age,money,birthday) values(?,?,?,?)",
		"张菲", 20, 101.5, 20230606)
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}

	//6.查询数据
	//传入切片
	var ps []Person
	err = opendb.Select(&ps, "SELECT id,name,age,money,birthday from person where id=?", 6)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", ps)
	//[{6 张菲 20 101.5 2023-06-06 00:00:00 +0000 UTC}]

	//查询全部
	err = opendb.Select(&ps, "SELECT id,name,age,money,birthday from person")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", ps)
	//[{1 李四 23 100.5 2023-06-06 00:00:00 +0000 UTC} {2 李四 23 100.5 2023-06-06 00:00:00 +0000 UTC} {4 李四 23 100.5 2023-06-06 00:00000 UTC} {5 李四 23 100.5 2023-06-06 00:00:00 +0000 UTC} {6 张菲 20 101.5 2023-06-06 00:00:00 +0000 UTC} {7 李四 23 100.5 2023-00:00:00 +0000 UTC} {8 张菲 20 101.5 2023-06-06 00:00:00 +0000 UTC} {9 李四 23 100.5 2023-06-06 00:00:00 +0000 UTC} {10 张菲 20 1023-06-06 00:00:00 +0000 UTC} {11 李四 23 100.5 2023-06-06 00:00:00 +0000 UTC} {12 张菲 20 101.5 2023-06-06 00:00:00 +0000 UTC}]
}
