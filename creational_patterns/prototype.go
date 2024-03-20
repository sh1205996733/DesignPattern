package main

import "encoding/json"

// 原型模式 一般推荐使用对象序列化的方式来实现
// 传统方式 字段复制原对象的字段，比较麻烦，而且是浅拷贝，（引用类型只拷贝地址，实际两个指向的还是同一个内存地址）

// 缺点每个类都要提供一个clone方法，违背了ocp原则
type Proto struct {
	Name     string
	Favs     []string
	Subjects []string `json:"-"` // - 字段不会被序列化
}
type Cloneable interface {
	Clone() *Proto
}

func (p *Proto) Clone() *Proto {
	str, _ := json.Marshal(p)
	var copy *Proto
	json.Unmarshal(str, &copy)
	return copy
}
