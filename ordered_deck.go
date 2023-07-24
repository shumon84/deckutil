package deckutil

type OrderedDeck interface {
	Size() int
	Shuffle()
	Draw(n int) (Card, error)
	Reveal(n int) ([]Card, error)
	RevealAll() []Card
	Search(card Card) (Card, error)
	AddTop(cards ...Card)
	AddBottom(cards ...Card)
	Insert(cards ...Card)
}

type orderedDeck struct {
	list map[int]Card
}
