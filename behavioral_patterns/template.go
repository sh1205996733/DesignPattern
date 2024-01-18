package main

import "fmt"

// 模版方法模式：当完成某个过程，该过程要执行一系列步骤，但这个步骤基本相同，但是只有其中个别步骤在实现时存在差异，通常考虑使用模版方法来实现
type Soya interface {
	AddCondiment()
}
type SoyaMilk struct {
	Soya Soya //方法一：内嵌接口，由父类去实现
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
