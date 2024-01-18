package main

import (
	"fmt"
	"testing"
)

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

}

func TestObserver(t *testing.T) {

}

func TestMediator(t *testing.T) {

}
