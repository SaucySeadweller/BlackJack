package main

import (
	"fmt"

	blackJack "github.com/SaucySeadweller/blackJack/bj"
)

func main() {
	var deck []blackJack.Card = blackJack.NewDeck()
	blackJack.OpenTheGame()
	blackJack.Shuffle(deck)
	blackJack.Deal(deck)
	fmt.Println(deck)
}
