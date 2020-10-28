package main

import "fmt"

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

//查看
func (s studentMgr) showStudents() {
	for _, v := range s.allStudent {
		fmt.Printf("学号：%d,名字%s\n", v.id, v.name)
	}
}

//增加
func (s studentMgr) addStudents() {
	var id int64
	var name string
	fmt.Print("输入学号：")
	fmt.Scanln(&id)
	fmt.Print("输入姓名：")
	fmt.Scanln(&name)

	s.allStudent[id] = student{
		id:id,
		name:name,
	}
}

//修改
func (s studentMgr) editStudents() {
	var id int64
	var name string
	fmt.Print("输入学号：")
	fmt.Scanln(&id)
	fmt.Print("输入姓名：")
	fmt.Scanln(&name)

	s.allStudent[id] = student{
		id:id,
		name:name,
	}
}

//删除
func (s studentMgr) deleteStudents() {
	var id int64
	fmt.Print("输入学号：")
	fmt.Scanln(&id)
	delete(s.allStudent,id)
}
