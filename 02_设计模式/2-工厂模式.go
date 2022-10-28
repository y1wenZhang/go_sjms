package main

import "fmt"

// 抽象层
type Fruit1 interface {
	Show()
}

type AbstractFactory1 interface {
	CreateFruit() Fruit1
}

// ====基础模块
type Apple1 struct {
}

func (a *Apple1) Show() {
	fmt.Println("我是苹果...")
}

type Banana1 struct {
}

func (b *Banana1) Show() {
	fmt.Println("我是香蕉...")
}

// ====工厂模块
type AppleFactory struct {
}

func (af *AppleFactory) CreateFruit() Fruit1 {
	var fruit Fruit1
	fruit = new(Apple1)

	return fruit
}

type BananaFactory struct {
}

func (ab *BananaFactory) CreateFruit() Fruit1 {
	var fruit Fruit1
	fruit = new(Banana1)

	return fruit
}

func main() {
	appleFac := new(AppleFactory)
	apple := appleFac.CreateFruit()
	apple.Show()

	bananaFac := new(BananaFactory)
	banana := bananaFac.CreateFruit()
	banana.Show()
}
