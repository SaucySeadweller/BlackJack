package main

import (
	"fmt"

	blackJack "github.com/SaucySeadweller/blackJack/bj"
)

func main() {
	var firstCard, secondCard int
	fmt.Println("Welcome to Blackjack. I am your dealer,my name is D'Arby,D-A-R-B-Y.")
	//var cmd string
	var deck []blackJack.Card = blackJack.NewDeck()
	deck, PlayerHand := blackJack.Deal(deck)
	deck, DealerHand := blackJack.Deal(deck)

	//blackJack.Shuffle(deck)

	fmt.Println("Let us begin... Open the game!")

	//blackJack.Commands(cmd)
	fmt.Println("[Player's Hand]")
	for i := 0; i < len(PlayerHand); i++ {
		fmt.Println(PlayerHand[i].Name+" of "+PlayerHand[i].Suit, blackJack.PlayerScore(firstCard, secondCard))

	}

	fmt.Println("[Dealer's Hand]")
	for i := 0; i < len(DealerHand); i++ {
		fmt.Println(DealerHand[i].Name+" of "+DealerHand[i].Suit, blackJack.DealerScore(firstCard, secondCard))
	}
	//fmt.Println("Would you like to hit or stand?")
}
