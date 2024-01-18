package main

// 装饰者模式：动态的将新功能加到对象上,增强原始对象功能
// 装饰者模式 VS 代理模式
type Drink interface {
	Cost() float64
	Desc() string
}

// 被装饰者
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

// 装饰者 继承+组合
type Decorator struct {
	Drink
	drink Drink
}

func NewDecorator(coffer, drink Drink) Decorator {
	return Decorator{Drink: coffer, drink: drink}
}

type MilkDecorator struct {
	Decorator
}

type SoyDecorator struct {
	Decorator
}

type ChocolateDecorator struct {
	Decorator
}

func (d *Decorator) Cost() float64 {
	return d.Drink.Cost() + d.drink.Cost()
}
