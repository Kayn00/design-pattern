package main

import "fmt"

// 备忘录模式是一种行为型设计模式。这种模式允许我们保存对象在某些关键节点时的必要信息，以便于在适当的时候可以将之恢复到之前的状态。通常它可以用来帮助设计撤销/恢复操作。

// 下面是备忘录设计模式的主要角色：
// 1、Originator（发起者）：Originator是当前的基础对象，它会将自己的状态保存进备忘录
// 2、Memento（备忘录） ： 存储着Originator的状态的对象
// 3、Caretaker（管理人）：Caretaker是保存着多条备忘录的对象，并维护着备忘录的索引，在需要的时候会返回相应的备忘录

// Originator有两个方法： savememento() 和 restorememento()。
// 1、createMemento(): Originator通过这个方法将其状态保存进一个备忘录对象
// 2、restorememento()： 这个方法将备忘录对象作为输入信息。Originator会通过传进来的备忘录信息执行重建。

// 优点:
// 1、备忘录模式仅做数据备忘，不论该数据是否正确。
// 2、设计模式最大的优点就是解耦，各司其职，发起人只需要提供备忘数据，不需要对其进行管理
//
// 缺点
// 1、实际应用中，备忘录模式大多是多状态的，如果进行大量备忘的话，会占用大量内存，当然，如果持久化在磁盘中的话，会减少内存占用，但会增加IO操作，这就需要开发者根据实际业务情况进行取舍了。

// 备忘录
type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

// 发起者
type originator struct {
	state string
}

func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

func (e *originator) restoreMemento(m *memento) {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}

// 管理人
type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}

func main() {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
