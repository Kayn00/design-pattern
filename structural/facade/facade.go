package main

import (
	"fmt"
)

// 通过提供一个统一的接口，封装了一个或多个复杂的子系统，为客户端提供简化的访问方式。
// 外观模式将复杂子系统的接口和实现细节与客户端代码解耦，使得客户端只需要与外观对象进行交互，而不需要了解底层子系统的复杂性。

// 外观模式包含以下几个角色：
// 1、Facade（外观）：外观对象是客户端与子系统之间的中间层，它提供了一个简化的接口，将客户端的请求委派给子系统处理。外观对象知道哪些子系统类负责处理请求，并将请求分派给它们。
// 2、Subsystem（子系统）：子系统是实现具体功能的一组类或接口。外观对象将客户端的请求转发给适当的子系统类进行处理。
// 3、Client（客户端）：客户端通过外观对象来访问子系统的功能，它只需要与外观对象进行交互，而不需要直接与子系统类打交道。

// 工作原理
// 1、客户端通过实例化外观对象来访问子系统的功能。
// 2、外观对象接收客户端的请求，并根据请求的类型和参数，将请求委派给适当的子系统类处理。
// 3、子系统类接收到请求后，执行相应的功能逻辑。
// 4、外观对象将子系统的处理结果返回给客户端。

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

// apiImpl facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("\n%s\n%s\n", aRet, bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

func main() {
	api := NewAPI()
	ret := api.Test()
	fmt.Println(ret)
}

/**输出
A module running
B module running
**/
