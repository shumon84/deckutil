package deckutil

import (
	"math"
	"math/rand"
	"reflect"
)

type OrderedDeck[T Card] interface {
	Size() int
	RevealAllWithoutShuffle() []T
	RevealAllWithShuffle() []T
	Shuffle()
	Draw() (T, error)
	DrawN(n int) ([]T, error)
	RevealTop(n int) ([]T, error)
	Search(card T) (T, error)
	AddTop(cards ...T)
	AddBottom(cards ...T)
	Insert(cards ...T)
}

type orderedDeck[T Card] struct {
	dict   cardDict[T]
	list   []T
	random rand.Source
}

func NewOrderedDeck[T Card](cards []T, random rand.Source) OrderedDeck[T] {
	dict := make(cardDict[T], len(cards))
	for i, card := range cards {
		dict[card.GetID()] = cardDictValue[T]{
			index: i,
			card:  card,
		}
	}
	return &orderedDeck[T]{
		dict:   dict,
		list:   cards,
		random: random,
	}
}

func (o *orderedDeck[T]) Size() int {
	return len(o.list)
}

func (o *orderedDeck[T]) RevealAllWithoutShuffle() []T {
	out := make([]T, len(o.list))
	copy(out, o.list)
	return out
}

func (o *orderedDeck[T]) RevealAllWithShuffle() []T {
	out := o.RevealAllWithoutShuffle()
	o.Shuffle()
	return out
}

func (o *orderedDeck[T]) Shuffle() {
	r := rand.New(o.random)
	r.Shuffle(len(o.list), reflect.Swapper(o.list))
	dict := make(cardDict[T], len(o.list))
	for i, card := range o.list {
		dict[card.GetID()] = cardDictValue[T]{
			index: i,
			card:  card,
		}
	}
	o.dict = dict
}

func (o *orderedDeck[T]) Draw() (x T, _ error) {
	cards, err := o.DrawN(1)
	if err != nil {
		return x, err
	}
	return cards[0], nil
}

func (o *orderedDeck[T]) DrawN(n int) ([]T, error) {
	if n >= len(o.list) {
		list := o.list
		o.list = []T{}
		o.dict = map[int]cardDictValue[T]{}
		return list, NewErrNoMoreCards()
	}
	list := make([]T, n)
	copy(list, o.list[:n])
	o.list = o.list[n:]
	for _, card := range list {
		delete(o.dict, card.GetID())
	}
	for i, card := range o.list {
		o.dict[card.GetID()] = cardDictValue[T]{
			index: i,
			card:  card,
		}
	}
	return list, nil
}

func (o *orderedDeck[T]) RevealTop(n int) ([]T, error) {
	sep := int(math.Min(
		float64(n),
		float64(len(o.list)),
	))
	list := make([]T, sep)
	copy(list, o.list)
	if n > len(o.list) {
		return list, NewErrNoMoreCards()
	}
	return list, nil
}

func (o *orderedDeck[T]) Search(card T) (x T, _ error) {
	cardInfo, ok := o.dict[card.GetID()]
	if !ok {
		return x, NewErrNotFound(card)
	}
	delete(o.dict, card.GetID())

	newList := make([]T, 0)
	if 0 < cardInfo.index {
		newList = o.list[:cardInfo.index]
	}
	if cardInfo.index+1 < len(o.list) {
		newList = append(newList, o.list[cardInfo.index+1:]...)
	}
	o.list = newList

	for i, c := range newList[cardInfo.index:] {
		o.dict[c.GetID()] = cardDictValue[T]{
			index: cardInfo.index + i,
			card:  c,
		}
	}
	return cardInfo.card, nil
}

// TODO: カードが重複している時にエラーを返すようにする
func (o *orderedDeck[T]) AddTop(cards ...T) {
	cardsCopy := make([]T, len(cards))
	copy(cardsCopy, cards)
	o.list = append(cardsCopy, o.list...)
	for i, card := range o.list {
		o.dict[card.GetID()] = cardDictValue[T]{
			index: i,
			card:  card,
		}
	}
}

// TODO: カードが重複している時にエラーを返すようにする
func (o *orderedDeck[T]) AddBottom(cards ...T) {
	for i, card := range cards {
		o.dict[card.GetID()] = cardDictValue[T]{
			index: len(o.list) + i,
			card:  card,
		}
	}
	o.list = append(o.list, cards...)
}

// TODO: カードが重複している時にエラーを返すようにする
func (o *orderedDeck[T]) Insert(cards ...T) {
	o.list = append(o.list, cards...)
	o.Shuffle()
	return
}
