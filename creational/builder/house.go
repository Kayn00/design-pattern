package main

import "fmt"

// 建造者（Builder）模式：指将一个复杂对象的构造与它的表示分离，使同样的构建过程可以创建不同的对象，这样的设计模式被称为建造者模式。
// 它是将一个复杂的对象分解为多个简单的对象，然后一步一步构建而成。它将变与不变相分离，即产品的组成部分是不变的，但每一部分是可以灵活选择的。

// 建造者（Builder）模式的主要角色如下。
// 抽象建造者（Abstract Builder）：它是一个包含创建产品各个子部件的抽象方法的接口，通常还包含一个返回复杂产品的方法 getResult()。
// 具体建造者(Concrete Builder）：实现 Builder 接口，完成复杂产品的各个部件的具体创建方法。
// 指挥者（Director）：它调用建造者对象中的部件构造与装配方法完成复杂对象的创建，在指挥者中不涉及具体产品的信息。
// 产品角色（Product）：它是包含多个组成部件的复杂对象，由具体建造者来创建其各个零部件。

// 建造者模式的使用场景：
// 1、建造者模式通常适用于有多个构造器参数或者需要较多构建步骤的场景。使用建造者模式可以精简构造器参数的数量，让构建过程更有条理。
// 2、可以为同一个产品提供两个不同的实现。比如，在下面的代码中，为house类型创建了两个不同的实现：通过iglooBuilder构建的igloo（冰屋），以及通过cabinBuilder构建的cabin（木屋）
// 3、可以应用于构建过程不允许被中断的场景。仍然参考刚才的代码，house类型的对象要么彻底完成，要么压根没有创建，不会存在中间状态，这是因为struct director封装了相应的过程，中间状态仅存在于ConcreteBuilder中。

// 抽象建造者
type IBuilder interface {
	setWindowType() // 窗
	setDoorType()   // 门
	setNumFloor()   // 地板
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// 具体建造者 木屋
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// 具体建造者 冰屋
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// 指挥者
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// 产品角色
type House struct {
	windowType string
	doorType   string
	floor      int
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
