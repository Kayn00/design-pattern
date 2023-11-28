package main

import "fmt"

// 简单工厂只有唯一的一个工厂控制着 所有产品的实例化，而 工厂方法 中包括一个工厂接口，我们可以动态的实现多种工厂，达到扩展的目的
// 因此，当增加一个产品时，只需增加一个相应的工厂类的子类, 以解决简单工厂生产太多产品时导致其内部代码臃肿（switch … case分支过多）的问题。

// 简单工厂需要:
// 工厂结构体
// 产品接口
// 产品结构体

// 工厂方法需要:
// 工厂接口
// 工厂结构体
// 产品接口
// 产品结构体

// 工厂方法模式的优点
// 灵活性增强，对于新产品的创建，只需多写一个相应的工厂类。
// 典型的解耦框架。高层模块只需要知道产品的抽象类，无须关心其他实现类，满足迪米特法则、依赖倒置原则和里氏替换原则。

// 工厂方法模式的缺点
// 类的个数容易过多，增加复杂度。
// 增加了系统的抽象性和理解难度。
// 只能生产一种产品，此弊端可使用抽象工厂模式解决。
// 无论是简单工厂还是工厂方法都只能生产一种产品，如果工厂需要创建生态里的多个产品，就需要更进一步，使用第三级的工厂模式--抽象工厂。

// OperatorFactory 工厂接口，由具体工厂类来实现
type OperatorFactory interface {
	Create() MathOperator
}

// MathOperator 实际产品实现的接口--表示数学运算器应该有哪些行为
type MathOperator interface {
	SetOperandA(int)
	SetOperandB(int)
	ComputeResult() int
}

// BaseOperator 是所有 Operator 的基类
// 封装公用方法，因为Go不支持继承，具体Operator类
// 只能组合它来实现类似继承的行为表现。
type BaseOperator struct {
	operandA, operandB int
}

func (o *BaseOperator) SetOperandA(operand int) {
	o.operandA = operand
}

func (o *BaseOperator) SetOperandB(operand int) {
	o.operandB = operand
}

//PlusOperatorFactory 是 PlusOperator 加法运算器的工厂类
type PlusOperatorFactory struct{}

func (pf *PlusOperatorFactory) Create() MathOperator {
	return &PlusOperator{
		BaseOperator: &BaseOperator{},
	}
}

//PlusOperator 实际的产品类--加法运算器
type PlusOperator struct {
	*BaseOperator
}

//ComputeResult 计算并获取结果
func (p *PlusOperator) ComputeResult() int {
	return p.operandA + p.operandB
}

// MultiOperatorFactory 是乘法运算器产品的工厂
type MultiOperatorFactory struct{}

func (mf *MultiOperatorFactory) Create() MathOperator {
	return &MultiOperator{
		BaseOperator: &BaseOperator{},
	}
}

// MultiOperator 实际的产品类--乘法运算器
type MultiOperator struct {
	*BaseOperator
}

func (m *MultiOperator) ComputeResult() int {
	return m.operandA * m.operandB
}

// 测试运行
func main() {
	var factory OperatorFactory
	var mathOp MathOperator
	factory = &PlusOperatorFactory{}
	mathOp = factory.Create()
	mathOp.SetOperandB(3)
	mathOp.SetOperandA(2)
	fmt.Printf("Plus operation reuslt: %d\n", mathOp.ComputeResult())

	factory = &MultiOperatorFactory{}
	mathOp = factory.Create()
	mathOp.SetOperandB(3)
	mathOp.SetOperandA(2)
	fmt.Printf("Multiple operation reuslt: %d\n", mathOp.ComputeResult())
}
