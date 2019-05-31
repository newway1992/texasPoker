package game

import "fmt"

type Room struct{
	Id int
	Player []*Player
	Bets	[]int //与玩家一一对应
	Pots []int //与玩家一一对应
	Min int
	Max int
	Banker int //庄家
	SB int //小盲
}

func NewRoom(Id int)*Room{
	room := &Room{Id:Id,Min:2,Max:9,SB:2}
	room.Player = make([]*Player,9)
	room.Bets = make([]int,9)
	room.Pots = make([]int,0,9)
	return room
}

func (this *Room)ShowDown(){
	//1.统计玩家下注情况
	this.CalBet()

	//2.计算分池
	pots := CalPots(this.Bets)

	this.PrintPots(pots)

	for i,_ := range this.Bets{
		this.Bets[i] = 0
	}

	for _,pot := range pots{
		var maxPlayer *Player
		for _,pos := range pot.Pos{
			player := this.Player[pos]
			if player != nil{
				if maxPlayer == nil{
					maxPlayer = player
					continue
				}
				if player.CardValue > maxPlayer.CardValue{
					maxPlayer = player
				}
			}
		}
		var winners []int
		for _,pos := range pot.Pos{
			player := this.Player[pos]
			if player != nil && player.CardValue == maxPlayer.CardValue{
				winners = append(winners,pos)
			}
		}
		if len(winners) == 0{
			fmt.Println("no winners")
			return
		}

		//多个玩家组合牌相等平分奖池
		for _,pos := range  winners{
			this.Bets[pos] += pot.Bet / len(winners)
		}
		this.Bets[winners[0]] += pot.Bet % len(winners) // odd chips
	}
	fmt.Println("比牌结果:")
	for i,v := range this.Bets{
		player := this.Player[i]
		if player == nil{
			continue
		}
		if v> 0{
			player := this.Player[i]
			player.Chips += v
		}
		fmt.Printf("uid:%d seat:%d result:%s win:%d chips:%d\n",player.Id,player.Seat,GetResult(player.CardType),v,player.Chips)
	}
}

func (this *Room)PrintPots(pots []PotNode){
	for k,v := range  pots {
		fmt.Printf("分池%d:(%d) ", k, v.Bet)
		fmt.Print("参与玩家(座位号):")
		for _, pos := range v.Pos {
			fmt.Printf("%d ", pos)
		}
		fmt.Println()
	}
}

func (this *Room)AddPlayer(p *Player){
	p.Room = this
}

func (this *Room)RemovePlayer(p *Player){
	p.Room = this
}

func (this *Room)SitDown(player *Player,seat int){
	player.Seat = seat
	this.Player[seat] = player
}

func (this *Room)CalBet(){
	for i,v := range this.Player{
		if v != nil{
			this.Bets[i] = v.Bet
		}else{
			this.Bets[i] = 0
		}
		//fmt.Printf("i: %d bet:%d\n",i,this.Bets[i])
	}
}




