package main

import "fmt"

// 简单工厂需要:
// 工厂结构体
// 产品接口
// 产品结构体

// 这些对象对应的类都实现了相同的接口
// 当我们想要创建一个对象的时候，调用同一个方法，传入不同的参数就可以返回给我们不同的对象了

// 简单工厂的优点是简单
// 缺点:如果具体产品扩产，就必须修改工厂内部，增加Case，一旦产品过多就会导致简单工厂过于臃肿，为了解决这个问题，才有了下一级别的工厂模式--工厂方法。

type Factory struct {
}

func (f Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	default:
		return nil
	}
}

type Product interface {
	create()
}

// 产品1，实现产品接口
type Product1 struct {
}

func (p1 Product1) create() {
	fmt.Println("this is product 1")
}

// 产品2，实现产品接口
type Product2 struct {
}

func (p1 Product2) create() {
	fmt.Println("this is product 2")
}

func main() {
	// 创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
	factory := new(Factory)

	// 在工厂类中传入不同的参数，获取不同的实例
	p1 := factory.Generate("product1")
	p1.create() // output:   this is product 1

	p2 := factory.Generate("product2")
	p2.create() // output:   this is product 2
}
