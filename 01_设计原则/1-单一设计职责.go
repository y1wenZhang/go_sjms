package main

import "fmt"

type Clothes struct{}

func (c *Clothes) OnWork() {
	fmt.Println("工作状态的装扮")
}

func (c *Clothes) OnShop() {
	fmt.Println("工作状态的装扮")
}

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街中的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("工作中的装扮")
}

func main() {
	cw := ClothesWork{}
	cw.Style()

	cs := ClothesShop{}
	cs.Style()

}
