package main

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	obj1 := GetInstance()
	obj2 := GetInstance()
	obj3 := GetInstance()
	obj4 := GetInstance()
	fmt.Printf("%p,%p,%p,%p\n", obj1, obj2, obj3, obj4)
	fmt.Println(obj1, obj2, obj3, obj4)
}

func TestSimpleFactory(t *testing.T) {
	obj := SimpleFactory("pepper")
	fmt.Println(obj)
}

func TestFactoryMethod(t *testing.T) {
	obj := FactoryMethod("c")
	fmt.Println(obj)
}

func TestAbstractFactory(t *testing.T) {
	//var f factory.Factory = factory.AbstractFactory("ab")
	var f Factory = AbstractFactory("cd")
	obj := f.CreateObj("a")
	fmt.Println(obj)
}

func TestPrototype(t *testing.T) {
	p1 := Proto{Name: "p1", Favs: []string{"play", "eat", "sleep"}}
	//p2 := prototype.Proto{Name: p1.Name, Favs: p1.Favs} //传统方式
	p2 := p1.Clone()
	p2.Favs[0] = "study"
	fmt.Println(p1.Favs, p2.Favs)
}

func TestBuilder(t *testing.T) {
	//p1 := &builder.CommonProduct{}
	p2 := &HeightProduct{}
	d := &Director{Builder: p2}
	p := d.Build()
	fmt.Println(p)
}
