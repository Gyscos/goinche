package main

import "testing"

func TestGame(t *testing.T) {
	game := NewGame(0)
	for _, p := range game.Players {
		// fmt.Printf("Player %v: %v\n", i, p.Cards)
		if p.Cards.Size() != 8 {
			t.Fail()
		}
	}

	// Phase 1: Bids

	// Phase 2: Tricks
}
