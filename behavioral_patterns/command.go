package main

import "fmt"

// 命令行模式：界面的按钮都是一条命令，模拟cmd。可能导致某些系统有过多的具体命令类，增加了系统的复杂度。

// 创建命令接口
type Command interface {
	//执行动作(操作)
	execute()
	//撤销动作(操作)
	undo()
}
type Light string

func (l *Light) on() {
	fmt.Println(" 电灯打开了.. ")
}

func (l *Light) off() {
	fmt.Println(" 电灯关闭了.. ")
}

type LightOnCommand struct {
	Light *Light //聚合Light
}

func (l *LightOnCommand) execute() {
	l.Light.on()
}

func (l *LightOnCommand) undo() {
	l.Light.off()
}

type LightOffCommand struct {
	Light *Light
}

func (l *LightOffCommand) execute() {
	l.Light.off()
}

func (l *LightOffCommand) undo() {
	l.Light.on()
}

// 没有任何命令，即空执行: 用于初始化每个按钮, 当调用空命令时，对象什么都不做.其实，这样是一种设计模式, 可以省掉对空判断
type NoCommand string

func (n *NoCommand) execute() {}

func (n *NoCommand) undo() {}

type remoteController struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command // 执行撤销的命令
}

func NewRemoteController() *remoteController {
	return &remoteController{
		onCommands:  []Command{new(NoCommand), new(NoCommand), new(NoCommand), new(NoCommand), new(NoCommand)},
		offCommands: []Command{new(NoCommand), new(NoCommand), new(NoCommand), new(NoCommand), new(NoCommand)},
	}
}

// 设置你需要的命令
func (r *remoteController) SetCommand(index int, do, undo Command) {
	r.onCommands[index] = do
	r.offCommands[index] = undo
}

// 按下开按钮
func (r *remoteController) OnButtonWasPushed(index int) {
	// 找到你按下的开的按钮， 并调用对应方法
	r.onCommands[index].execute()
	// 记录这次的操作，用于撤销
	r.undoCommand = r.onCommands[index]
}

// 按下开按钮
func (r *remoteController) OffButtonWasPushed(index int) {
	// 找到你按下的关的按钮， 并调用对应方法
	r.offCommands[index].execute()
	// 记录这次的操作，用于撤销
	r.undoCommand = r.offCommands[index]
}

func (r *remoteController) UndoButtonWasPushed() {
	r.undoCommand.undo()
}
