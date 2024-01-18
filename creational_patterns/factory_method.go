package main

import "fmt"

// 工厂方法模式 可扩展性不是很好
func FactoryMethod(objType string) string {
	var o string
	switch objType {
	case "a", "b":
		o = factoryAB(objType)
	case "c", "d":
		o = factoryCD(objType)
	}
	return o
}

func factoryAB(t string) string {
	fmt.Println("ab factory")
	var oj string
	switch t {
	case "a":
		oj = "A"
	case "b":
		oj = "B"
	}
	return oj
}

func factoryCD(t string) string {
	fmt.Println("cd factory")
	var oj string
	switch t {
	case "c":
		oj = "C"
	case "d":
		oj = "D"
	}
	return oj
}
