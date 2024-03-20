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
func GetInstanceByLazyLoadV1() *singleton {
	if obj == nil { //此处存在多线程并发问题
		obj = new(singleton)
	}
	return obj
}
func GetInstanceByLazyLoadV2() *singleton {
	lock.Lock() //此处加锁效率太低
	defer lock.Unlock()
	if obj == nil { //此处存在多线程并发问题
		obj = new(singleton)
	}
	return obj
}
func GetInstanceByLazyLoadV3() *singleton {
	if obj != nil {
		return obj
	}
	lock.Lock() //此处加锁效率太低
	defer lock.Unlock()
	//此处存在多线程并发问题 会重复初始化
	obj = new(singleton)
	return obj
}
func GetInstanceByLazyLoadV4() *singleton {
	if obj != nil {
		return obj
	}
	lock.Lock() //此处加锁效率太低
	defer lock.Unlock()
	if obj != nil { //双检加锁
		return obj
	}
	obj = new(singleton)
	return obj
}

func GetInstanceByLazyLoad() *singleton {
	once.Do(func() { //推荐
		obj = new(singleton)
	})
	return obj
}

var obj2 = new(singleton)

// 饿汉式：线程安全
func GetInstance() *singleton {
	return obj2
}
