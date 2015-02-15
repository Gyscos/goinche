package main

import "math/rand"

type Deck struct {
	Cards []Card
}

// NewDeck returns a full, sorted 32-cards deck
func NewDeck() Deck {
	result := Deck{
		Cards: make([]Card, 32),
	}

	for i, _ := range result.Cards {
		result.Cards[i] = GetCard(uint(i))
	}

	return result
}

func (d *Deck) Shuffle() {
	ids := rand.Perm(len(d.Cards))

	cards := make([]Card, len(d.Cards))

	for i, v := range ids {
		cards[v] = d.Cards[i]
	}

	d.Cards = cards
}

func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) String() string {
	result := "["
	for i, c := range d.Cards {
		if i != 0 {
			result += ","
		}
		result += c.String()
	}

	result = result + "]"
	return result
}

func (d *Deck) Empty() bool {
	return len(d.Cards) == 0
}
