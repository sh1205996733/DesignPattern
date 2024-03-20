package main

import (
	"fmt"
	"github.com/bmizerany/assert"
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

func TestSingleton2(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

func BenchmarkSingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if GetInstance() != GetInstance() {
			b.Error("test fail")
		}
	}
}

func TestSimpleFactory(t *testing.T) {
	obj := SimpleFactory("pepper")
	fmt.Println(obj)
}

func TestFactoryMethod(t *testing.T) {
	var f CoffeeFactory = &MeishiCoffeeFactory{}
	//var f CoffeeFactory = &NatieCoffeeFactory{}
	cli := &FactoryMethod{f}
	cli.OrderCoffer()
}

func TestAbstractFactory(t *testing.T) {
	var f AbstractFactory = new(ItalyFactory)
	//var f AbstractFactory = new(AmericaFactory)
	c := f.createCoffer()
	c.AddTang()
	c.AddNai()

	d := f.createDessert()
	d.show()
}

func TestPrototype(t *testing.T) {
	p1 := Proto{Name: "p1", Favs: []string{"play", "eat", "sleep"}}
	//p2 := prototype.Proto{Name: p1.Name, Favs: p1.Favs} //传统方式
	p2 := p1.Clone()
	p2.Favs[0] = "study"
	fmt.Println(p1.Favs, p2.Favs)
}

func TestBuilder(t *testing.T) {
	//var p1 Builder = &builder.CommonProduct{}
	var p2 Builder = &HeightProduct{}
	d := &Director{Builder: p2}
	p := d.Build()
	fmt.Println(p)
}

func TestOptions(t *testing.T) {
	p := InitOptions(WithStr1("str1"), WithStr2("str2"))
	fmt.Println(p)
}
