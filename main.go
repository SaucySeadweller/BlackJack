package main

import (
	"fmt"

	blackJack "github.com/SaucySeadweller/blackJack/bj"
)

func main() {
	var deck []blackJack.Card = blackJack.NewDeck()
	blackJack.OpenTheGame()
	blackJack.Shuffle(deck)
	deck, PlayerHand := blackJack.Deal(deck)
	deck, DealerHand := blackJack.Deal(deck)
	fmt.Println("Player's Hand")
	for i := 0; i < len(PlayerHand); i++ {
		fmt.Println(PlayerHand[i].Name + " of " + PlayerHand[i].Suit)
	}
	fmt.Println("Dealer's Hand")
	for i := 0; i < len(DealerHand); i++ {
		fmt.Println(DealerHand[i].Name + " of " + DealerHand[i].Suit)
	}
}
