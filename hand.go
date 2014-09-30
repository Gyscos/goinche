package main

import "strings"

type Hand uint32

func (h *Hand) Add(c Card) {
	*h = Hand(uint32(*h) | uint32(c))
}

func (h *Hand) Remove(c Card) {
	*h = Hand(uint32(*h) &^ uint32(c))
}

func (h Hand) Has(c Card) bool {
	return (uint32(h) & uint32(c)) != 0
}

func (h Hand) Empty() bool {
	return h == 0
}

func (h Hand) Size() int {
	s := 0

	for !h.Empty() {
		h.Remove(h.GetCard())
		s++
	}

	return s
}

func (h Hand) GetCard() Card {
	return Card((h ^ (h - 1) + 1) >> 1)
}

func (h Hand) List() []Card {
	result := []Card{}

	for !h.Empty() {
		c := h.GetCard()
		h.Remove(c)
		result = append(result, c)
	}

	return result
}

func (h Hand) String() string {
	cards := h.List()
	list := make([]string, len(cards))
	for i, c := range cards {
		list[i] = c.String()
	}

	return "[" + strings.Join(list, ", ") + "]"
}
