package main

import "fmt"

type deck []string

func newDeck() deck {

	cards := deck{}

	cardsSuit := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardsValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuit {
		for _, value := range cardsValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
