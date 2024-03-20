package main

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	// 假设这里获取数据，以及数据是否敏感
	strategyType := "file"
	strategyType = "encrypt_file"
	storage, _ := NewStorageStrategy(strategyType)
	storage.Save()
}

func TestTemplate(t *testing.T) {
	//制作红豆豆浆
	fmt.Println("----制作红豆豆浆----")
	var redBeanSoyaMilk = &SoyaMilk{Soya: new(RedBeanSoyaMilk)}
	//var redBeanSoyaMilk = &SoyaMilk{new(RedBeanSoyaMilk).AddCondiment}
	redBeanSoyaMilk.Make()

	fmt.Println("----制作花生豆浆----")
	//var peanutSoyaMilk = &SoyaMilk{Soya: new(PeanutSoyaMilk)}
	var peanutSoyaMilk = new(PeanutSoyaMilk)
	//var peanutSoyaMilk Soya = &SoyaMilk{new(PeanutSoyaMilk).AddCondiment}
	peanutSoyaMilk.Make()

}

func TestCommand(t *testing.T) {
	//使用命令设计模式，完成通过遥控器，对电灯的操作

	//创建电灯的对象(接受者)
	light := new(Light)

	//创建电灯相关的开关命令
	lightOnCommand := &LightOnCommand{light}
	lightOffCommand := &LightOffCommand{light}

	//需要一个遥控器
	remoteController := NewRemoteController()

	//给我们的遥控器设置命令, 比如 no = 0 是电灯的开和关的操作
	remoteController.SetCommand(0, lightOnCommand, lightOffCommand)

	fmt.Println("--------按下灯的开按钮-----------")
	remoteController.OnButtonWasPushed(0)
	fmt.Println("--------按下灯的关按钮-----------")
	remoteController.OffButtonWasPushed(0)
	fmt.Println("--------按下撤销按钮-----------")
	remoteController.UndoButtonWasPushed()

	fmt.Println("=========使用遥控器操作电视机==========")

	//TV tv = new TV();
	//
	//TVOffCommand tvOffCommand = new TVOffCommand(tv);
	//TVOnCommand tvOnCommand = new TVOnCommand(tv);
	//
	////给我们的遥控器设置命令, 比如 no = 1 是电视机的开和关的操作
	//remoteController.setCommand(1, tvOnCommand, tvOffCommand);
	//
	//fmt.Println("--------按下电视机的开按钮-----------");
	//remoteController.onButtonWasPushed(1);
	//fmt.Println("--------按下电视机的关按钮-----------");
	//remoteController.offButtonWasPushed(1);
	//fmt.Println("--------按下撤销按钮-----------");
	//remoteController.undoButtonWasPushed();
}

func TestVisitor(t *testing.T) {
	var circle = new(Circle)
	rectangle := new(Rectangle)
	//计算面积
	areas := new(AreaCalculator)
	circle.calc(areas)
	rectangle.calc(areas)
	//计算周长
	perimeters := new(PerimeterCalculator)
	circle.calc(perimeters)
	rectangle.calc(perimeters)
}

func TestObserver(t *testing.T) {
	var publish Subject
	publish = &Publisher{Name: "Publisher"}
	sub1 := &PhoneObserver{Name: "PhoneObserver"}
	sub2 := &EmailObserver{Name: "EmailObserver"}
	publish.Subscribe(sub1)
	publish.Subscribe(sub2)
	publish.NotifyAll("first")
	publish.UnSubscribe(1)
	publish.NotifyAll("second")
}

func TestMediator(t *testing.T) {
	mediator := new(Mediator)
	owner := &HouseOwner{mediator}
	tenant := &Tenant{mediator}
	mediator.owners, mediator.tenants = []Person{owner}, []Person{tenant}
	//owner.fabuxinxi("商户发布出租信息")
	tenant.fabuxinxi("租户发布租房信息")
}

func TestState(t *testing.T) {
	// 创建活动对象，奖品有1个奖品
	activity := &RaffleActivity{count: 1}
	activity.state = GetNoRaffleState(activity)
	// 我们连续抽300次奖
	for i := 0; i < 30; i++ {
		fmt.Println("--------第", i+1, "次抽奖----------")
		// 参加抽奖，第一步点击扣除积分
		activity.debuctMoney()
		// 第二步抽奖
		activity.raffle()
	}
}

func TestResponsibilitychain(t *testing.T) {
	//创建一个请求
	purchaseRequest := &PurchaseRequest{1, 5000, 1}

	//创建相关的审批人
	departmentApprover := &DepartmentApprover{BaseApprover: BaseApprover{Name: "张主任"}}
	collegeApprover := &CollegeApprover{BaseApprover: BaseApprover{Name: "李院长"}}
	viceSchoolMasterApprover := &ViceSchoolMasterApprover{BaseApprover: BaseApprover{Name: "王副校"}}
	schoolMasterApprover := &SchoolMasterApprover{BaseApprover: BaseApprover{Name: "佟校长"}}

	//需要将各个审批级别的下一个设置好 (处理人构成环形: )
	departmentApprover.SetApprover(collegeApprover)
	collegeApprover.SetApprover(viceSchoolMasterApprover)
	viceSchoolMasterApprover.SetApprover(schoolMasterApprover)
	schoolMasterApprover.SetApprover(departmentApprover)

	//departmentApprover.ProcessRequest(purchaseRequest)
	viceSchoolMasterApprover.ProcessRequest(purchaseRequest)
}

func TestIterator(t *testing.T) {
	stu, stu2, stu3 := new(Student), new(Student), new(Student)
	list := new(StudentAggregate)
	list.addStudent(stu, stu2, stu3)
	iterator := list.getIterator()
	for iterator.hasNext() {
		fmt.Println(iterator.next())
	}
}

func TestMemento(t *testing.T) {

}

func TestInterpreter(t *testing.T) {
	//cxt := Context(1)
	//expression := &Minus{&Plus{&Plus{&Plus{Context(1), 2}, c}, d}, e}
	//fmt.Println(expression.interpret(cxt))
	// ((((a + b) + c) + d) - e)= 5
}
