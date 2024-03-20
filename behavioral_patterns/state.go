package main

import (
	"fmt"
	"math/rand"
)

// 状态模式：解决对象存在多种状态转换时，需要对外输出不同的行为的问题 状态模式允许对象在运行时根据其内部状态选择它的行为，从而使得对象看起来似乎改变了它的类。
// 扩展性很差。如果新加了断电的状态，我们需要修改上面判断逻辑
// 状态抽象类
type State interface {
	// 扣除积分 - 50
	deductMoney()
	// 是否抽中奖品
	raffle() bool
	// 发放奖品
	dispensePrize()
}

type RaffleActivity struct {
	// state 表示活动当前的状态，是变化
	state State
	// 奖品数量
	count int
}

func (a *RaffleActivity) setState(state State) {
	a.state = state
}

// 扣分, 调用当前状态的 deductMoney
func (a *RaffleActivity) debuctMoney() {
	a.state.deductMoney()
}

// 抽奖
func (a *RaffleActivity) raffle() {
	// 如果当前的状态是抽奖成功
	if a.state.raffle() {
		//领取奖品
		a.state.dispensePrize()
	}
}

// 这里请大家注意，每领取一次奖品，count--
func (a *RaffleActivity) getCount() int {
	curCount := a.count
	a.count--
	return curCount
}

func GetNoRaffleState(ac *RaffleActivity) State {
	return &NoRaffleState{ac}
}

func GetCanRaffleState(ac *RaffleActivity) State {
	return &CanRaffleState{ac}
}

func GetDispenseState(ac *RaffleActivity) State {
	return &DispenseState{ac}
}

func GetDispenseOutState(ac *RaffleActivity) State {
	return &DispenseOutState{ac}
}

// 不能抽奖状态
type NoRaffleState struct {
	activity *RaffleActivity // 初始化时传入活动引用，扣除积分后改变其状态
}

// 当前状态可以扣积分 , 扣除后，将状态设置成可以抽奖状态
func (n *NoRaffleState) deductMoney() {
	fmt.Println("扣除50积分成功，您可以抽奖了")
	n.activity.setState(GetCanRaffleState(n.activity))
}

// 当前状态不能抽奖
func (n *NoRaffleState) raffle() bool {
	fmt.Println("扣了积分才能抽奖喔！")
	return false
}

// 当前状态不能发奖品
func (n *NoRaffleState) dispensePrize() {
	fmt.Println("不能发放奖品")
}

// 可以抽奖的状态
type CanRaffleState struct {
	activity *RaffleActivity // 初始化时传入活动引用，扣除积分后改变其状态
}

// 已经扣除了积分，不能再扣
func (n *CanRaffleState) deductMoney() {
	fmt.Println("已经扣取过了积分!")
	n.activity.setState(GetCanRaffleState(n.activity))
}

// 可以抽奖, 抽完奖后，根据实际情况，改成新的状态
func (n *CanRaffleState) raffle() bool {
	fmt.Println("正在抽奖，请稍等！")
	num := rand.Intn(10)
	// 10%中奖机会
	if num == 0 {
		// 改变活动状态为发放奖品 context
		n.activity.setState(GetDispenseState(n.activity))
		return true
	} else {
		fmt.Println("很遗憾没有抽中奖品！")
		// 改变状态为不能抽奖
		n.activity.setState(GetNoRaffleState(n.activity))
		return false
	}
}

// 当前状态不能发奖品
func (n *CanRaffleState) dispensePrize() {
	fmt.Println("没中奖，不能发放奖品")
}

// 发放奖品的状态
type DispenseState struct {
	activity *RaffleActivity // 初始化时传入活动引用，扣除积分后改变其状态
}

func (n *DispenseState) deductMoney() {
	fmt.Println("不能扣除积分")
}

// 当前状态不能抽奖
func (n *DispenseState) raffle() bool {
	fmt.Println("不能抽奖！")
	return false
}

// 发奖品
func (n *DispenseState) dispensePrize() {
	if n.activity.getCount() > 0 {
		fmt.Println("恭喜中奖了")
		// 改变状态为不能抽奖
		n.activity.setState(GetNoRaffleState(n.activity))
	} else {
		fmt.Println("很遗憾，奖品发送完了")
		// 改变状态为奖品发送完毕, 后面我们就不可以抽奖
		n.activity.setState(GetDispenseOutState(n.activity))
	}
}

// 奖品发放完毕状态
// 说明，当我们activity 改变成 DispenseOutState， 抽奖活动结束
type DispenseOutState struct {
	activity *RaffleActivity // 初始化时传入活动引用，扣除积分后改变其状态
}

func (n *DispenseOutState) deductMoney() {
	fmt.Println("奖品发送完了，请下次再参加")
}

// 当前状态不能抽奖
func (n *DispenseOutState) raffle() bool {
	fmt.Println("奖品发送完了，请下次再参加")
	return false
}

// 当前状态不能发奖品
func (n *DispenseOutState) dispensePrize() {
	fmt.Println("奖品发送完了，请下次再参加")
}
