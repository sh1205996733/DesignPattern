package main

import "fmt"

// 工厂方法模式只考虑生产同等级的产品。
/**
工厂方法模式通常由以下几种角色组成：
	抽象产品 (Product) ：定义产品的属性和行为。
	具体产品 (Concrete Product) ：实现抽象产品接口的具体类。
	抽象工厂 (Factory) ：定义工厂的属性和行为以及返回实例化产品的方法。
	具体工厂 (Concrete Factory) ：实现抽象工厂接口的具体类。
*/
// 定义一个抽象工程方法，由其他工厂实现  将具体的产品实现与其它部分的代码分离开来，同时也具备扩展方便、可维护性高等优点。

// 1、创建一个抽象产品类，定义产品的基本属性与方法
type Coffer interface {
	AddNai()
	AddTang()
}

// 2、创建一个抽象工厂类，在其中定义一个创建产品的抽象方法。 咖啡工厂【CoffeeFactory】：接口或者抽象类，里面提供生产咖啡的抽象方法
type CoffeeFactory interface {
	MakeCoffee() Coffer
}

// 3、创建具体产品类，实现抽象产品接口，即定义具体的产品类型。
//
//	美式咖啡【MeishiCoffee】：Coffee 的子类，具体化 Coffee 中的两个方法
//	拿铁咖啡【NatieCoffee】：Coffee 的子类，具体化 Coffee 中的两个方发
type MeishiCoffee struct {
}

func (m MeishiCoffee) AddNai() {
	fmt.Println("MeishiCoffee AddNai")
}

func (m MeishiCoffee) AddTang() {
	fmt.Println("MeishiCoffee AddTang")
}

type NatieCoffee struct {
}

func (m NatieCoffee) AddNai() {
	fmt.Println("NatieCoffee AddNai")
}

func (m NatieCoffee) AddTang() {
	fmt.Println("NatieCoffee AddTang")
}

// 4、创建一个具体工厂类，实现抽象工厂接口，即具体决定应该创建哪种产品。
//
//	美式咖啡工厂【MeishiCoffeeFactory】：制作美式咖啡
//	拿铁咖啡工厂【NatieCoffeeFactory】：制作拿铁咖啡
type MeishiCoffeeFactory struct {
}

func (m MeishiCoffeeFactory) MakeCoffee() Coffer {
	fmt.Println("MeishiCoffeeFactory")
	return &MeishiCoffee{}
}

type NatieCoffeeFactory struct {
}

func (m NatieCoffeeFactory) MakeCoffee() Coffer {
	fmt.Println("NatieCoffeeFactory")
	return &NatieCoffee{}
}

type FactoryMethod struct {
	f CoffeeFactory
}

// FactoryMethod
func (ab *FactoryMethod) OrderCoffer() Coffer {
	coffer := ab.f.MakeCoffee()
	coffer.AddTang()
	coffer.AddNai()
	return coffer
}
