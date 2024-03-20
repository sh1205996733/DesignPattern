package main

// 装饰者模式：是一种结构型设计模式，它允许通过将对象放入包装器中来为原始对象增强功能。
// 这些包装器（称为装饰器），能够在不改变原始对象代码的情况下，动态地为其添加新的行为和属性。
// 装饰者模式 VS 代理模式
/**

适配器模式（Adapter Pattern）和装饰器模式（Decorator Pattern）都是常见的结构型模式，它们用于在现有类的基础上增加新的功能或改变原有类的接口。它们的主要区别在于：

意图不同：适配器模式的主要意图是将一个已有的类的接口转换成客户所期望的另一个接口，以满足不同的需求。而装饰器模式的主要意图则是为对象动态地添加新的行为或责任。
对象类型不同：适配器模式通常使用组合方式来包装被适配者对象，从而对它进行转换；而装饰器模式则是包装同种类型的对象，使其能够动态地增加新的行为或责任。
适用场景不同：适配器模式适用于需要在保持原有接口、功能和实现的同时，通过对现有对象的适应来扩展新的功能。例如，将一种数据格式转换成另一种格式。
而装饰器模式适用于在不改变对象接口的情况下为对象动态地添加新的职责或行为。例如，为文本编辑器增加拼写检查、撤销操作等。

综上，适配器模式和装饰器模式虽然都是用于类的扩展和变换，但适配器模式更注重接口的转换和功能的实现；而装饰器模式则更注重对象的动态性，能够在运行时动态地为对象添加新的职责或行为。
*/

/**

桥接模式 VS 装饰器模式
桥接模式：主要是将抽象部分和具体部分分离，让他们可以相互独立的变化。桥接模式通过一个桥接接口来实现两个抽象体系的相互调用
装饰器模式：主要在不改变原有对象结构的前提下，通过创建一个包装对象，这样可以动态的为原有对象添加新的功能或者行为。

二者最主要的区别：
* 装饰器模式的应用场景通常是在需要在不改变原有类结构的情况下，为对象添加新的行为或功能。
* 而桥接模式的应用场景通常是在需要分离抽象部分和实现部分，并且这两部分独立进行扩展和变化的情况下。

总之，装饰器模式主要关注于为对象增加行为，而桥接模式主要强调从程序中消除复杂性并提高系统的灵活性。
*/

/**
代理模式 VS 装饰器模式

代理模式和装饰器模式虽然都是设计模式，但它们有着不同的目的和应用场景。以下是它们之间的主要区别：

* 目的不同。代理模式主要用于控制对对象的访问，而装饰器模式主要用于增强对象的功能。
* 应用场景不同。代理模式通常用于访问权限管理、资源控制等场景，它创建一个代理对象来控制对原对象的访问，而装饰器模式则在不改变原对象的前提下，动态地给对象添加新的功能或修改其行为。
* 实现方式不同。代理模式在编译时确定代理和真实对象之间的关系，而装饰器模式可以在运行时动态地构造新的功能
* 对对象的影响不同。代理模式对访问对象的控制更为严格，而装饰器模式则更加灵活，可以动态地添加或修改功能，而不影响对象的其他方面
简而言之，代理模式更关注于访问控制和流程管理，而装饰器模式更侧重于功能的增强和扩展。
*/
// 1、创建抽象构建【咖啡，Coffee】，提供获取名称和价格两个抽象方法
type Drink interface {
	Cost() float64
	Desc() string
}

// 2.创建具体组件，咖啡的一种具体类型【卡布奇诺，KabuqinuoCoffee】，具体组件需要实现抽象组件
type coffer struct {
	desc  string
	price float64
}

func (c *coffer) Cost() float64 {
	return c.price
}

func (c *coffer) Desc() string {
	return c.desc
}

func NewCoffee(desc string, price float64) Drink {
	return &coffer{desc: desc, price: price}
}

type DeCafCoffee struct {
	Drink
}
type ShortBlackCoffee struct {
	Drink
}
type LongBlackCoffee struct {
	Drink
}
type EspressoCoffee struct {
	Drink
}

// 3、创建抽象装饰器 将抽象组件作为成员变量聚合到抽象装饰器中 继承+组合
type Decorator struct {
	Drink
	origin Drink //增强的对象
}

func NewDecorator(coffer, drink Drink) Decorator {
	return Decorator{Drink: coffer, origin: drink}
}

// 4、声明具体装饰器
type MilkDecorator struct {
	Decorator
}

type ChocolateDecorator struct {
	Decorator
}

func (d *MilkDecorator) Cost() float64 {
	return d.Cost() + d.origin.Cost()
}
