package main

import "fmt"

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫在吃...")
}

type CatA struct {
	Cat // 继承
}

func (ca *CatA) Sleep() {
	fmt.Println("小猫在睡觉...")
}

type CatB struct {
	C *Cat // 组合
}

type CatC struct{}

func (cc *CatC) Eat(c *Cat) {
	c.Eat() // 组合
}

func main() {
	cat := Cat{}
	cat.Eat()
	fmt.Println("---------")
	catA := CatA{}
	catA.Eat()
	catA.Sleep()
	fmt.Println("---------")
	catB := CatB{
		C: &cat,
	}
	catB.C.Eat()
	fmt.Println("---------")
	catC := CatC{}
	catC.Eat(&cat)
}
