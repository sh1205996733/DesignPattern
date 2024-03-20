package main

import "fmt"

// 模版方法模式：当完成某个过程，该过程要执行一系列步骤，但这个步骤基本相同，但是只有其中个别步骤在实现时存在差异，通常考虑使用模版方法来实现

// 总之，模板方法模式适用于当算法的基本骨架已经确定，但某些特定步骤的实现可能会在子类中发生变化的场景，它能通过抽象出模板方法来封装算法过程，并通过抽象方法延迟到子类中实现特定步骤，从而提高代码的重用性和扩展性。

/*
外观模式是结构型模式，旨在提供一个统一的接口，用于访问子系统中的各种不同的类。通过提供一个简单的统一接口，外观模式可以简化客户端与子系统之间的交互。外观模式将多个复杂的子系统封装起来，并暴露出一个简单的统一 API，使得客户端可以更容易地使用这些功能而不需要了解其内部实现。
模板方法模式是行为型模式。相对而言，模板方法模式则是定义一个算法的骨架，而将特定步骤的实现延迟到子类中。该模式使得子类可以重载算法的某些步骤，从而不改变算法的整体结构。模板方法模式通常用于编写框架或库的核心代码。它提供了一个结构性框架，以便针对具体情况进行扩展。
因此，在两者之间的区别特点方面，外观模式是更集中的抽象，它旨在简化接口、集成和调用流程；而模板方法模式是一个纵向的抽象级别，它旨在优化代码或设计框架，使得算法的可维护性和扩展性更好。
*/
type Soya interface {
	AddCondiment()
}
type SoyaMilk struct {
	Soya //方法一：内嵌接口，由子类去实现
	//AddCondiment func() // 定义函数，
}

func (s *SoyaMilk) choose() {
	fmt.Println("choose")
}

func (s *SoyaMilk) soak() {
	fmt.Println("soak")
}

func (s *SoyaMilk) beat() {
	fmt.Println("beat")
}

func (s *SoyaMilk) Make() {
	s.choose()
	if s.Soya != nil {
		s.Soya.AddCondiment()
	}
	//if s.AddCondiment != nil {
	//	s.AddCondiment()
	//}
	s.beat()
	s.soak()
}

type PeanutSoyaMilk struct {
	SoyaMilk
}

// PeanutSoyaMilk PeanutSoyaMilk
func NewPeanutSoyaMilk() *PeanutSoyaMilk {
	tel := &PeanutSoyaMilk{}
	// 这里有点绕，是因为 go 没有继承，用嵌套结构体的方法进行模拟
	// 这里将子类作为接口嵌入父类，就可以让父类的模板方法 Send 调用到子类的函数
	// 实际使用中，我们并不会这么写，都是采用组合+接口的方式完成类似的功能
	tel.SoyaMilk = SoyaMilk{Soya: tel}
	return tel
}

func (p *PeanutSoyaMilk) AddCondiment() {
	fmt.Println("Add Peanut...")
}

// 推荐使用这种的
type RedBeanSoyaMilk string

func (r *RedBeanSoyaMilk) AddCondiment() {
	//if !r.CustomerWantCondiments() {
	//	return
	//}
	fmt.Println("Add RedBean...")
}

// 钩子方法，决定是否需要添加配料
func (r *RedBeanSoyaMilk) CustomerWantCondiments() bool {
	return false
}
