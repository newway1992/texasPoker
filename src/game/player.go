package game

import "fmt"

type Player struct{
	Id int
	Name string
	Room *Room
	Seat int
	CardType int //牌型
	CardValue int	//牌型权值
	Bet int	//当轮下注
	Chips int	//总筹码
}

var Id = 1000

func NewPlayer()*Player{
	player := &Player{Id:generateUserId()}
	return player
}

func generateUserId()int{
	Id++
	return Id
}

//下注
func (this *Player)DoBet(bet int){
	this.Bet += bet
	this.Chips -= bet
	fmt.Printf("uid:%d seat:%d bet:%d\n",this.Id,this.Seat,bet)
}

