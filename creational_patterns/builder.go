package main

import "fmt"

//建造者模式（Builder Pattern）是一种创建型设计模式，将产品和创建过程解藕，使得相同的创建过程可以创建不同的产品。
//
//在建造者模式中，定义一个Director（指挥者）对象，其职责是按照一定的顺序调用Builder对象来构建最终的对象。而Builder（建造者）对象则负责实际构建对象的各个部分，对外隐藏了具体的实现细节。 同时， Builder 将返回构建的对象给 Director 使用。
//
//建造者模式的组成部分：
//抽象建造者（Builder）：接口或者抽象类，目的是实现复杂对象有哪些部分需要进行建造，不提供建造细节
//具体建造者：抽象建造者的实现类，将抽象建造者里面的方法具体化
//产品：需要生产的东西
//指挥者（Director）：调用具体建造者完成各个部分

// 建造者模式 VS 抽象工厂模式
//
//	抽象工厂模式:实现对产品家族的创建，具有不同分类的产品组合（例如，同一个厂家有不同的手机）,用于创建一些简单对象，而且通常只需要一步就可以创建完成
//	建造者模式:按照相同的步骤生产不同的产品（例如，同样的步骤盖不同的房子）
//		对象复杂度不同：工厂方法模式用于创建简单对象，而建造者模式用于创建相对复杂的对象；
//		过程不同：工厂方法模式通常只需要一步就可以创建完成，而建造者模式会分成多个步骤、多次操作来构建一个完整的对象，增加了可定制性和灵活性；
//		结构不同：工厂方法模式是选择性地隐藏了对象创建的细节，通过向用户暴露工厂接口来创建对象；而建造者模式则直接将对象构造和表示分开，隔离了对象的创建和使用。

// 例如：当我们要考虑生产一个超人；
// 若使用工厂方法，模式，一般将生产超人和生产机器人放在一起比较，我们提供生产超人的工厂就可以将超人生产出来了。
// 若使用建造者模式考虑，就是将超人分成头、手、脚等等部分生产好，然后进行组装好。最后穿上内裤就可以了，在这个过程中就需要一个指挥者了，用于指定生产顺序。
//
// 1.角色一：产品
type Product struct {
	Name string
}

// 2. 角色二：抽象建造者
type Builder interface {
	buildBasic()
	buildWalls()
	roofed()
	build() Product
}

// 3. 角色四：具体建造者
type CommonProduct struct {
	Product
}

func (c *CommonProduct) buildBasic() {
	fmt.Println("common buildBasic")
}

func (c *CommonProduct) buildWalls() {
	fmt.Println("common buildWalls")
}

func (c *CommonProduct) roofed() {
	fmt.Println("common roofed")
}

func (c *CommonProduct) build() Product {
	return Product{Name: "common"}
}

type HeightProduct struct {
	Product
}

func (h *HeightProduct) buildBasic() {
	fmt.Println("height buildBasic")
}

func (h *HeightProduct) buildWalls() {
	fmt.Println("height buildWalls")
}

func (h *HeightProduct) roofed() {
	fmt.Println("height roofed")
}

func (h *HeightProduct) build() Product {
	return Product{Name: "height"}
}

// 4.角色三：指挥者
type Director struct {
	Builder Builder
}

func (d *Director) Build() Product {
	d.Builder.buildBasic()
	d.Builder.buildWalls()
	d.Builder.roofed()
	return d.Builder.build()
}
