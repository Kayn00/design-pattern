package main

import "fmt"

// 主题 发布者
type ISubject interface {
	AddObservers(observers ...IObserver) // 添加观察者
	NotifyObservers()                    // 通知观察者
}

// 具体主题
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

// 具体主题
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

// 抽象观察者 订阅者
type IObserver interface {
	Notify() // 当被观察对象有更改的时候，出发观察者的Notify() 方法
}

// 具体订阅者
type Observer2 struct {
}

func (o *Observer2) Notify() {
	fmt.Println("已经触发了观察者2")
}

// 具体订阅者
type Observer1 struct {
}

func (o *Observer1) Notify() {
	fmt.Println("已经触发了观察者1")
}



func main() {
	// 创建主题
	s2 := new(Subject2)
	s1 := new(Subject1)
	// 创建订阅者
	o2 := new(Observer2)
	o1 := new(Observer1)
	// 为主题添加订阅者
	s2.AddObservers(o2, o1)
	s1.AddObservers(o1, o2)

	// 这里的主题要做各种更改...

	// 更改完毕，触发订阅者
	s2.NotifyObservers() // output: 已经触发了订阅者
	fmt.Println()
	s1.NotifyObservers()
}
