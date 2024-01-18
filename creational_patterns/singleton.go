// singleton 单例模式
// https://refactoringguru.cn/design-patterns/singleton

package main

import "sync"

// 1.私有化
type singleton struct {
}

var obj *singleton
var once sync.Once
var lock sync.Mutex

// 2.提供公共访问方法
// 懒汉式 线程不安全，可能造成内存浪费
func GetInstanceByLazyLoad() *singleton {
	//lock.Lock() 此处加锁效率太低
	//defer lock.Unlock()
	if obj == nil { //此处多线程存在问题
		//lock.Lock() 此处加锁效率太低
		//defer lock.Unlock()
		//if obj == nil { //双重检查
		//	obj = new(singleton)
		//}
		once.Do(func() { //推荐
			obj = new(singleton)
		})
	}
	return obj
}

var obj2 = new(singleton)

// 饿汉式：线程安全
func GetInstance() *singleton {
	return obj2
}
