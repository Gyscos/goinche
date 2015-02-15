package main

import "errors"

type Auction struct {
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

	return nil
}
