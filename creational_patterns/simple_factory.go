package main

// 简单工厂模式 可扩展性不是很好
// 简单工厂模式:定义了一个创建对象的类，由这个类来封装实例化对象的行为(代码)
// 在软件开发中，当我们会用到大量的创建某种、某类或者某批对象时，就会使用到工厂模式.
func SimpleFactory(objType string) string { //返回一个公共接口
	var o string
	if objType == "greek" {
		o = "greek"
	} else if objType == "cheese" {
		o = "cheese"
	} else if objType == "pepper" {
		o = "pepper"
	}
	return o
}

/**
• 对于拟构造的组件，需要依据其共性，抽离出一个公共 interface
• 每个具体的组件类型对 interface 加以实现
• 定义一个具体的工厂类，在构造器方法接受具体的组件类型，完成对应类型组件的构造

简单工厂模式的优势包括：
• 属于工厂模式最为简单直观的一种类型
• 构造各类组件时的聚拢收口效果最好，提供的公共切面最全面到位

存在的劣势为：
• 组件类扩展时，需要直接修改工厂的组件构造方法，不符合开闭原则
*/
