package main

import (
	"fmt"
	"os"
	"project02/Students"
)

var (
	AllStudents []*Students.Students
)

//	实现一个简单的学生管理系统，每个学生有分数、年级、性别、名字等字段，用户可以在控制台输入1添加学生、输入2修改学生信息、输入3打印所有学生列表，输入4结束程序，如下：

func OutputStudent() {
	welcome := `
1.添加学生
2.修改学生
3.打印所有学生
4.结束程序
`
	fmt.Printf("%s请输入你要做的操作编号:", welcome)
}

func InputStudent() *Students.Students {
	var (
		name  string
		score float32
		sex   int
		grade int
	)
	fmt.Println("请输入学生姓名[中文]")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学生分数[小数｜整数]")
	fmt.Scanf("%f\n", &score)
	fmt.Println("请输入学生性别[1|0]")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("请输入学生年级[整数]")
	fmt.Scanf("%d\n", &grade)

	Stu := Students.NewStudents(name, score, sex, grade)
	return Stu

}

func AddStudents() {
	Si := InputStudent()
	for index, v := range AllStudents {
		if v.Name == Si.Name {
			//同名则更新
			AllStudents[index] = Si
			fmt.Println("更新成功")
			return
		}
	}

	AllStudents = append(AllStudents, Si)
	fmt.Println("添加成功")
}

func ModifyStudents() {
	Si := InputStudent()
	for index, v := range AllStudents {
		if v.Name == Si.Name {
			//同名则更新
			AllStudents[index] = Si
			fmt.Println("更新成功")
			return
		}
	}
	fmt.Println("\n查无此人!!!!")
}

func PrintStudents() {
	for index, v := range AllStudents {
		fmt.Printf("%s,%#v", AllStudents[index].Name, v)
	}
}

func main() {
	for {
		OutputStudent()
		var i int
		fmt.Scanf("%d\n", &i)
		switch i {
		case 1:
			AddStudents()
		case 2:
			ModifyStudents()
		case 3:
			PrintStudents()
		case 4:
			os.Exit(0)
		}

	}

}
