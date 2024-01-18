package main

type ABFactory struct {
}

func (f *ABFactory) CreateObj(objType string) interface{} {
	return factoryAB(objType)
}

type CDFactory struct {
}

func (f *CDFactory) CreateObj(objType string) interface{} {
	return factoryCD(objType)
}

// 定义一个抽象工程方法，由其他工厂实现
type Factory interface {
	CreateObj(string) interface{}
}

func AbstractFactory(objType string) Factory {
	var o Factory
	switch objType {
	case "ab":
		o = new(ABFactory)
	case "cd":
		o = new(CDFactory)
	}
	return o
}
