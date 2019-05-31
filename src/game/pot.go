package game

import "sort"

type BetNode struct{
	Bet int
	Pos int
}

type PotNode struct{
	Bet int
	Pos []int
}

type Bets []BetNode

func (p Bets) Len() int {
	return len(p)
}

func (p Bets) Less(i, j int) bool {
	return p[i].Bet < p[j].Bet
}

func (p Bets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//计算分池
func CalPots(chips []int)(pots []PotNode){
	var bets Bets
	for i,v := range chips{
		if v > 0{
			bets = append(bets,BetNode{v,i})
		}
	}
	//将各个玩家所下注额从小到大排序
	sort.Sort(bets)

	for i,v := range bets{
		if v.Bet > 0{
			s := bets[i:]
			//记录每份池子总注额及参与玩家座位号
			pot := PotNode{ Bet:v.Bet*len(s) }
			for j,_ := range s{
				s[j].Bet -= v.Bet
				pot.Pos = append(pot.Pos,s[j].Pos)
			}
			pots = append(pots,pot)
		}
	}
	return
}