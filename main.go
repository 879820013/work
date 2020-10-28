package main

import (
	"fmt"
	"os"
)
var smr studentMgr
func showMenu()  {
	fmt.Println("hi")
	fmt.Println(`
	1.查看
	2.添加
	3.修改
	4.删除
	5.退出
	`)
}

func main()  {
	var smr  = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for  {
		showMenu()
		fmt.Print("输入：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("输入了：",choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("没有这个选项：",choice)
		}
	}
}
