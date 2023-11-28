package main

import (
	"fmt"
	"strconv"
)

// https://juejin.cn/post/7242596585330098232
// 责任链模式是一种行为型的设计模式，也叫职责链，定义：在这样的场景中，请求需要经过一系列的处理者，通过列表或者链表将处理者组合起来，
// 请求依次被处理者处理，如果需要中断，也可以及时退出处理。上述的这种的应用场景，典型的就是在 web 应用中的 中间件或者拦截器，比如 Gin框架 的中间件。

// 结构
// 1.抽象处理者（Handler）角色。定义一个处理请求的接口，包含抽象处理方法和一个后继连接。
// 2.具体处理者（Concrete Handler）角色。实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
// 3.客户类（Client）角色。创建处理链，并向链头的具体处理者对象提交请求，它不关心处理细节和请求的传递过程。

// 适用场景
// 1.多个对象可以处理一个请求，但具体由哪个对象处理该请求在运行时自动确定。
// 2.可动态指定一组对象处理请求，或添加新的处理者。
// 3.需要在不明确指定请求处理者的情况下，向多个处理者中的一个提交请求。

// 优点
// 1.降低了对象之间的耦合度。该模式使得一个对象无须知道到底是哪一个对象处理其请求以及链的结构，发送者和接收者也无须拥有对方的明确信息。
// 2.增强了系统的可扩展性。可以根据需要增加新的请求处理类，满足开闭原则。
// 3.增强了给对象指派职责的灵活性。当工作流程发生变化，可以动态地改变链内的成员或者调动它们的次序，也可动态地新增或者删除责任。
// 4.责任链简化了对象之间的连接。每个对象只需保持一个指向其后继者的引用，不需保持其他所有处理者的引用，这避免了使用众多的 if 或者 if···else 语句。
// 5.责任分担，明确各类的责任范围，符合类的单一职责原则。每个类只需要处理自己该处理的工作，不该处理的传递给下一个对象完成，明确各类的责任范围，符合类的单一职责原则。
//
// 缺点
// 1.不能保证每个请求一定被处理。由于一个请求没有明确的接收者，所以不能保证它一定会被处理，该请求可能一直传到链的末端都得不到处理。
// 2.对比较长的职责链，请求的处理可能涉及多个处理对象，系统性能将受到一定影响。
// 3.职责链建立的合理性要靠客户端来保证，增加了客户端的复杂性。会由于职责链的错误设置而导致系统出错，如可能会造成循环调用。

// chain of responsibility mode
// 1.abstract handler interface
type iHandler interface {
	Handler(handlerID int) string
}

// 2.concrete handler
type handler struct {
	name     string
	next     iHandler
	reqLevel int
}

// new
func NewHandler(name string, next iHandler, reqLevel int) *handler {
	return &handler{
		name:     name,
		next:     next,
		reqLevel: reqLevel,
	}
}

func (r *handler) Handler(reqLevel int) string {
	if r.reqLevel == reqLevel {
		return r.name + " handled " + strconv.Itoa(reqLevel)
	}

	if r.next == nil {
		return ""
	}

	return r.next.Handler(reqLevel)
}

// 3.client
func main() {
	zh := NewHandler("zhangSan", nil, 1)
	li := NewHandler("liSi", zh, 2)
	w := NewHandler("wangWu", li, 3)

	r := w.Handler(3) // 通过next依次向下传递，直到得到处理或者到最后
	fmt.Println(r)
}
