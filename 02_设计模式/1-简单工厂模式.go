package main

import "fmt"

// -----抽象层------
type Fruit interface {
	Show()
}

// 业务层
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("我是Apple...")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("我是Banana...")
}

func NewFruit(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		apple := Apple{}
		fruit = &apple
	} else if name == "banana" {
		fruit = new(Banana)
	}

	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.Show()

	banana := NewFruit("banana")
	banana.Show()
}
