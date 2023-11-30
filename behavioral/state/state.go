package main

// 状态模式是一种行为设计模式， 让你能在一个对象的内部状态变化时改变其行为， 使其看上去就像改变了自身所属的类一样。
// 状态模式包含以下主要角色。
// 1、上下文（Context）角色：它定义了客户端需要的接口，内部维护一个当前状态，并负责具体状态的切换。
// 2、抽象状态（State）角色：定义一个接口，用以封装环境对象中的特定状态所对应的行为，可以有一个或多个行为。
// 3、具体状态（Concrete State）角色：实现抽象状态所对应的行为，并且在需要的情况下进行状态切换。

import "fmt"

type Mario struct {
	score  int64
	status MarioStatus
}

type MarioStatus interface {
	Name()
	ObtainMushroom()
	ObtainCape()
	MeetMonster()
	SetMario(mario *Mario)
}

/**
 * @Author: Jason Pang
 * @Description: 小马里奥
 */
type SmallMarioStatus struct {
	mario *Mario
}

/**
 * @Author: Jason Pang
 * @Description: 设置马里奥
 * @receiver s
 * @param mario
 */
func (s *SmallMarioStatus) SetMario(mario *Mario) {
	s.mario = mario
}

func (s *SmallMarioStatus) Name() {
	fmt.Println("小马里奥")
}

/**
 * @Author: Jason Pang
 * @Description: 获得蘑菇变为超级马里奥
 * @receiver s
 */
func (s *SmallMarioStatus) ObtainMushroom() {
	s.mario.status = &SuperMarioStatus{
		mario: s.mario,
	}
	s.mario.score += 100
}

/**
 * @Author: Jason Pang
 * @Description: 获得斗篷变为斗篷马里奥
 * @receiver s
 */
func (s *SmallMarioStatus) ObtainCape() {
	s.mario.status = &CapeMarioStatus{
		mario: s.mario,
	}
	s.mario.score += 200
}

/**
 * @Author: Jason Pang
 * @Description: 遇到怪兽减100
 * @receiver s
 */
func (s *SmallMarioStatus) MeetMonster() {
	s.mario.score -= 100
}

/**
 * @Author: Jason Pang
 * @Description: 超级马里奥
 */

type SuperMarioStatus struct {
	mario *Mario
}

/**
 * @Author: Jason Pang
 * @Description: 设置马里奥
 * @receiver s
 * @param mario
 */
func (s *SuperMarioStatus) SetMario(mario *Mario) {
	s.mario = mario
}

func (s *SuperMarioStatus) Name() {
	fmt.Println("超级马里奥")
}

/**
 * @Author: Jason Pang
 * @Description: 获得蘑菇无变化
 * @receiver s
 */
func (s *SuperMarioStatus) ObtainMushroom() {

}

/**
 * @Author: Jason Pang
 * @Description:获得斗篷变为斗篷马里奥
 * @receiver s
 */
func (s *SuperMarioStatus) ObtainCape() {
	s.mario.status = &CapeMarioStatus{
		mario: s.mario,
	}
	s.mario.score += 200
}

/**
 * @Author: Jason Pang
 * @Description: 遇到怪兽变为小马里奥
 * @receiver s
 */
func (s *SuperMarioStatus) MeetMonster() {
	s.mario.status = &SmallMarioStatus{
		mario: s.mario,
	}
	s.mario.score -= 200
}

/**
 * @Author: Jason Pang
 * @Description: 斗篷马里奥
 */
type CapeMarioStatus struct {
	mario *Mario
}

/**
 * @Author: Jason Pang
 * @Description: 设置马里奥
 * @receiver s
 * @param mario
 */
func (c *CapeMarioStatus) SetMario(mario *Mario) {
	c.mario = mario
}

func (c *CapeMarioStatus) Name() {
	fmt.Println("斗篷马里奥")
}

/**
 * @Author: Jason Pang
 * @Description:获得蘑菇无变化
 * @receiver c
 */
func (c *CapeMarioStatus) ObtainMushroom() {

}

/**
 * @Author: Jason Pang
 * @Description: 获得斗篷无变化
 * @receiver c
 */
func (c *CapeMarioStatus) ObtainCape() {

}

/**
 * @Author: Jason Pang
 * @Description: 遇到怪兽变为小马里奥
 * @receiver c
 */
func (c *CapeMarioStatus) MeetMonster() {
	c.mario.status = &SmallMarioStatus{
		mario: c.mario,
	}
	c.mario.score -= 200
}
func main() {
	mario := Mario{
		status: &SmallMarioStatus{},
		score:  0,
	}
	mario.status.SetMario(&mario)

	mario.status.Name()
	fmt.Println("-------------------获得蘑菇\n")
	mario.status.ObtainMushroom()

	mario.status.Name()
	fmt.Println("-------------------获得斗篷\n")
	mario.status.ObtainCape()

	mario.status.Name()
	fmt.Println("-------------------遇到怪兽\n")
	mario.status.MeetMonster()

	mario.status.Name()
}
