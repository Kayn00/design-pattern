package main

import "fmt"

// 装饰模式是一种结构型设计模式， 允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。
// 由于目标对象和装饰器遵循同一接口， 因此你可用装饰来对对象进行无限次的封装。 结果对象将获得所有封装器叠加而来的行为。

// 结构
// 1.抽象构件（Component）角色。定义一个抽象接口以规范准备接收附加责任的对象。
// 2.具体构件（ConcreteComponent）角色。被装饰的真实对象，实现抽象构件，通过装饰角色为其添加一些职责。
// 3.抽象装饰（Decorator）角色(如果有扩展需要)。继承抽象构件，并包含具体构件的实例，可以通过其子类扩展具体构件的功能。
// 4.具体装饰（ConcreteDecorator）角色。实现抽象装饰的相关方法，并给具体构件对象添加附加的责任。

// 优点
// 1.通过使用不用装饰类及这些装饰类的排列组合，可以实现不同效果。
// 2.装饰器模式完全遵守开闭原则。
// 3.不改变原对象下，动态给对象增加行为，比继承更灵活。

// 缺点
// 1.装饰器模式会增加许多子类，过度使用会增加程序得复杂性。

// 适用场景
// 1.给现有类增加职责，又不通过增加子类进行扩充。例如，该类被隐藏或者该类是终极类或者采用继承方式会产生大量的子类。
// 2.对现有基本功能排列组合产生许多功能，通过继承不好实现，通过装饰器很好实现。
// 3.当对象的功能要求可以动态地添加，也可以再动态地撤销时。

/*
业务场景：
- 定义披萨，披萨面饼+不同馅料装饰 -> 不同价格
*/

// 1.抽象组件，接口，具体组件和装饰器都要实现该接口
type iPizza interface {
	getPrice() float64
}

// 2.具体组件，声明实际行为
type pizzaBase struct {
	price float64
}

func (r *pizzaBase) getPrice() float64 {
	return r.price
}

// 3.不同装饰器，实际是加不同馅料实现对pizza饼的装饰
// 培根
type BaconFilling struct {
	pizza iPizza
}

func (r *BaconFilling) getPrice() float64 {
	return r.pizza.getPrice() + 5
}

// 奶酪
type CheeseFilling struct {
	pizza iPizza
}

func (r *CheeseFilling) getPrice() float64 {
	return r.pizza.getPrice() + 6
}

// client
func main() {
	pizza := &pizzaBase{3.0}

	// add cheese 3+6
	pizzaWithCheese := &CheeseFilling{pizza: pizza}
	fmt.Println("pizzaWithCheese price is : ", pizzaWithCheese.getPrice())

	// add bacon
	pizzaWithBacon := &BaconFilling{pizza: pizzaWithCheese}

	// calculate total price, 3+5+6
	fmt.Printf("The price of pizza with bacon and cheese is: $ %.2f\n", pizzaWithBacon.getPrice())
}
