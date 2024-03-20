package main

import "fmt"

// 访问者模式 访问者模式是一种行为型设计模式，它可以将算法与对象结构分离开来，使得算法可以独立于对象来变化。
// 该模式主要解决的问题是在不修改对象自身的基础上，对对象的结构进行增删改等操作。
// 在访问者模式中，对象结构包含多个具体元素，每个具体元素都可以接受一个访问者进行访问，而访问者可以对不同的具体元素进行不同的操作。
// 1、抽象访问者 首先我们需要定义图形元素的基本接口，包括接受访问者的方法
type Element interface {
	calc(Visitor)
}

// 2、然后我们定义几个具体访问类
type Circle struct {
}

func (e *Circle) calc(visitor Visitor) { //计算
	visitor.visitCircle(e)
}

type Rectangle struct {
}

func (e *Rectangle) calc(visitor Visitor) {
	visitor.visitRectangle(e)
}

// 3、接下来，我们定义一个访问者接口
type Visitor interface {
	visitCircle(Element)
	visitRectangle(Element)
}

// 4、最后，我们定义具体的访问者实现
type AreaCalculator string // 计算面积的访问者
func (v *AreaCalculator) visitCircle(c Element) {
	fmt.Println("打印圆的面积", c)
}
func (v *AreaCalculator) visitRectangle(c Element) {
	fmt.Println("打印矩形的面积", c)
}

type PerimeterCalculator string //计算周长的访问者
func (v *PerimeterCalculator) visitCircle(c Element) {
	fmt.Println("打印圆的周长", c)
}
func (v *PerimeterCalculator) visitRectangle(c Element) {
	fmt.Println("打印矩形的周长", c)
}
