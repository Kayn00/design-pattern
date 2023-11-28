package main

import "fmt"

/**
 * 业务场景: 农民种菜,村里安排农民收菜
 * 产品:韭菜
 * 抽象建造者:人
 * 具体建造者:农民伯伯
 * 指挥者:村
 */
//韭菜
type Leek struct {
}

// 韭菜建造者抽象-农民
type PeasantInter interface {
	Water() error         // 浇水
	Fertilization() error // 施肥
	Reap() *Leek          // 收割
}

// 建造者实体 - 农民
type Peasant struct {
}

func (p *Peasant) Water() error {
	fmt.Println("给韭菜浇水")
	return nil
}
func (p *Peasant) Fertilization() error {
	fmt.Println("给韭菜施肥")
	return nil
}
func (p *Peasant) Reap() *Leek {
	fmt.Println("收割韭菜")
	return &Leek{}
}

// 指挥者 - 农村
type Village struct {
	peasant Peasant
}

func (v *Village) SetF(p Peasant) {
	v.peasant = p
}

func (v *Village) getProduct() *Leek {
	return v.peasant.Reap()
}

// 开始指挥建造者割韭菜
func (v *Village) Gen() *Leek {
	v.peasant.Water()
	v.peasant.Fertilization()
	return v.peasant.Reap()
}

func main() {
	// 创建一个指挥者
	village := new(Village)
	// 创建建造者
	peasant := Peasant{}
	village.SetF(peasant)
	leek := village.Gen()
	fmt.Println(leek)
}
