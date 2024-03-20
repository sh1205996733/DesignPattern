package main

// 迭代器模式：它允许通过对集合对象进行迭代来访问集合中的元素，而不需要暴露集合内部的结构。

// 1、创建抽象迭代器和想要迭代遍历的对象
type Iterator interface {
	hasNext() bool
	next() *Student
}

type Student struct {
}

// 2、创建具体迭代器角色
type StudentIterator struct {
	stus []*Student
}

func (i *StudentIterator) hasNext() bool {
	return len(i.stus) > 0
}

func (i *StudentIterator) next() *Student {
	stu := i.stus[0]
	i.stus = i.stus[1:]
	return stu
}

// 3、创建抽象聚合角色
// 4、创建具体聚合角色
type StudentAggregate struct {
	stus []*Student
}

func (s *StudentAggregate) addStudent(stu ...*Student) {
	s.stus = append(s.stus, stu...)
}

func (s *StudentAggregate) removeStudent() {
	s.stus = s.stus[:len(s.stus)-1]
}
func (s *StudentAggregate) getIterator() Iterator {
	return &StudentIterator{s.stus}
}
