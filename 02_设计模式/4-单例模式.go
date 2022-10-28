package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var initialized uint32 // 标记
var lock sync.Mutex

/*
三个要点：
		一是某个类只能有一个实例；
		二是它必须自行创建这个实例；
		三是它必须自行向整个系统提供这个实例。
*/

/*
	保证一个类永远只能有一个对象
*/

//1、保证这个类非公有化，外界不能通过这个类直接创建一个对象
//   那么这个类就应该变得非公有访问 类名称首字母要小写
type singleton struct{}

//2、但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
//   Golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
var instance *singleton = new(singleton)

/*
3、如果全部为私有化，那么外部模块将永远无法访问到这个类和对象，
所以需要对外提供一个方法来获取这个唯一实例对象
注意：这个方法是否可以定义为singelton的一个成员方法呢？
答案是不能，因为如果为成员方法就必须要先访问对象、再访问函数
但是类和对象目前都已经私有化，外界无法访问，所以这个方法一定是一个全局普通函数
*/
//GetInstance 饿汉式创建
/*
func GetInstance() *singleton {
	return instance
}
*/

/*
//GetInstance 懒汉式创建
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
		return instance
	}

	return instance
}
*/

/*
//GetInstance 线程安全式创建(加锁)
func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
		return instance
	}

	return instance
}
*/

/*
//GetInstance 线程安全式创建(原子类操作)
func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&initialized) == 0 {
		instance = new(singleton)
		//atomic.StoreUint32(&initialized, 1)
		atomic.AddUint32(&initialized, 1)
	}
	return instance
}
*/

//GetInstance sync.Once创建(Once装饰的方法只会被执行一次)
func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
