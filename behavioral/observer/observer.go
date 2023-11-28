package main

import "fmt"

// https://juejin.cn/post/6844903729703944205
// 角色
// 1、抽象观察者
// 2、具体观察者
// 3、抽象被观察者
// 4、具体被观察者

// 抽象观察者
type IObserver interface {
	Notify() // 当被观察对象有更改的时候，出发观察者的Notify() 方法
}

// 抽象被观察者
type ISubject interface {
	AddObservers(observers ...IObserver) // 添加观察者
	NotifyObservers()                    // 通知观察者
}

// 具体观察者
type Observer2 struct {
}

func (o *Observer2) Notify() {
	fmt.Println("已经触发了观察者2")
}

// 具体观察者
type Observer1 struct {
}

func (o *Observer1) Notify() {
	fmt.Println("已经触发了观察者1")
}

// 具体被观察者
type Subject2 struct {
	observers []IObserver
}

func (s *Subject2) AddObservers(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject2) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify() // 触发观察者
	}
}

type Subject1 struct {
	observers []IObserver
}

func (s *Subject1) AddObservers(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject1) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify() // 触发观察者
	}
}

func main() {
	// 创建被观察者
	s2 := new(Subject2)
	s1 := new(Subject1)
	// 创建观察者
	o2 := new(Observer2)
	o1 := new(Observer1)
	// 为主题添加观察者
	s2.AddObservers(o2, o1)
	s1.AddObservers(o1, o2)

	// 这里的被观察者要做各种更改...

	// 更改完毕，触发观察者
	s2.NotifyObservers() // output: 已经触发了观察者
	fmt.Println()
	s1.NotifyObservers()
}
