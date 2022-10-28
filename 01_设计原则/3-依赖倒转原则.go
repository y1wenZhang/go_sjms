package main

import "fmt"

type Car interface {
	Run()
}

type Person interface {
	Drive(car Car)
}

type Benz struct{}

func (bz *Benz) Run() {
	fmt.Println("benz is running...")
}

type Bmw struct{}

func (bm *Bmw) Run() {
	fmt.Println("bmw is running...")
}

type Zhangsan struct{}

func (z3 *Zhangsan) Drive(car Car) {
	fmt.Println("zhang3 is drive...")
	car.Run()
}

type Lisi struct{}

func (l4 *Lisi) Drive(car Car) {
	fmt.Println("lisi is drive...")
	car.Run()
}

func main() {
	var benz Car
	bz := Benz{}
	benz = &bz
	// benz = new(Benz)

	var bmw Car
	bm := Bmw{}
	bmw = &bm

	var zhangsan Person
	z3 := Zhangsan{}
	zhangsan = &z3
	// zhangsan = new(Zhangsan)
	zhangsan.Drive(benz)
	zhangsan.Drive(bmw)
}
