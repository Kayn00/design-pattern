package main

import "fmt"

// 工厂模式用来创建一组相关或者相互依赖的对象，与工厂方法模式的区别就在于，工厂方法模式针对的是一个产品等级结构；而抽象工厂模式则是针对的多个产品等级结构，
// 我们可以将一种产品等级想象为一个产品族，所谓的产品族，是指位于不同产品等级结构中功能相关联的产品组成的家族。

// 抽象工厂模式的优点
// 当需要产品族时，抽象工厂可以保证客户端始终只使用同一个产品的产品族。
// 抽象工厂增强了程序的可扩展性，对于新产品族的增加，只需实现一个新的具体工厂即可，不需要对已有代码进行修改，符合开闭原则。

// 抽象工厂模式的缺点
// 规定了所有可能被创建的产品集合，产品族中扩展新的产品困难，需要修改抽象工厂的接口。
// 增加了系统的抽象性和理解难度。

// 工厂接口
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}

// TV产品接口
type ITelevision interface {
	Watch()
}

// 空调产品接口
type IAirConditioner interface {
	SetTemperature(int)
}

// 具体工厂 华为工厂
type HuaWeiFactory struct{}

func (hf *HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTV{}
}
func (hf *HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{}
}

// 具体产品 华为TV
type HuaWeiTV struct{}

func (ht *HuaWeiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

// 具体产品 华为空调
type HuaWeiAirConditioner struct{}

func (ha *HuaWeiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("HuaWei AirConditioner set temperature to %d ℃\n", temp)
}

// 具体工厂 小米工厂
type MiFactory struct{}

func (mf *MiFactory) CreateTelevision() ITelevision {
	return &MiTV{}
}
func (mf *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}

// 具体产品 小米TV
type MiTV struct{}

func (mt *MiTV) Watch() {
	fmt.Println("Watch Mi TV")
}

// 具体产品 小米空调
type MiAirConditioner struct{}

func (ma *MiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("Mi AirConditioner set temperature to %d ℃\n", temp)
}

func main() {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaWeiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(25)

	factory = &MiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(26)
}
