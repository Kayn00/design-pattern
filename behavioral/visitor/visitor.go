package main

import "fmt"

// https://juejin.cn/post/6844903728101720071
// 访问者模式（Visitor Pattern）是一种行为型设计模式，它允许你在不修改现有代码的情况下向现有对象结构添加新的行为。
// 该模式建立在两个核心组件上：访问者和元素。访问者是一个能够访问所有元素的对象，而元素则是需要接受访问者的对象。在这种模式下，访问者可以在不改变元素本身的情况下对其进行操作。

// 角色组成
// 1、抽象访问者
// 2、访问者
// 3、抽象元素类
// 4、元素类
// 5、结构容器: (非必须) 保存元素列表，可以放置访问者

// 大概的流程就是:
// 1、从结构容器中取出元素
// 2、创建一个访问者
// 3、将访问者载入传入的元素（即让访问者访问元素）
// 4、获取输出

// 定义访问者接口
type IVisitor interface {
	Visit() // 访问者的访问方法
}

type ProductionVisitor struct {
}

func (v ProductionVisitor) Visit() {
	fmt.Println("这是生产环境")
}

type TestingVisitor struct {
}

func (t TestingVisitor) Visit() {
	fmt.Println("这是测试环境")
}

// 定义元素接口
type IElement interface {
	Accept(visitor IVisitor)
}

// 实现元素接口
type Element struct {
}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

// 修改 Print() 方法
type EnvExample struct {
	Element
}

func (e EnvExample) Print(visitor IVisitor) {
	e.Element.Accept(visitor)
}

func main() {
	// 创建一个元素
	e := new(Element)
	e.Accept(new(ProductionVisitor)) // output: 这是生产环境
	e.Accept(new(TestingVisitor))    // output: 这是测试环境

	m := new(EnvExample)
	m.Print(new(ProductionVisitor))
	m.Print(new(TestingVisitor))
}
