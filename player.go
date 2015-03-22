package main

type Player struct {
	Cards Hand
}

type PlayerID int

func (p PlayerID) Next() PlayerID {
	if p == 3 {
		return PlayerID(0)
	} else {
		return p + 1
	}
}

func (p PlayerID) Prev() PlayerID {
	if p == 0 {
		return PlayerID(3)
	} else {
		return p - 1
	}
}
