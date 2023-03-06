package main

import (
	"fmt"
	"io"
	"os"
)

/*
func Create(name string) (file *File, err Error)
根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666

func NewFile(fd uintptr, name string) *File
根据文件描述符创建相应的文件，返回一个文件对象

func Open(name string) (file *File, err Error)
只读方式打开一个名称为name的文件

func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

func (file *File) Write(b []byte) (n int, err Error)
写入byte类型的信息到文件

func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
在指定位置开始写入byte类型的信息

func (file *File) WriteString(s string) (ret int, err Error)
写入string信息到文件

func (file *File) Read(b []byte) (n int, err Error)
读取数据到b中

func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
从off开始读取数据到b中

func Remove(name string) Error
删除文件名为name的文件
*/

func main() {
	//创建文件:
	myfile, err := os.Create("./feichi.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer myfile.Close()

	//写文件:
	for i := 0; i <= 5; i++ {
		myfile.WriteString("fei\n")
		myfile.Write([]byte("qq\n"))
	}

	//打开文件获取打开的文件句柄
	myfile2, err := os.Open("feichi.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer myfile2.Close()

	//读取文件:文件读取可以用file.Read()和file.ReadAt()，读到文件末尾会返回io.EOF的错误
	//创建叫buf的[]byte slice用于接受文件读取的数据，并创建一个chuck []byte充当缓冲区
	buf := make([]byte, 1024)
	var chuck []byte
	for {
		//读取到的字节数, 报错
		n, err := myfile2.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		//读取到的字节数为0，标识读取结束
		if n == 0 {
			break
		}
		//将读取出来的数据追加到缓冲区
		chuck = append(chuck, buf[:n]...)
	}
	fmt.Println(chuck, string(chuck))

	//拷贝文件 先读取，再写入
	feichi2, err := os.Create("./feichi2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer feichi2.Close()
	feichi2.Write(chuck)
}

/*
84行输出:
[102 101 105 10 113 113 10 102 101 105 10 113 113 10 102 101 105 10 113 113 10 102 101 105 10 113 113 10 102 101 105 10 113 113 10 102 101 105 10 113 113 10] fei
qq
fei
qq
fei
qq
fei
qq
fei
qq
fei
qq

93行copy文件:
felix@MacBook-Pro project02 % ls
feichi.txt      feichi2.txt     go.mod          main.go
*/
