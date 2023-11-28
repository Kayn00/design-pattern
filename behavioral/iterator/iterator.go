package main

import "fmt"

// 角色
// 1、抽象聚合类: 定义一个抽象的容器
// 2、具体聚合类: 实现上面的抽象类，作为一个容器，用来存放元素，等待迭代
// 3、抽象迭代器: 迭代器接口，每个容器下都有一个该迭代器接口的具体实现
// 4、具体迭代器: 根据不同的容器，需要定义不同的具体迭代器，定义了游标移动的具体实现

// 容器接口
type IAggregate interface {
	Iterator() IIterator
}

// 具体容器
type Aggregate struct {
	container []int // 容器中装载 int 型容器
}

// 创建一个迭代器，并让迭代器中的容器指针指向当前对象
func (a *Aggregate) Iterator() IIterator {
	i := new(Iterator)
	i.aggregate = a
	return i
}

// 迭代器接口
type IIterator interface {
	HasNext() bool
	Current() int
	Next() bool
}

type Iterator struct {
	cursor    int        // 当前游标
	aggregate *Aggregate // 对应的容器指针
}

// 判断是否迭代到最后，如果没有，则返回true
func (i *Iterator) HasNext() bool {
	if i.cursor+1 < len(i.aggregate.container) {
		return true
	}
	return false
}

// 获取当前迭代元素（从容器中取出当前游标对应的元素）
func (i *Iterator) Current() int {
	return i.aggregate.container[i.cursor]
}

// 将游标指向下一个元素
func (i *Iterator) Next() bool {
	if i.cursor < len(i.aggregate.container) {
		i.cursor++
		return true
	}
	return false
}

func main() {
	// 创建容器，并放入初始化数据
	c := &Aggregate{container: []int{1, 2, 3, 4}}
	// 获取迭代器
	iterator := c.Iterator()
	for {
		// 打印当前数据
		fmt.Println(iterator.Current())
		// 如果有下一个元素，则将游标移动到下一个元素
		// 否则跳出循环，迭代结束
		if iterator.HasNext() {
			iterator.Next()
		} else {
			break
		}
	}
}
