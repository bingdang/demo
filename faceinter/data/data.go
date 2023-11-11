package data

import "fmt"

type Sql interface {
	Open() error
}

type mysql struct {
	url string
}

func Newmysql(url string) *mysql {
	return &mysql{url: url}
}

func (sql *mysql) Open() error {
	fmt.Println("打开了", sql.url)
	return nil
}

type mgdb struct {
	url string
}

func Newmgdb(url string) *mgdb {
	return &mgdb{url: url}
}

func (sql *mgdb) Open() error {
	fmt.Println("打开了", sql.url)
	return nil
}
