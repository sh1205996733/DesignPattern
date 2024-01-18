package main

import "fmt"

// 享元模式：解决重复对象的内存浪费问题，当系统中有大量的相似对象时，就需要缓存池。不需要总是创建新的对象，可以从缓冲池中拿，这样可以降低系统内存，提高效率
// 经典的应用场景就是池技术，String常量池，数据库连接池，缓冲池等
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
