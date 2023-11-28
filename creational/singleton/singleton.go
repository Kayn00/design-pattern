package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

// Q:为什么要使用单例模式?
// A:1、处理资源访问冲突;2、表示全局唯一类;
// https://gitlab.gf.com.cn/gfstore/finance-backend/blob/develop/models/pcenter_asset.go#L754
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{} // 懒汉模式
		} else {
			fmt.Println("Single instance already created.11111")
		}
	} else {
		fmt.Println("Single instance already created.22222")
	}

	return singleInstance
}

// 饿汉式
// 包在被导入时自动初始化
//var singleInstance = &single{}
//func GetInstance() *single {
//	return singleInstance
//}

var once sync.Once

//func GetInstance() *single {
//	once.Do(func() {
//		fmt.Println("Creating single instance now.")
//		singleInstance = &single{} // 如果singleInstance设置为nil就不会在初始化
//	})
//	return singleInstance
//}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	//time.Sleep(3 * time.Second)
	//for i := 0; i < 30; i++ {
	//	go GetInstance()
	//}
	//singleInstance = nil
	//GetInstance()

	//for i := 0; i < 30; i++ {
	//	go GetInstance()
	//}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

//var initialized uint32
//
//func GetInstance() *single {
//	// 一次判断即可返回
//	if atomic.LoadUint32(&initialized) == 1 {
//		fmt.Println("Single instance already created.33333")
//		return singleInstance
//	}
//	lock.Lock()
//	defer lock.Unlock()
//	if initialized == 0 {
//		fmt.Println("Creating single instance now.")
//		singleInstance = &single{}
//		atomic.StoreUint32(&initialized, 1) // 原子装载
//	} else {
//		fmt.Println("Single instance already created.44444")
//	}
//	return singleInstance
//}
