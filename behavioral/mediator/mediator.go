package main

import "fmt"

// 中介者模式（Mediator Pattern）又叫作调解者模式或调停者模式。 用一个中介对象封装一系列对象交互， 中介者使各对象不需要显式地相互作用，
// 从而使其耦合松散， 而且可以独立地改变它们之间的交互， 属于行为型设计模式。

// 中介者模式主要适用于以下应用场景。
// 1、系统中对象之间存在复杂的引用关系，产生的相互依赖关系结构混乱且难以理解。
// 2、交互的公共行为，如果需要改变行为，则可以增加新的中介者类。

// 中介者模式的优点
// 1、减少类间依赖，将多对多依赖转化成一对多，降低了类间耦合。
// 2、类间各司其职，符合迪米特法则。

// 中介者模式的缺点
// 1、中介者模式将原本多个对象直接的相互依赖变成了中介者和多个组件类的依赖关系。
// 2、当组件类越多时，中介者就会越臃肿，变得复杂且难以维护。

// 中介者模式的一个绝佳例子就是火车站交通系统。 两列火车互相之间从来不会就站台的空闲状态进行通信。
//​ station­Manager车站经理可充当中介者， 让平台仅可由一列入场火车使用， 而将其他火车放入队列中等待。 离场火车会向车站发送通知， 便于队列中的下一列火车进站。

// 火车接口
type Train interface {
	// 到达
	arrive()
	// 离开
	depart()
	// 许可进站
	permitArrival()
}

// 客车
type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}

// 货车
type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

// 中介者接口
type Mediator interface {
	// 是否可以进站
	canArrive(Train) bool
	// 通知火车进站
	notifyAboutDeparture()
}

// 车站管理员
type StationManager struct {
	// 站台是否空闲
	isPlatformFree bool
	// 等待进站火车列表
	trainQueue []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

func main() {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
