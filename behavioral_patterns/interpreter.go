package main

import "fmt"

//解释器模式

// 1、创建抽象表达式
type Expression interface {
	interpret(Context) int
}

// 2、创建终结符表达式
type Value struct {
	value int
}

func (e *Value) interpret(ctx Context) int {
	return e.value
}

// 3、创建非终结符表达式
type Minus struct {
	left  Expression
	right Expression
}

func (e *Minus) interpret(cxt Context) int {
	fmt.Println("(", e.left, "-", e.right, ")")
	return e.left.interpret(cxt) - e.right.interpret(cxt)
}

type Plus struct {
	left  Expression
	right Expression
}

func (e *Plus) interpret(cxt Context) int {
	fmt.Println("(", e.left, "+", e.right, ")")
	return e.left.interpret(cxt) + e.right.interpret(cxt)
}

type Context int
