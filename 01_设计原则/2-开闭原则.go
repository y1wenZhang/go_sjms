package main

import "fmt"

type AbstractBanker interface {
	Dobusi()
}

// 存款业务员
type SaveBanker struct{}

func (sb *SaveBanker) Dobusi() {
	fmt.Println("进行了 存款业务...")
}

type PayBanker struct{}

func (p *PayBanker) Dobusi() {
	fmt.Println("进行了 支付业务...")
}

func BankerDobusi(banker AbstractBanker) {
	banker.Dobusi()
}

func main() {

	// sb := SaveBanker{}

	// var ab AbstractBanker
	// ab = &sb
	// ab.Dobusi()

	// sp := PayBanker{}
	// ab = &sp
	// ab.Dobusi()
	BankerDobusi(&SaveBanker{})
	BankerDobusi(&PayBanker{})
}
