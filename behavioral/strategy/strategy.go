package main

import "fmt"

// 策略模式定义一组算法类，将每个算法分别封装起来，让他们可以相互替换，策略模式可以使得算法独立于客户端，策略模式用来解耦策略的定义，创建，使用。

type Vehicle interface {
	Go()
}

type Car struct {
}

func (r *Car) Go() {
	fmt.Println("use car")
}

type Bicycle struct {
}

func (r *Bicycle) Go() {
	fmt.Println("use Bicycle")
}

type Traveler struct {
	impl Vehicle
}

func (r *Traveler) SetVehicle(i Vehicle) {
	r.impl = i
}

func (r *Traveler) Go() {
	r.impl.Go()
}

func main() {
	traveler := Traveler{}
	traveler.SetVehicle(&Car{})
	traveler.Go()
	traveler.SetVehicle(&Bicycle{})
	traveler.Go()
}
