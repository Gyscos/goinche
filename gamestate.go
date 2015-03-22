package main

import "errors"

type GameState struct {
	// Player hands
	Players [4]Player

	// Next player to act
	CurrentPlayer PlayerID

	Trump Suit

	CurrentTrick Trick

	// 1 then 2 for belote & rebelote
	BeloteLevel int

	// 0 for regular games, 1 for coinched, 2 for super-coinched
	// Final score is multiplied by 2^CoincheLevel
	CoincheLevel int

	// Are we bidding, or playing?
	Bidding bool
}

func NewGame(first PlayerID) *GameState {
	result := &GameState{
		CurrentPlayer: first,
		Bidding:       true,
	}
	hands := DealHands()
	for i, h := range hands {
		result.Players[i].Cards = h
	}
	return result
}

// Find the highest trump card in a trick, right up to the given player.
func GetHigherTrump(trick Trick, trump Suit, player PlayerID) int {

	highestTrump := -1
	// Start with the trick first player.
	for p := trick.FirstPlayer; p != player; p = p.Next() {
		if trick.Cards[p].Suit() == trump {
			str := trick.Cards[p].Rank().getTrumpOrder()
			if str > highestTrump {
				highestTrump = str
			}
		}
	}
	return highestTrump
}

func HasHigherTrump(hand Hand, order int, trump Suit) bool {
	for _, r := range Ranks {
		if r.getTrumpOrder() > order && hand.Has(MakeCard(trump, r)) {
			return true
		}
	}
	return false
}

func (gs *GameState) IsValidMove(playerID PlayerID, card Card) error {
	player := &gs.Players[playerID]
	if !player.Cards.Has(card) {
		// Does he even has this card??
		return errors.New("card is not in player's hand")
	}
	if playerID == gs.CurrentTrick.FirstPlayer {
		// The first player in a trick has more choice.
		return nil
	}
	cardSuit := card.Suit()
	if cardSuit == gs.Trump {
		// Check for increasing trump
		highestTrump := GetHigherTrump(gs.CurrentTrick, gs.Trump, playerID)
		if card.Rank().getTrumpOrder() < highestTrump {
			// He didn't raise. Did he have an excuse?
			if HasHigherTrump(gs.Players[playerID].Cards, highestTrump, gs.Trump) {
				return errors.New("player did not raise on trumps, but could do so")
			}
		}
	}
	startingSuit := gs.CurrentTrick.Cards[gs.CurrentTrick.FirstPlayer].Suit()
	if cardSuit != startingSuit {
		// The only excuse to play this color is to have no card of startingSuit
		if player.Cards.HasAny(startingSuit) {
			return errors.New("card is not the right suit")
		}

		// Now, was it a trump at least?
		if cardSuit != gs.Trump {
			// If our partner is currently winning, the following rule doesn't apply
			partnerWinning := false
			if !partnerWinning && player.Cards.HasAny(gs.Trump) {
				return errors.New("cannot piss with trump in hand")
			}
		}
	}

	return nil
}

func (gs *GameState) Play(player PlayerID, card Card) error {
	if player != gs.CurrentPlayer {
		return errors.New("invalid player turn")
	}
	if gs.Bidding {
		return errors.New("cannot play while biddind")
	}

	// Is the player allowed to play this card?
	err := gs.IsValidMove(player, card)
	if err != nil {
		return err
	}

	gs.CurrentTrick.playCard(player, card, gs.Trump)

	if gs.CurrentPlayer.Next() == gs.CurrentTrick.FirstPlayer {
		// Trick is over
		// Archive it.

		// Was this the LAST TRICK???
		// GAME OVER BITCHES.
		// Broadcast to everyone
	}

	// Is the trick complete?
	gs.CurrentPlayer = gs.CurrentPlayer.Next()

	return nil
}
