package blackJack

import (
	"fmt"
	"math/rand"
	"time"
)

//Player Defines the player
type Player struct {
	Cards      []Card
	PointTotal int
	Munnies    int
	Bet        int
}

//Card is a struct that defines a card
type Card struct {
	Value int
	Suit  string
	Name  string
}

//Deck is a deck of cards
type Deck []Card

//CardValue defines the value for each card
var CardValue = map[string]int{
	"Ace: ": 1,
	"Two: ": 2,
	"Three": 3,
	"Four":  4,
	"Five":  5,
	"Six":   6,
	"Seven": 7,
	"Eight": 8,
	"Nine":  9,
	"Ten":   10,
	"Jack":  10,
	"Queen": 10,
	"King":  10,
}

//NewDeck Generates a deck for BlackJack
func NewDeck() []Card {
	deck := Deck{}
	Cardnames := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	Cardsuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	for _, s := range Cardsuits {
		for _, r := range Cardnames {
			card := Card{
				Value: CardValue[r],
				Suit:  s,
				Name:  r,
			}
			deck = append(deck, card)
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

//Scores calculates the scores
func Scores(FirstCard int, SecondCard int) int {
	fmt.Println(rand.Intn(52))
	return FirstCard + SecondCard
}

//Deal appends the cards to the hand
func Deal(deck []Card) ([]Card, []Card) {
	var hand []Card
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
		randorino := rand.New(rand.NewSource(time.Now().UnixNano()))
		for len(deck) > 0 {
			n := len(deck)
			rando := randorino.Intn(n)
			deck[n-1], deck[rando] = deck[rando], deck[n-1]
			deck = deck[:n-1]
		}
	}
}

//OpenTheGame starts the games
func OpenTheGame() {
	fmt.Println("Welcome to Blackjack. I am your dealer,my name is D'Arby,D-A-R-B-Y.")

}
