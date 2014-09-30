package main

type Card uint32

type Suit uint32
type Rank uint32

const (
	Suits_Mask Suit = (1 << 9) - 1
	Ranks_Mask Rank = 1
)

func GetSuit(n uint) Suit {
	return Suit(1 << (8 * n))
}

func GetRank(n uint) Rank {
	return Rank(1 << n)
}

func GetCard(n uint) Card {
	return Card(1 << n)
}

const (
	Suit_Hearts Suit = 1 << (8 * iota)
	Suit_Diamonds
	Suit_Spades
	Suit_Clubs
)

const (
	Rank_7 Rank = 1 << iota
	Rank_8
	Rank_9
	Rank_X
	Rank_J
	Rank_Q
	Rank_K
	Rank_A
)

func (r Rank) String() string {
	switch r {
	case Rank_7:
		return "7"
	case Rank_8:
		return "8"
	case Rank_9:
		return "9"
	case Rank_X:
		return "X"
	case Rank_J:
		return "J"
	case Rank_Q:
		return "Q"
	case Rank_K:
		return "K"
	case Rank_A:
		return "A"
	}

	return "?"
}

func (s Suit) String() string {
	switch s {
	case Suit_Clubs:
		return "♣"
	case Suit_Hearts:
		return "♥"
	case Suit_Spades:
		return "♠"
	case Suit_Diamonds:
		return "♦"
	}

	return "?"
}

func (c Card) String() string {
	return c.Rank().String() + c.Suit().String()
}

func MakeCard(s Suit, r Rank) Card {
	return Card(uint32(s) * uint32(r))
}

func (c Card) Rank() Rank {
	v := uint32(c)

	var r uint32 = 0
	var mask uint32 = 255

	r |= (mask & v)
	v = v >> 8
	r |= (mask & v)
	v = v >> 8
	r |= (mask & v)
	v = v >> 8
	r |= v

	return Rank(r)
}

func (c Card) Suit() Suit {
	v := uint32(c)
	r := uint32(c.Rank())

	return Suit(v / r)
}

const (
	Hearts_7 Card = Card(uint32(Suit_Hearts) * uint32(Rank_7))
	Hearts_8 Card = Card(uint32(Suit_Hearts) * uint32(Rank_8))
	Hearts_9 Card = Card(uint32(Suit_Hearts) * uint32(Rank_9))
	Hearts_X Card = Card(uint32(Suit_Hearts) * uint32(Rank_X))
	Hearts_J Card = Card(uint32(Suit_Hearts) * uint32(Rank_J))
	Hearts_Q Card = Card(uint32(Suit_Hearts) * uint32(Rank_Q))
	Hearts_K Card = Card(uint32(Suit_Hearts) * uint32(Rank_K))
	Hearts_A Card = Card(uint32(Suit_Hearts) * uint32(Rank_A))

	Spades_7 Card = Card(uint32(Suit_Spades) * uint32(Rank_7))
	Spades_8 Card = Card(uint32(Suit_Spades) * uint32(Rank_8))
	Spades_9 Card = Card(uint32(Suit_Spades) * uint32(Rank_9))
	Spades_X Card = Card(uint32(Suit_Spades) * uint32(Rank_X))
	Spades_J Card = Card(uint32(Suit_Spades) * uint32(Rank_J))
	Spades_Q Card = Card(uint32(Suit_Spades) * uint32(Rank_Q))
	Spades_K Card = Card(uint32(Suit_Spades) * uint32(Rank_K))
	Spades_A Card = Card(uint32(Suit_Spades) * uint32(Rank_A))

	Diamonds_7 Card = Card(uint32(Suit_Diamonds) * uint32(Rank_7))
	Diamonds_8 Card = Card(uint32(Suit_Diamonds) * uint32(Rank_8))
	Diamonds_9 Card = Card(uint32(Suit_Diamonds) * uint32(Rank_9))
	Diamonds_X Card = Card(uint32(Suit_Diamonds) * uint32(Rank_X))
	Diamonds_J Card = Card(uint32(Suit_Diamonds) * uint32(Rank_J))
	Diamonds_Q Card = Card(uint32(Suit_Diamonds) * uint32(Rank_Q))
	Diamonds_K Card = Card(uint32(Suit_Diamonds) * uint32(Rank_K))
	Diamonds_A Card = Card(uint32(Suit_Diamonds) * uint32(Rank_A))

	Clubs_7 Card = Card(uint32(Suit_Clubs) * uint32(Rank_7))
	Clubs_8 Card = Card(uint32(Suit_Clubs) * uint32(Rank_8))
	Clubs_9 Card = Card(uint32(Suit_Clubs) * uint32(Rank_9))
	Clubs_X Card = Card(uint32(Suit_Clubs) * uint32(Rank_X))
	Clubs_J Card = Card(uint32(Suit_Clubs) * uint32(Rank_J))
	Clubs_Q Card = Card(uint32(Suit_Clubs) * uint32(Rank_Q))
	Clubs_K Card = Card(uint32(Suit_Clubs) * uint32(Rank_K))
	Clubs_A Card = Card(uint32(Suit_Clubs) * uint32(Rank_A))
)
