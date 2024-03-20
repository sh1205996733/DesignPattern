package main

import "fmt"

// 选项模式 首先我们定义一个结构体，初始化这个结构体，然后给结构体赋值。
// 如果开发后期需要对结构体内的参数进行增加或者删除操作，也就需要对应的对初始化的结构体进行修改。
type Options struct {
	str1 string
	str2 string
	int1 int
	int2 int
	int3 int
}

// 传统方式
func Init(str1 string, str2 string, int1 int, int2 int) {
	options := Options{}
	options.str1 = str1
	options.str2 = str2
	options.int1 = int1
	options.int2 = int2
	fmt.Printf("options:%#v\n", options)
}

// 选项模式
type OptionFunc func(*Options)

func WithStr1(str string) OptionFunc {
	return func(o *Options) {
		o.str1 = str
	}
}

func WithStr2(str string) OptionFunc {
	return func(o *Options) {
		o.str2 = str
	}
}

func InitOptions(opt ...OptionFunc) *Options {
	options := &Options{}
	for _, f := range opt {
		f(options)
	}
	return options
}

// 个人总结满足下面条件可以考虑使用选项模式
// 1.参数确实比较复杂，影响调用方使用
// 2.参数确实有比较清晰明确的默认值
// 3.为参数的后续拓展考虑
