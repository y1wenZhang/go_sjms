package main

import "fmt"

//适配的目标
type V5 interface {
	Use5V()
}

//业务类，依赖V5接口
type Phone1 struct {
	v V5
}

func (p *Phone1) Charge() {
	fmt.Println("Phone正在充电")
	p.v.Use5V()
}

func NewPhone(v5 V5) *Phone1 {
	return &Phone1{v: v5}
}

//被适配的目标
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V充电")
}

//电源适配器
type Adapter struct {
	v220 V220
}

func (ad *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")
	ad.v220.Use220V()
}

func NewAdapter(v220 V220) *Adapter {
	return &Adapter{v220: v220}
}

func main() {
	phone := NewPhone(NewAdapter(V220{}))
	phone.Charge()
}
