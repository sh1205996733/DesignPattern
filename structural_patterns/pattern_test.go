package main

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	p := new(Phone)
	p.Charging(new(VoltageAdapter))
	p.Charging(new(VoltageAdapter2))
}

func TestBridge(t *testing.T) {
	//phone := new(bridge.Android)
	phone := new(IPhone)
	//phone.Open()
	//phone.Call()
	//phone.Close()

	//phone1 := new(bridge.FoldedPhone)
	phone1 := new(UpRightPhone)
	phone1.Brand = phone
	phone1.Open()
	phone1.Call()
	phone1.Close()
}

func TestDecorator(t *testing.T) {
	// 装饰者模式下的订单：2份巧克力+一份牛奶的LongBlack
	// 1. 点一份 LongBlack
	var drink Drink = &LongBlackCoffee{Drink: NewCoffee("LongBlackCoffee", 10)}
	fmt.Println("费用1=", drink.Cost())
	fmt.Println("描述=" + drink.Desc())

	// 2. drink 加入一份牛奶
	drink = &MilkDecorator{Decorator: NewDecorator(NewCoffee("Milk", 2), drink)}
	fmt.Println("drink 加入一份牛奶 费用 =", drink.Cost())
	fmt.Println("drink 加入一份牛奶 描述 = " + drink.Desc())

	// 3. drink 加入一份巧克力

	drink = &ChocolateDecorator{Decorator: NewDecorator(NewCoffee("Chocolate", 3), drink)}

	fmt.Println("drink 加入一份牛奶 加入一份巧克力  费用 =", drink.Cost())
	fmt.Println("drink 加入一份牛奶 加入一份巧克力 描述 = " + drink.Desc())

	// 4 drink 加入一份巧克力

	drink = &ChocolateDecorator{Decorator: NewDecorator(NewCoffee("Chocolate", 3), drink)}

	fmt.Println("drink 加入一份牛奶 加入2份巧克力   费用 =", drink.Cost())
	fmt.Println("drink 加入一份牛奶 加入2份巧克力 描述 = " + drink.Desc())

	fmt.Println("===========================")

	var drink2 Drink = &DeCafCoffee{Drink: NewCoffee("DeCafCoffee", 1.5)}
	fmt.Println("drink2 无因咖啡  费用 =", drink2.Cost())
	fmt.Println("drink2 无因咖啡 描述 = " + drink2.Desc())

	drink2 = &MilkDecorator{Decorator: NewDecorator(NewCoffee("Milk", 2), drink2)}

	fmt.Println("drink2 无因咖啡 加入一份牛奶  费用 =", drink2.Cost())
	fmt.Println("drink2 无因咖啡 加入一份牛奶 描述 = " + drink2.Desc())
}

func TestComposite(t *testing.T) {
	//从大到小创建对象 学校
	var university Component = &University{component: NewComponent("清华大学", " 中国顶级大学 ")}

	//创建 学院
	var computerCollege Component = &College{component: component{"计算机学院", " 计算机学院 "}}
	var infoEngineercollege Component = &College{component: component{"信息工程学院", " 信息工程学院 "}}

	//创建各个学院下面的系(专业)
	computerCollege.Add(&Department{component: component{"软件工程", " 软件工程不错 "}})
	computerCollege.Add(&Department{component: component{"网络工程", " 网络工程不错 "}})
	computerCollege.Add(&Department{component: component{"计算机科学与技术", " 计算机科学与技术是老牌的专业 "}})

	//
	infoEngineercollege.Add(&Department{component: component{"通信工程", " 通信工程不好学 "}})
	infoEngineercollege.Add(&Department{component: component{"信息工程", " 信息工程好学 "}})

	//将学院加入到 学校
	university.Add(computerCollege)
	university.Add(infoEngineercollege)

	//university.Print()
	infoEngineercollege.Print()
}

func TestFacade(t *testing.T) {
	homeTheaterFacade := NewHomeTheaterFacade()
	homeTheaterFacade.Ready()
	homeTheaterFacade.Play()
	homeTheaterFacade.End()
}

func TestFlyweight(t *testing.T) {
	// 创建一个工厂类
	factory := new(WebSiteFactory)

	// 客户要一个以新闻形式发布的网站
	webSite1 := factory.GetWebSiteCategory("新闻")

	webSite1.Use(NewUser("tom"))

	// 客户要一个以博客形式发布的网站
	webSite2 := factory.GetWebSiteCategory("博客")

	webSite2.Use(NewUser("jack"))

	// 客户要一个以博客形式发布的网站
	webSite3 := factory.GetWebSiteCategory("博客")

	webSite3.Use(NewUser("smith"))

	// 客户要一个以博客形式发布的网站
	webSite4 := factory.GetWebSiteCategory("博客")

	webSite4.Use(NewUser("king"))

	fmt.Println("网站的分类共=", factory.GetWebSiteCount())
}

func TestProxy(t *testing.T) {
	// TODO Auto-generated method stub
	//创建目标对象(被代理对象)
	teacherDao := new(TeacherDao)

	//创建代理对象, 同时将被代理对象传递给代理对象
	teacherDaoProxy := &TeacherDaoProxy{Target: teacherDao}

	//通过代理对象，调用到被代理对象的方法
	//即：执行的是代理对象的方法，代理对象再去调用目标对象的方法
	teacherDaoProxy.Teach()
}
