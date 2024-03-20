// observer 观察者模式（发布-订阅者模式）
// https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247495132&idx=1&sn=0c42ff03123e188c4de44df1b67ef4de&chksm=fa833c4bcdf4b55da316c82a7edaf33fbeca427116ab57caec1501993a8ef5905ce04f399de3&scene=178&cur_album_id=2531498848431669249#rd
// https://refactoringguru.cn/design-patterns/observer?_gl=1*scgsby*_ga*Njc3MzQzNTA5LjE2ODk3NTc4NjE.*_ga_SR8Y3GYQYC*MTY4OTc1Nzg2MS4xLjEuMTY4OTc1OTA0Ny41NC4wLjA.
package main

import "fmt"

// Subject 发布者接口
type Subject interface {
	Subscribe(Observer)
	UnSubscribe(int)
	NotifyAll(string)
}

type Publisher struct {
	Name        string
	subscribers []Observer
}

// Subscribe 订阅
func (p *Publisher) Subscribe(observer Observer) {
	fmt.Println(observer.GetName(), " subscribe success")
	p.subscribers = append(p.subscribers, observer)
}

// UnSubscribe 取消订阅
func (p *Publisher) UnSubscribe(index int) {
	if index < 0 || index >= len(p.subscribers) {
		return
	}
	fmt.Println(p.subscribers[index].GetName(), " down")
	p.subscribers[index] = nil
}

// NotifyAll 通知订阅者
func (p *Publisher) NotifyAll(message string) {
	fmt.Printf("subscribers %#v\n", p.subscribers)
	fmt.Println(p.Name, " notify start")
	for _, sub := range p.subscribers {
		if sub == nil {
			continue
		}
		sub.Notify(message)
	}
	fmt.Println(p.Name, "  notify end")
}

// Observer 观察者接口 或者叫 Subscriber订阅者接口
type Observer interface {
	Notify(string)
	GetName() string
}

// PhoneObserver 短信订阅者
type PhoneObserver struct {
	Name string
}

func (p *PhoneObserver) Notify(msg string) {
	fmt.Println("PhoneObserver receive success msg: ", msg)
}
func (p *PhoneObserver) GetName() string {
	return p.Name
}

// EmailObserver 邮箱订阅者
type EmailObserver struct {
	Name string
}

func (p *EmailObserver) Notify(msg string) {
	fmt.Println("EmailObserver receive success msg: ", msg)
}
func (p *EmailObserver) GetName() string {
	return p.Name
}
