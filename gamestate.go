package main

type GameState struct {
	// Player hands
	Players [4]Player

	// Next player to act
	CurrentPlayer int

	// 0 for regular games, 1 for coinched, 2 for super-coinched
	// Final score is multiplied by 2^CoincheLevel
	CoincheLevel int
}

func NewGame(first int) *GameState {
	result := &GameState{
		CurrentPlayer: first,
	}
	hands := DealHands()
	for i, h := range hands {
		result.Players[i].Cards = h
	}
	return result
}
