package main

import "fmt"

// 组合模式

type Component interface {
	Add(Component)
	Remove(Component)
	Print()
}

type component struct {
	name string
	desc string
}

func NewComponent(name, desc string) component {
	return component{
		name: name,
		desc: desc,
	}
}

type University struct {
	component
	Colleges map[Component]bool
}

func (u *University) Add(c Component) {
	if u.Colleges == nil {
		u.Colleges = make(map[Component]bool)
	}
	u.Colleges[c] = true
}

func (u *University) Remove(c Component) {
	delete(u.Colleges, c)
}

func (u *University) Print() {
	fmt.Println("--------------" + u.name + "--------------")
	for c := range u.Colleges {
		c.Print()
	}
}

type College struct {
	component
	Departments map[Component]bool
}

func (o *College) Add(c Component) {
	if o.Departments == nil {
		o.Departments = make(map[Component]bool)
	}
	o.Departments[c] = true
}

func (o *College) Remove(c Component) {
	delete(o.Departments, c)
}

func (o *College) Print() {
	fmt.Println("--------------" + o.name + "--------------")
	for d := range o.Departments {
		d.Print()
	}
}

type Department struct {
	component
}

func (d *Department) Add(c Component) {
}

func (d *Department) Remove(c Component) {
}

func (d *Department) Print() {
	fmt.Println(d.name)
}
