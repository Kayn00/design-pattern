package main

import (
	"fmt"
	"strings"
	"sync"
)

// 享元模式的主要角色有如下。
// 1、抽象享元角色（Flyweight）：是所有的具体享元类的基类，为具体享元规范需要实现的公共接口，非享元的外部状态以参数的形式通过方法传入。
// 2、具体享元（Concrete Flyweight）角色：实现抽象享元角色中所规定的接口。
// 3、非享元（Unsharable Flyweight)角色：是不可以共享的外部状态，它以参数的形式注入具体享元的相关方法中。
// 4、享元工厂（Flyweight Factory）角色：负责创建和管理享元角色。当客户对象请求一个享元对象时，享元工厂检査系统中是否存在符合要求的享元对象，如果存在则提供给客户；如果不存在的话，则创建一个新的享元对象。

// 场景
// 1、某火车票查问, 可依据发站和到站, 查问余票信息
// 2、火车票蕴含根本信息(发站, 到站, 经停站, 登程工夫, 到站工夫…)和剩余票数信息
// 3、根本信息字段较多, 且只跟发站和到站相干, 因而可采纳享元模式进行池化解决
// 4、残余票数信息因为实时变动, 因而由余票服务另外提供

// 设计
// ITicket: 定义车票根本信息接口
// ITicketRemaining: 继承ITicket, 并增加余票数信息
// ITicketService: 定义车票信息服务接口
// ITicketRemainingService: 定义余票信息服务接口. 依据发站和到站, 查问余票信息.
// tMockTicket: 车票信息实体, 实现ITicket接口
// tMockTicketService: 车票信息服务, 通过享元模式池化了车票信息.
// tMockTicketRemaining: 余票信息实体, 实现ITicketRemaining接口
// tMockTicketRemainingService: 余票信息服务, 通过ITicketService获取车票根本信息. 依据发站和到站, 查问余票信息.

var (
	MockTicketService          ITicketService          = newMockTicketService()
	MockTicketRemainingService ITicketRemainingService = newMockTicketRemainingService()
)

// 定义车票根本信息接口
type ITicket interface {
	ID() int
	From() string
	To() string
	LeavingTime() string
	ArrivalTime() string
	InterList() []string
	Price() float64
}

// 继承ITicket, 并增加余票数信息
type ITicketRemaining interface {
	ITicket
	Remaining() int
}

// 定义车票信息服务接口
type ITicketService interface {
	Get(from string, to string) ITicket
	Save(it ITicket)
}

// 定义余票信息服务接口, 依据发站和到站, 查问余票信息.
type ITicketRemainingService interface {
	Get(from string, to string) ITicketRemaining
	Save(id int, num int)
}

// 车票信息实体, 实现ITicket接口
type tMockTicket struct {
	iID          int
	sFrom        string
	sTo          string
	sLeavingTime string
	sArrivalTime string
	mInterList   []string
	fPrice       float64
	iRemaining   int
}

func NewMockTicket(id int, from string, to string, price float64) *tMockTicket {
	return &tMockTicket{
		iID:          id,
		sFrom:        from,
		sTo:          to,
		sLeavingTime: "09:00",
		sArrivalTime: "11:30",
		mInterList:   strings.Split("深圳北,虎门", ","),
		fPrice:       price,
	}
}

func (me *tMockTicket) ID() int {
	return me.iID
}

func (me *tMockTicket) From() string {
	return me.sFrom
}

func (me *tMockTicket) To() string {
	return me.sTo
}

func (me *tMockTicket) LeavingTime() string {
	return me.sLeavingTime
}

func (me *tMockTicket) ArrivalTime() string {
	return me.sArrivalTime
}

func (me *tMockTicket) InterList() []string {
	return me.mInterList
}

func (me *tMockTicket) Price() float64 {
	return me.fPrice
}

// 车票信息服务, 实现ITicketService接口, 通过享元模式池化了车票信息.
type tMockTicketService struct {
	mTickets map[string]ITicket
	mRWMutex *sync.RWMutex
}

func newMockTicketService() *tMockTicketService {
	return &tMockTicketService{
		make(map[string]ITicket, 0),
		new(sync.RWMutex),
	}
}

func (me *tMockTicketService) Get(from string, to string) ITicket {
	k := from + "-" + to

	me.mRWMutex.RLock()
	defer me.mRWMutex.RUnlock()
	it, ok := me.mTickets[k]

	if ok {
		return it

	} else {
		return nil
	}
}

func (me *tMockTicketService) Save(it ITicket) {
	k := it.From() + "-" + it.To()

	me.mRWMutex.Lock()
	defer me.mRWMutex.Unlock()
	me.mTickets[k] = it
}

// 余票信息实体, 实现ITicketRemaining接口
type tMockTicketRemaining struct {
	ITicket
	iRemaining int
}

func newMockTicketRemaining(it ITicket, num int) *tMockTicketRemaining {
	return &tMockTicketRemaining{
		it, num,
	}
}

func (me *tMockTicketRemaining) Remaining() int {
	return me.iRemaining
}

// 余票信息服务, 实现ITicketRemainingService接口. 通过ITicketService获取车票根本信息. 依据发站和到站, 查问余票信息.
type tMockTicketRemainingService struct {
	mRemaining map[int]int
	mRWMutex   *sync.RWMutex
}

func newMockTicketRemainingService() *tMockTicketRemainingService {
	return &tMockTicketRemainingService{
		make(map[int]int, 16),
		new(sync.RWMutex),
	}
}

func (me *tMockTicketRemainingService) Get(from string, to string) ITicketRemaining {
	ticket := MockTicketService.Get(from, to)
	if ticket == nil {
		return nil
	}
	r := newMockTicketRemaining(ticket, 0)

	me.mRWMutex.RLock()
	defer me.mRWMutex.RUnlock()
	num, ok := me.mRemaining[ticket.ID()]

	if ok {
		r.iRemaining = num
	}

	return r
}

func (me *tMockTicketRemainingService) Save(id int, num int) {
	me.mRWMutex.Lock()
	defer me.mRWMutex.Unlock()
	me.mRemaining[id] = num
}

func main() {
	from := "福田"
	to := "广州南"
	ticket := NewMockTicket(1, from, to, 100)
	MockTicketService.Save(ticket)
	MockTicketRemainingService.Save(ticket.ID(), 10)

	remaining := MockTicketRemainingService.Get(from, to)
	fmt.Println(fmt.Sprintf("from=%s, to=%s, price=%v, remaining=%v\n", remaining.From(), remaining.To(), remaining.Price(), remaining.Remaining()))
}
