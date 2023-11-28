package main

import "fmt"

// 原型模式也是一种创建型模式，它可以帮助我们优雅地创建对象的拷贝。
// 在这种设计模式里面，将克隆某个对象的职责交给了要被克隆的这个对象。被克隆的对象需要提供一个clone()方法。通过这个方法可以返回该对象的拷贝。

// 原型模式的使用场景：
// 创建新对象的操作比较耗资源（如数据库操作）或代价比较高时。比较起从头创建新对象，克隆对象明显更加可取
// 要被克隆的对象创建起来比较复杂时：比如对象克隆的过程中存在深度拷贝或分层拷贝时；又比如要被克隆的对象存在无法被直接访问到的私有成员时。

// 使用原型模式，需要了解深拷贝和浅拷贝。
// 浅拷贝是指被复制对象的所有变量都含有与原来的对象相同的值，而所有的对其他对象的引用都仍然指向原来的对象。
// 深拷贝把引用对象的变量指向复制过的新对象，而不是原有的被引用的对象。

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}
	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")
	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
