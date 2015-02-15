package main

import "strings"

// Hand represents a set of cards.
type Hand uint32

// Adds the given card to the hand. It has no effect if the card is already in
// the hand.
func (h *Hand) Add(c Card) {
	*h = Hand(uint32(*h) | uint32(c))
}

// Removes the given card from the hand. It has no effect if the card is not
// in the hand.
func (h *Hand) Remove(c Card) {
	*h = Hand(uint32(*h) &^ uint32(c))
}

// Returns TRUE if the given card is in the hand.
func (h Hand) Has(c Card) bool {
	return (uint32(h) & uint32(c)) != 0
}

// Returns TRUE if the hand contains no card.
func (h Hand) Empty() bool {
	return h == 0
}

// Returns the number of cards in the hand.
func (h Hand) Size() int {
	s := 0

	for !h.Empty() {
		h.Remove(h.GetCard())
		s++
	}

	return s
}

// Returns one card from the hand. Returns 0 if the hand is empty.
func (h Hand) GetCard() Card {
	if h.Empty() {
		return Card(0)
	}

	n := (h ^ (h - 1) + 1)
	if n == 0 {
		// Oops, overflow. So we want the left-most card
		return GetCard(31)
	} else {
		return Card(n >> 1)
	}
}

// Returns a list representing the cards currently in the hand.
// Modifying the returned list will not affect the hand.
func (h Hand) List() []Card {
	result := []Card{}

	for !h.Empty() {
		c := h.GetCard()
		h.Remove(c)
		result = append(result, c)
	}

	return result
}

// Returns a text representation of the hand as a list of cards.
func (h Hand) String() string {
	cards := h.List()
	list := make([]string, len(cards))
	for i, c := range cards {
		list[i] = c.String()
	}

	return "[" + strings.Join(list, ", ") + "]"
}

// Deal n cards at a time to each player.
func DealEach(d *Deck, hands *[4]Hand, n int) {
	if len(d.Cards) < 4*n {
		panic("Deck has too few cards!")
	}

	for hi, _ := range hands {
		for i := 0; i < n; i++ {
			hands[hi].Add(d.Draw())
		}
	}
}

// Deals four hands from a full shuffled deck of cards.
func DealHands() [4]Hand {
	var result [4]Hand
	d := NewDeck()
	d.Shuffle()

	DealEach(&d, &result, 3)
	DealEach(&d, &result, 2)
	DealEach(&d, &result, 3)

	return result
}
