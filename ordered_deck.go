package deckutil

import "math/rand"

type OrderedDeck interface {
	Size() int
	RevealAllWithoutShuffle() []Card
	RevealAllWithShuffle() []Card
	Shuffle()
	Draw() (Card, error)
	DrawN(n int) ([]Card, error)
	RevealTop(n int) ([]Card, error)
	Search(card Card) (Card, error)
	AddTop(cards ...Card)
	AddBottom(cards ...Card)
	Insert(cards ...Card)
}

type orderedDeck struct {
	dict   cardDict
	list   []Card
	random rand.Source
}

func NewOrderedDeck(cards []Card, random rand.Source) OrderedDeck {
	dict := make(cardDict, len(cards))
	for i, card := range cards {
		dict[card.GetID()] = cardDictValue{
			index: i,
			card:  card,
		}
	}
	return &orderedDeck{
		dict:   dict,
		list:   cards,
		random: random,
	}
}

func (o *orderedDeck) Size() int {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) RevealAllWithoutShuffle() []Card {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) RevealAllWithShuffle() []Card {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) Shuffle() {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) Draw() (Card, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) DrawN(n int) ([]Card, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) RevealTop(n int) ([]Card, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) Search(card Card) (Card, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) AddTop(cards ...Card) {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) AddBottom(cards ...Card) {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) Insert(cards ...Card) {
	//TODO implement me
	panic("implement me")
}
