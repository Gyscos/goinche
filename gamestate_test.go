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
	a := game.StartAuction()
	a.Bid(&Contract{
		Bet:    Bet_80,
		Trump:  Suit_Hearts,
		Author: 0,
	})
	a.Pass()
	a.Pass()
	a.Bid(&Contract{
		Bet:    Bet_100,
		Trump:  Suit_Spades,
		Author: 3,
	})
	a.Pass()
	a.Pass()
	a.Pass()

	// Phase 2: Tricks
}
