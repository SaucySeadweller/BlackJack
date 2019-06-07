package blackJack

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

//Player Defines the player
type Player struct {
	Cards      []Card
	PointTotal int
}

//Card is a struct that defines a card
type Card struct {
	Value int
	Suit  string
	Name  string
}

//Deck is a deck of cards
type Deck []Card

//NewDeck Generates a deck for BlackJack
func NewDeck() []Card {
	deck := Deck{}
	CardNames := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	CardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	CardValue := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "10", "10"}
	for _, s := range CardSuits {
		for _, r := range CardNames {
			for v := 0; v < len(CardValue); v++ {
				card := Card{
					Value: v,
					Suit:  s,
					Name:  r,
				}
				deck = append(deck, card)
			}
		}
	}
	return deck
}

//DisplayCard displays the cards
func (d *Deck) DisplayCard() {
	for i, Card := range *d {
		fmt.Println(i, Card.Value)
	}
}

//PlayerScore calculates the player's score
func PlayerScore(FirstCard int, SecondCard int) int {
	return FirstCard + SecondCard
}

//DealerScore calculates the dealer's score
func DealerScore(FirstCard int, SecondCard int) int {
	return FirstCard + SecondCard
}

//Deal appends the cards to the hand
func Deal(deck []Card) ([]Card, []Card) {
	var hand []Card
	Shuffle(deck)
	for i := 0; i < 2; i++ {
		hand = append(hand, deck[i])
		deck[i] = deck[len(deck)-1]
		deck = deck[:len(deck)-1]

	}
	return deck, hand
}

//Shuffle randomizes the deck
func Shuffle(deck []Card) {
	for i := 0; i < 10; i++ {
		randm := rand.New(rand.NewSource(time.Now().UnixNano()))
		for len(deck) > 0 {
			n := len(deck)
			rando := randm.Intn(n)
			deck[n-1], deck[rando] = deck[rando], deck[n-1]
			deck = deck[:n-1]
		}
	}
}

//Stand creates the stand command
func Stand() bool {

	return true
}

//Hit creats the hit command
func Hit(deck []Card) bool {
	var hand []Card
	for i := 0; i < 1; i++ {
		hand = append(hand, deck[i])
		deck[i] = deck[len(deck)-1]
		deck = deck[:len(deck)-1]
	}
	return true
}

//Dealer will control how the dealer behaves when playing the game
func Dealer() {
}

//Commands are the things the player can do
func Commands(cmd string) {
	var firstCard, secondCard int
	var deck []Card
	{
		switch strings.ToLower(cmd) {
		default:
			log.Println("You cannot do that.")
		case "Stand":

		case "Hit":
			Hit(deck)
		case "Split":
		}

		if PlayerScore(firstCard, secondCard) == 21 {
			fmt.Println("Congratualtions, you beat me.")
		} else if PlayerScore(firstCard, secondCard) < 21 {
			fmt.Println("I'm sorry but you have busted, it's my win")

		} else if DealerScore(firstCard, secondCard) == 21 {
			fmt.Println("My aren't I lucky, My win")

		} else if PlayerScore(firstCard, secondCard) > DealerScore(firstCard, secondCard) && PlayerScore(firstCard, secondCard) < 22 {
			println(" Congratulations, you have won.")

		} else if PlayerScore(firstCard, secondCard) < DealerScore(firstCard, secondCard) && DealerScore(firstCard, secondCard) < 22 {
			fmt.Println("I win this round, better luck next time.")

		}

	}

}

//CardModels contains the unicode for cards
func CardModels() {
	//	🂠	🂡	🂢	🂣	🂤	🂥	🂦	🂧	🂨	🂩	🂪	🂫	🂬	🂭	🂮
	//	🂱	🂲	🂳	🂴	🂵	🂶	🂷	🂸	🂹	🂺	🂻	🂼	🂽	🂾
	//	🃁	🃂	🃃	🃄	🃅	🃆	🃇	🃈	🃉	🃊	🃋	🃌	🃍	🃎
	//	🃑	🃒	🃓	🃔	🃕	🃖	🃗	🃘	🃙	🃚	🃛	🃜	🃝	🃞
}
