package main

import "fmt"

// --- 抽象层
type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {
}

// 实现层
type Huawei struct{}

func (hw *Huawei) Show() {
	fmt.Println("秀出了Huawei手机...")
}

type Xiaomi struct{}

func (xm *Xiaomi) Show() {
	fmt.Println("秀出了xiaomi手机...")
}

// 具体的装饰类
type MoDecorator struct {
	Decorator // 继承基础类装饰器
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("加壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

func main() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.Show()

	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()

	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei)
	keHuawei.Show()
}
