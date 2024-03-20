package main

import "fmt"

// 抽象工厂模式:将考虑多等级产品的生产，将同一个具体工厂所生产的位于不同等级的一组产品称为一个产品族。
//		抽象工厂 (Abstract Factory) ：定义了一个接口，用于创建不同种类的产品。
//		具体工厂 (Concrete Factory) ：实现抽象工厂接口，并负责实例化产品。
//		抽象产品 (Abstract Product) ：定义产品的基本属性和方法。
//		具体产品 (Concrete Product) ：实现抽象产品接口，即给出具体的实现。

// 1、创建抽象产品，咖啡【Coffee】和甜点【Dessert】抽象类
type Dessert interface {
	show()
}

// 2、创建抽象工厂接口【AbstractFactory】，里面提供生产咖啡和甜点的抽象方法
type AbstractFactory interface {
	createCoffer() Coffer
	createDessert() Dessert
}

// 3、创建具体产品类
//
//	具体的甜点【提拉米苏（TiLaMiSu）和抹茶慕斯（MoChaMuSi）】
//	具体的咖啡【拿铁（NatieCoffee）和美式（AmericaCoffee）】
type TiLaMiSu struct {
}

func (t *TiLaMiSu) show() {
	fmt.Println("TiLaMiSu")
}

type MoChaMuSi struct {
}

func (t *MoChaMuSi) show() {
	fmt.Println("MoChaMuSi")
}

// 4、创建具体的工厂
//
//	【美式风格工厂（AmericaFactory）和意大利风格工厂（ItalyFactory）】
type AmericaFactory struct {
}

func (f *AmericaFactory) createCoffer() Coffer {
	return &MeishiCoffee{}
}

func (f *AmericaFactory) createDessert() Dessert {
	return &TiLaMiSu{}
}

type ItalyFactory struct {
}

func (f *ItalyFactory) createCoffer() Coffer {
	return &NatieCoffee{}
}

func (f *ItalyFactory) createDessert() Dessert {
	return &MoChaMuSi{}
}

/**
工厂方法模式和抽象工厂模式都是创建型设计模式，它们的主要区别在于所创建的对象范围不同。

工厂方法模式（Factory Method）通过让子类实现工厂接口，来决定具体应该创建哪一个产品类的实例对象。它允许我们在不改变现有代码基础上添加新的产品类型，并且可以将具体产品的实现与调用方分离开来。
抽象工厂模式（Abstract Factory）与工厂方法模式类似，也是用于创建一系列相关的对象。不同之处在于，抽象工厂是针对多个产品族而言的，即每个工厂可以创建多种不同类型的产品。这样的话，抽象工厂为创建一组相关或独立的对象提供了一种方式。

另外，工厂方法模式通常只针对一个抽象产品类进行创建，而抽象工厂模式则需要针对多种抽象产品进行创建。此外，使用工厂方法模式时，往往仅需选用所需的工厂方法即可，而使用抽象工厂模式时，则需要创建必须的所有抽象产品对象。
*/
