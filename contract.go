package main

type ContractBet int

type Contract struct {
	Bet    ContractBet
	Trump  Suit
	Author PlayerID
}

const (
	Bet_80 ContractBet = iota
	Bet_90
	Bet_100
	Bet_110
	Bet_120
	Bet_130
	Bet_140
	Bet_150
	Bet_160
	Bet_Capot
)
