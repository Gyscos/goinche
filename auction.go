package main

import "errors"

type Auction struct {
	game      *GameState
	passCount int

	History []*Contract
	Current *Contract
}

func (a *Auction) Bid(c *Contract) error {
	if a.Current != nil {
		if a.Current.Bet >= c.Bet {
			return errors.New("New bet should be higher than current!")
		}
		a.History = append(a.History, a.Current)
	}
	a.Current = c
	if c.Bet == Bet_Capot {
		a.complete()
	}

	return nil
}

func (a *Auction) Pass() bool {
	a.passCount++

	result := a.passCount > 3
	if result {
		a.complete()
	}
	return result
}

func (a *Auction) complete() {
	a.game.CurrentContract = a.Current
	a.game.Trump = a.Current.Trump
	a.game.Bidding = false
}
