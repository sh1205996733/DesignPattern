package main

// 简单工厂模式 可扩展性不是很好
func SimpleFactory(objType string) string {
	var o string
	if objType == "greek" {
		o = "greek"
	} else if objType == "cheese" {
		o = "cheese"
	} else if objType == "pepper" {
		o = "pepper"
	}
	return o
}
