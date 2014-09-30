package main

import "testing"

func TestHand(t *testing.T) {
	hand := Hand(0)

	cards := []Card{Diamonds_J, Hearts_8, Spades_A}
	for _, c := range cards {
		hand.Add(c)
	}

	if hand.Size() != len(cards) {
		t.Errorf("Hand should have %v cards, but has %v : %v", len(cards), hand.Size(), hand)
	}

	for _, c := range cards {
		if !hand.Has(c) {
			t.Errorf("Hand should have %v but contains %v.", c, hand)
		}
		hand.Remove(c)
	}

	if !hand.Empty() {
		t.Errorf("Hand should be empty, but contains %v.", hand)
	}

}
