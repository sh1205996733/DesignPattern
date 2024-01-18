package main

import "fmt"

// 代理模式

// 静态代理:被代理者和代理者都要实现同一个接口
type Teacher interface {
	Teach()
}

type TeacherDao struct {
}

func (t *TeacherDao) Teach() {
	fmt.Println(" 老师授课中  。。。。。")
}

type TeacherDaoProxy struct {
	Target Teacher
}

func (t *TeacherDaoProxy) Teach() {
	fmt.Println("开始代理  完成某些操作。。。。。 ") //方法
	t.Target.Teach()
	fmt.Println("提交。。。。。") //方法
}

//动态代理
//todo
