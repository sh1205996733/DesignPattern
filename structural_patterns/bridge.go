package main

import "fmt"

// 桥接模式
//在实现桥接模式时，有两个重要的概念：抽象部分【FileParser】和实现部分【VideoFileParser 和 MusicFileParser】。
//抽象部分定义了系统中所需功能的通用接口，而实现部分则负责提供这些接口的具体实现。这两个部分之间通过桥接连接器进行交互。
//在桥接模式中，系统必须包含一个抽象类或接口，它定义了所有实施具体功能所需的方法。
//这个抽象基类与实现它的所有子类之间都应该存在一个桥接连接器【FileParser 作为 OperateSystem 的成员变量，聚合到 OperateSystem 里面】。
//这个桥接连接器就是一个抽象类或接口

// 1.定义抽象接口【FileParser，文件解析器】，其中包含文件解析的抽象方法。
type FileParser interface {
	Parser()
}

// 2、说明具体的文件解析器，用于解析不同类型的文件【视频文件解析器和音频文件解析器】
type VideoParser string

func (p *VideoParser) Parser() {
	fmt.Println("VideoParser Successful")
}

type PictureParser string

func (p *PictureParser) Parser() {
	fmt.Println("PictureParser Successful")
}

// 3、定义桥接连接器【OperateSystem】，需要声明文件解析方法 parseFile() 。
type OperateSystem struct {
	parser FileParser
}

func (p *OperateSystem) ParserFile() {
	if p.parser != nil {
		p.parser.Parser()
		fmt.Println("FileParser Successful")
	}
}

// 4、有很多种不同的操作系统，直接继承自 OperateSystem 即可。
type Windows struct {
	OperateSystem
}

func (p *Windows) ParserFile() {
	fmt.Println("Windows FileParser Start")
	p.OperateSystem.ParserFile()
	fmt.Println("Windows FileParser End")
}

type Linux struct {
	OperateSystem
}

func (p *Linux) ParserFile() {
	fmt.Println("Linux FileParser Start")
	p.OperateSystem.ParserFile()
	fmt.Println("Linux FileParser End")
}

type Mac struct {
	OperateSystem
}

func (p *Mac) ParserFile() {
	fmt.Println("Mac FileParser Start")
	p.OperateSystem.ParserFile()
	fmt.Println("Mac FileParser End")
}
