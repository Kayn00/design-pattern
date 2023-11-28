package main

import "fmt"

// 桥接是一种结构型设计模式，可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用。

// 1、抽象部分（Abstraction）提供高层控制逻辑，依赖于完成底层实际工作的实现对象。
// 2、实现部分（Implementation）为所有具体实现声明通用接口。抽象部分仅能通过在这里声明的方法与实现对象交互。
// 抽象部分可以列出和实现部分一样的方法，但是抽象部分通常声明一些复杂行为，这些行为依赖于多种由实现部分声明的原语操作。
// 3、具体实现（Concrete Implementations）中包括特定于平台的代码。
// 4、精确抽象（Refined Abstraction）提供控制逻辑的变体。与其父类一样，它们通过通用实现接口与不同的实现进行交互。
//
// 通常情况下，客户端（Client）仅关心如何与抽象部分合作。但是，客户端需要将抽象对象与一个实现对象连接起来。

// 1.优点
// 你可以创建与平台无关的类和程序。
// 客户端代码仅与高层抽象部分进行互动，不会接触到平台的详细信息。
// *开闭原则:*你可以新增抽象部分和实现部分，且它们之间不会相互影响。
// *单一职责原则:*抽象部分专注于处理高层逻辑， 实现部分处理平台细节。
//
// 2.缺点
// 对高内聚的类使用该模式可能会让代码更加复杂。

// 使用场景
// 如果你想要拆分或重组一个具有多重功能的庞杂类（例如能与多个数据库服务器进行交互的类），可以使用桥接模式。
// 如果你希望在几个独立维度上扩展一个类，可使用该模式。
// 如果你需要在运行时切换不同实现方法，可使用桥接模式。

// 抽象部分 高层控制
type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// 实现部分（底层实际工作）
// 打印机 打印文件
type Printer interface {
	PrintFile()
}

// 爱普生
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// 惠普
type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
