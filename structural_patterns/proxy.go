package main

import "fmt"

// 代理模式 它允许通过使用代理对象来控制对另一个对象的访问。代理对象通常会拦截对被代理对象的访问，并在其间插入额外的逻辑功能或控制

/**
适配器模式 VS 代理模式

适配器模式旨在解决接口不兼容的问题，而代理模式则旨在控制对其它对象的访问。
适配器模式旨在使两个不兼容的接口可以协同工作；而代理模式则旨在实现对目标对象的控制，在不改变目标对象的基础上增强其功能或保护其安全性。
*/
// 静态代理:被代理者和代理者都要实现同一个接口

// 1、抽象角色
type Teacher interface {
	Teach()
}

// 2、真实角色
type TeacherDao struct {
}

func (t *TeacherDao) Teach() {
	fmt.Println(" 老师授课中  。。。。。")
}

// 3、创建代理角色
type TeacherDaoProxy struct {
	Target Teacher // 需要将真实角色聚合到代理角色中
}

func (t *TeacherDaoProxy) Teach() {
	fmt.Println("开始代理  完成某些操作。。。。。 ") //方法
	t.Target.Teach()
	fmt.Println("提交。。。。。") //方法
}

//动态代理
//todo
