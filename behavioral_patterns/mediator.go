// mediator 中介者模式
// https://refactoringguru.cn/design-patterns/mediator?_gl=1*ac44gv*_ga*Njc3MzQzNTA5LjE2ODk3NTc4NjE.*_ga_SR8Y3GYQYC*MTY4OTgxNzkyMy40LjAuMTY4OTgxNzkyMy42MC4wLjA.
package main

/**
1、观察者模式：
察者模式属于行为型模式，指多个对象间存在一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
这种模式有时又称作发布-订阅模式、模型-视图模式。
2、中介者模式：
中介者模式，属于行为型模式；定义一个中介对象来封装一系列对象之间的交互，使原有对象之间的耦合松散，且可以独立地改变它们之间的交互。 ​
中介者模式又叫调停模式，它是迪米特法则的典型应用。

观察者模式：只能从从一的一方循环的通知，属于单向。（好比上课，只能老实授课给班上所有的学生，反过来不行）
中介者模式：可以从任一方循环通知，属于双向。（好比分享会，每个人都可以分享自己的事情给别人）
*/
import "fmt"

// 1、创建抽象工作者
type AbsMediator interface {
	publicHouse(string)
	needHouse(string)
}

// 2、创建抽象同事类【Person】，抽象同事有名字，有自己注册到哪一个中介公司中，所以需要聚合抽象中介者。
type Person interface {
	fabuxinxi(string)
	contact(string)
}

// 3、创建具体同事，房东【HouseOwner】和租户【Tenant】
type HouseOwner struct {
	mediator AbsMediator
}
type Tenant struct {
	mediator AbsMediator
}

func (p *HouseOwner) fabuxinxi(info string) {
	p.mediator.publicHouse(info)
}

func (p *HouseOwner) contact(info string) {
	fmt.Println(info)
}

func (p *Tenant) fabuxinxi(info string) {
	p.mediator.needHouse(info)
}

func (p *Tenant) contact(info string) {
	fmt.Println(info)
}

// 4. 创建中介
type Mediator struct {
	owners, tenants []Person
}

func (p *Mediator) publicHouse(info string) {
	for _, tenant := range p.tenants {
		tenant.contact("中介给租户推送:" + info)
	}
}

func (p *Mediator) needHouse(info string) {
	for _, owner := range p.owners {
		owner.contact("中介给房主推送:" + info)
	}
}
