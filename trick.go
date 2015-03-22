package main

type Trick struct {
	FirstPlayer PlayerID
	Cards       [4]Card
	Score       int
}

func (t *Trick) findWinner(trump Suit) PlayerID {
	best := 0
	bestStrength := -1
	for i, card := range t.Cards {
		strength := card.getStrength(trump)
		if strength > bestStrength {
			best = i
			bestStrength = strength
		}
	}
	return PlayerID(best)
}

func (t *Trick) playCard(player PlayerID, card Card, trump Suit) {
	t.Cards[player] = card
	t.Score += card.getScore(trump)
}
