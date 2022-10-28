package main

import "fmt"

/*
假设江湖有一名无事不知，无话不说的大嘴巴，“江湖百晓生”，任何江湖中发生的事件都会被百晓生知晓，且进行广播。
先江湖中有两个帮派，分别为：
丐帮：黄蓉、洪七公、乔峰。
明教：张无忌、灭绝师太、金毛狮王。
现在需要用观察者模式模拟如下场景：
（1）事件一：丐帮的黄蓉把明教的张无忌揍了，这次武林事件被百晓生知晓，并且进行广播。
         主动打人方的帮派收到消息要拍手叫好。
         被打的帮派收到消息应该报酬，如：灭绝师太得知消息进行报仇，将丐帮黄蓉揍了。触发事件二。
（2）事件二：明教的灭绝师太把丐帮的黄蓉揍了，这次武林事件被百姓生知晓，并且进行广播。
*/

/*
               百晓生
	[丐帮]               [明教]
    洪七公               张无忌
    黄蓉					韦一笑
    乔峰				    金毛狮王
*/

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

//-------- 抽象层 -------
type Listener1 interface {
	//当同伴被揍了该怎么办
	OnFriendBeFight(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

type Notifier1 interface {
	//添加观察者
	AddListener1(Listener1 Listener1)
	//删除观察者
	RemoveListener1(Listener1 Listener1)
	//通知广播
	Notify(event *Event)
}

type Event struct {
	Noti    Notifier1 //被知晓的通知者
	One     Listener1 //事件主动发出者
	Another Listener1 //时间被动接收者
	Msg     string    //具体消息
}

//-------- 实现层 -------
//英雄(Listener1)
type Hero1 struct {
	Name  string
	Party string
}

func (Hero1 *Hero1) Fight(another Listener1, baixiao Notifier1) {
	msg := fmt.Sprintf("%s 将 %s 揍了...", Hero1.Title(), another.Title())

	//生成事件
	event := new(Event)
	event.Noti = baixiao
	event.One = Hero1
	event.Another = another
	event.Msg = msg

	baixiao.Notify(event)
}

func (Hero1 *Hero1) Title() string {
	return fmt.Sprintf("[%s]:%s", Hero1.Party, Hero1.Name)
}

func (Hero1 *Hero1) OnFriendBeFight(event *Event) {
	//判断是否为当事人
	if Hero1.Name == event.One.GetName() || Hero1.Name == event.Another.GetName() {
		return
	}

	//本帮派同伴将其他门派揍了，要拍手叫好!
	if Hero1.Party == event.One.GetParty() {
		fmt.Println(Hero1.Title(), "得知消息，拍手叫好！！！")
		return
	}

	//本帮派同伴被其他门派揍了，要主动报仇反击!
	if Hero1.Party == event.Another.GetParty() {
		fmt.Println(Hero1.Title(), "得知消息，发起报仇反击！！！")
		Hero1.Fight(event.One, event.Noti)
		return
	}
}

func (Hero1 *Hero1) GetName() string {
	return Hero1.Name
}

func (Hero1 *Hero1) GetParty() string {
	return Hero1.Party
}

//百晓生(Nofifier)
type BaiXiao struct {
	Hero1List []Listener1
}

//添加观察者
func (b *BaiXiao) AddListener1(Listener1 Listener1) {
	b.Hero1List = append(b.Hero1List, Listener1)
}

//删除观察者
func (b *BaiXiao) RemoveListener1(Listener1 Listener1) {
	for index, l := range b.Hero1List {
		//找到要删除的元素位置
		if Listener1 == l {
			//将删除的点前后的元素链接起来
			b.Hero1List = append(b.Hero1List[:index], b.Hero1List[index+1:]...)
			break
		}
	}
}

//通知广播
func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("【世界消息】 百晓生广播消息: ", event.Msg)
	for _, Listener1 := range b.Hero1List {
		//依次调用全部观察的具体动作
		Listener1.OnFriendBeFight(event)
	}
}

func main() {
	Hero11 := Hero1{
		"黄蓉",
		PGaiBang,
	}

	Hero12 := Hero1{
		"洪七公",
		PGaiBang,
	}

	Hero13 := Hero1{
		"乔峰",
		PGaiBang,
	}

	Hero14 := Hero1{
		"张无忌",
		PMingJiao,
	}

	Hero15 := Hero1{
		"韦一笑",
		PMingJiao,
	}

	Hero16 := Hero1{
		"金毛狮王",
		PMingJiao,
	}

	baixiao := BaiXiao{}

	baixiao.AddListener1(&Hero11)
	baixiao.AddListener1(&Hero12)
	baixiao.AddListener1(&Hero13)
	baixiao.AddListener1(&Hero14)
	baixiao.AddListener1(&Hero15)
	baixiao.AddListener1(&Hero16)

	fmt.Println("武林一片平静.....")
	Hero11.Fight(&Hero15, &baixiao)
}
