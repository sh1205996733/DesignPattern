package main

import "fmt"

// 享元模式：解决重复对象的内存浪费问题，当系统中有大量的相似对象时，就需要缓存池。不需要总是创建新的对象，可以从缓冲池中拿，这样可以降低系统内存，提高效率
// 经典的应用场景就是池技术，String常量池，数据库连接池，缓冲池等
// 享元模式通过共享对象来减少重复对象的创建。从而提高系统的效率和性能。

/*
单例模式通过一个类只能创建一个对象的机制来保证全局唯一性。典型的实现方式是使用静态成员变量或枚举类型，并将构造函数设为私有，避免外部直接创建对象。这样就可以在整个系统中共享该对象，避免重复创建对象所带来的开销。
享元模式则是通过共享对象来减少内存使用和提高系统性能。享元对象分为两类：内部状态和外部状态，其中内部状态指的是那些多个对象之间可以共享的属性，而外部状态则是那些因对象而异的属性。享元模式通过维护一个对象缓存池，将内部状态作为一个公共资源进行共享，从而避免了创建大量重复对象的开销。

简单来说，单例模式关注的是全局唯一性，而享元模式关注的是对象共享。两种模式的应用场景并不完全相同，需要根据实际情况选择合适的模式来提高程序效率和性能。
*/
type WebSite interface {
	Use(*user)
}
type user struct {
	name string
}

func NewUser(name string) *user {
	return &user{name: name}
}

type WebSiteFactory struct {
	m map[string]WebSite
}

func (w *WebSiteFactory) GetWebSiteCategory(typ string) WebSite {
	if w.m == nil {
		w.m = make(map[string]WebSite)
	}
	if _, ok := w.m[typ]; !ok {
		w.m[typ] = &ConcreteWebSite{typ: typ}
	}
	return w.m[typ]
}

func (w *WebSiteFactory) GetWebSiteCount() int {
	return len(w.m)
}

type ConcreteWebSite struct {
	typ string
}

func (c *ConcreteWebSite) Use(u *user) {
	fmt.Println("网站的发布形式为:", c.typ, " 在使用中 .. 使用者是", u.name)
}
