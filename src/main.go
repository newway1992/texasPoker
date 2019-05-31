package main

import (
	"fmt"
	"game"
)
func main(){
	fmt.Println("Welcome to TexasPoker!")
	room := game.NewRoom(1001)

	player1 := game.NewPlayer()
	player2 := game.NewPlayer()
	player3 := game.NewPlayer()
	player4 := game.NewPlayer()
	player5 := game.NewPlayer()

	room.SitDown(player1,0)
	room.SitDown(player2,1)
	room.SitDown(player3,2)
	room.SitDown(player4,3)
	room.SitDown(player5,4)

	player1.Chips = room.SB * 200
	player2.Chips = room.SB * 200
	player3.Chips = room.SB * 200
	player4.Chips = room.SB * 200
	player5.Chips = room.SB * 200

	player1.CardType = game.FOUR
	player1.CardValue = game.FOUR

	player2.CardType = game.ROYAL_FLUSH
	player2.CardValue = game.ROYAL_FLUSH

	player3.CardType = game.ONE_PAIR
	player3.CardValue = game.ONE_PAIR

	player4.CardType = game.TWO_PAIR
	player4.CardValue = game.TWO_PAIR

	player5.CardType = game.HIGH_CARD
	player5.CardValue = game.HIGH_CARD

	player1.DoBet(100)
	player2.DoBet(80)
	player3.DoBet(60)
	player4.DoBet(40)
	player5.DoBet(20)
	room.ShowDown()

}
