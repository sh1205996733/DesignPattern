package main

import "fmt"

// 建造者模式：将产品和创建过程解藕，使得相同的创建过程可以创建不同的产品
// 建造者模式 VS 抽象工厂模式
// 	抽象工厂模式:实现对产品家族的创建，具有不同分类的产品组合（例如，同一个手机有不同的厂家）
// 	建造者模式:按照相同的步骤生产不同的产品（例如，同样的步骤盖不同的房子）

type Product struct {
	Name string
}

type Builder interface {
	buildBasic()
	buildWalls()
	roofed()
	build() any
}

type Director struct {
	Builder Builder
}

func (d *Director) Build() any {
	d.Builder.buildBasic()
	d.Builder.buildWalls()
	d.Builder.roofed()
	return d.Builder.build()
}

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

func (c *CommonProduct) build() any {
	c.Name = "Common"
	return c
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

func (h *HeightProduct) build() any {
	h.Name = "height"
	return h
}
