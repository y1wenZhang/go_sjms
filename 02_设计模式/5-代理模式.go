package main

import "fmt"

type Goods struct {
	Kinds string
	Fact  bool
}

type Shopping interface {
	Buy(goods *Goods)
}

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(g *Goods) {
	fmt.Println("在韩国购买了:", g.Kinds)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(g *Goods) {
	fmt.Println("在美国购买了:", g.Kinds)
}

type AfricaShopping struct{}

func (af *AfricaShopping) Buy(g *Goods) {
	fmt.Println("在非洲购买了:", g.Kinds)
}

type ProxyShopping struct {
	shopping Shopping
}

func (ps *ProxyShopping) Buy(g *Goods) {
	// 1. 先验货
	if ps.distinguish(g) == true {
		//2. 进行购买
		ps.shopping.Buy(g) //调用原被代理的具体主题任务
		//3 海关安检
		ps.check(g)
	}
}

//验货流程
func (ps *ProxyShopping) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kinds, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kinds, ", 不应该购买。")
	}
	return goods.Fact
}

//安检流程
func (ps *ProxyShopping) check(goods *Goods) {
	fmt.Println("对[", goods.Kinds, "] 进行了海关检查， 成功的带回祖国")
}

func NewProxy(shopping Shopping) Shopping {
	return &ProxyShopping{shopping: shopping}
}

func main() {
	g1 := Goods{
		Kinds: "韩国面膜",
		Fact:  true,
	}
	g2 := Goods{Kinds: "CET4证书", Fact: false}

	var shopping Shopping
	shopping = new(AmericanShopping)
	var proxy Shopping
	proxy = NewProxy(shopping)
	proxy.Buy(&g1)
	proxy.Buy(&g2)
}
