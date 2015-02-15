package main

import "testing"

func TestDeck(t *testing.T) {
	d := NewDeck()
	d.Shuffle()

	var count [32]int

	for !d.Empty() {
		c := d.Draw()
		count[c.Id()]++
	}

	for _, c := range count {
		if c != 1 {
			t.Fail()
		}
	}
}
