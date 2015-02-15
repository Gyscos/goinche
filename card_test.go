package main

import "testing"

type Result struct {
	c Card
	r string
}

func TestCardCombinations(t *testing.T) {
	for s := uint(0); s < uint(4); s++ {
		for r := uint(0); r < uint(8); r++ {
			suit := GetSuit(s)
			rank := GetRank(r)
			c := MakeCard(suit, rank)
			if suit != c.Suit() || rank != c.Rank() {
				t.Errorf("%v + %v became %v + %v", suit, rank, c.Suit(), c.Rank())
			}
		}
	}
}

func TestCardIds(t *testing.T) {
	for i := uint(0); i < uint(32); i++ {
		c := GetCard(i)
		if c.Id() != i {
			t.Fail()
		}
	}
}

func TestCardStrings(t *testing.T) {
	data := []Result{
		{c: Hearts_8, r: "8♥"},
		{c: Hearts_J, r: "J♥"},
		{c: Diamonds_7, r: "7♦"},
		{c: Diamonds_9, r: "9♦"},
		{c: Spades_X, r: "X♠"},
		{c: Spades_Q, r: "Q♠"},
		{c: Clubs_K, r: "K♣"},
		{c: Clubs_A, r: "A♣"}}

	for _, d := range data {
		if d.c.String() != d.r {
			t.Errorf("Cannot serialize %v, got %v.", d.r, d.c)
		}
	}
}
