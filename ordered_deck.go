package deckutil

import (
	"math/rand"
	"sort"
)

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
	return len(o.list)
}

func (o *orderedDeck) RevealAllWithoutShuffle() []Card {
	return o.list
}

func (o *orderedDeck) RevealAllWithShuffle() []Card {
	//TODO implement me
	panic("implement me")
}

func (o *orderedDeck) Shuffle() {
	type tuple struct {
		r int64
		c Card
	}
	shuffle := make([]tuple, len(o.list))
	for i, card := range o.list {
		shuffle[i] = tuple{
			o.random.Int63(),
			card,
		}
	}
	sort.Slice(shuffle, func(i, j int) bool {
		return shuffle[i].r < shuffle[j].r
	})

	cards := make([]Card, len(o.list))
	dict := make(cardDict, len(cards))
	for i, t := range shuffle {
		cards[i] = t.c
		dict[t.c.GetID()] = cardDictValue{
			index: i,
			card:  t.c,
		}
	}
	o.list = cards
	o.dict = dict
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
