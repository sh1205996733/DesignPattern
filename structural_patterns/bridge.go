package main

import "fmt"

// 桥接模式

// 品牌
type Brand interface {
	Open()
	Close()
	Call()
}

type Android string
type IPhone string

func (p Android) Open() {
	fmt.Println("Android Open...")
}
func (p Android) Close() {
	fmt.Println("Android Close...")
}
func (p Android) Call() {
	fmt.Println("Android Call...")
}

func (p IPhone) Open() {
	fmt.Println("IPhone Open...")
}
func (p IPhone) Close() {
	fmt.Println("IPhone Close...")
}
func (p IPhone) Call() {
	fmt.Println("IPhone Call...")
}

type FoldedPhone struct {
	Brand Brand
}

func (p FoldedPhone) Open() {
	fmt.Print("Folded ")
	p.Brand.Open()
}
func (p FoldedPhone) Close() {
	fmt.Print("Folded ")
	p.Brand.Close()
}
func (p FoldedPhone) Call() {
	fmt.Print("Folded ")
	p.Brand.Call()
}

type UpRightPhone struct {
	Brand Brand
}

func (p UpRightPhone) Open() {
	fmt.Print("UpRight ")
	p.Brand.Open()
}
func (p UpRightPhone) Close() {
	fmt.Print("UpRight ")
	p.Brand.Close()
}
func (p UpRightPhone) Call() {
	fmt.Print("UpRight ")
	p.Brand.Call()
}
