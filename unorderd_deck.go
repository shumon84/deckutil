package deckutil

import "math/rand"

type UnorderedDeck[T Card] interface {
	Size() int
	RevealAll() []T
	RandomTrash() (T, error)
	Trash(card T) (T, error)
	TrashN(cards ...T) ([]T, error)
	Insert(cards ...T)
}

type unorderedDeck[T Card] struct {
	dict   cardDict[T]
	list   []T
	random rand.Source
}

func NewUnorderedDeck[T Card](cards []T, random rand.Source) UnorderedDeck[T] {
	dict := make(cardDict[T], len(cards))
	for i, card := range cards {
		dict[card.GetID()] = cardDictValue[T]{
			index: i,
			card:  card,
		}
	}
	return &unorderedDeck[T]{
		dict:   dict,
		list:   cards,
		random: random,
	}
}

func (u *unorderedDeck[T]) Size() int {
	return len(u.dict)
}

func (u *unorderedDeck[T]) RevealAll() []T {
	out := make([]T, len(u.list))
	copy(out, u.list)
	return out
}

func (u *unorderedDeck[T]) RandomTrash() (x T, _ error) {
	if len(u.list) == 0 {
		return x, NewErrNoMoreCards()
	}
	index := rand.New(u.random).Intn(len(u.dict))
	card := u.list[index]
	return u.Trash(card)
}

func (u *unorderedDeck[T]) Trash(card T) (x T, _ error) {
	cards, err := u.TrashN(card)
	if err != nil {
		return x, err
	}
	return cards[0], nil
}

func (u *unorderedDeck[T]) TrashN(cards ...T) ([]T, error) {
	cardInfos := make([]cardDictValue[T], 0, len(cards))
	notFoundCards := make([]Card, 0)
	foundCards := make([]T, 0, len(cards))
	for _, card := range cards {
		cardInfo, ok := u.dict[card.GetID()]
		if ok {
			foundCards = append(foundCards, cardInfo.card)
			cardInfos = append(cardInfos, cardInfo)
		} else {
			notFoundCards = append(notFoundCards, card)
		}
	}
	if len(notFoundCards) != 0 {
		return foundCards, NewErrNotFound(notFoundCards...)
	}
	for _, cardInfo := range cardInfos {
		delete(u.dict, cardInfo.card.GetID())
	}
	newList := make([]T, len(u.dict))
	index := 0
	for id, cardInfo := range u.dict {
		newList[index] = cardInfo.card
		cardInfo.index = index
		u.dict[id] = cardInfo
		index++
	}
	u.list = newList
	return foundCards, nil
}

func (u *unorderedDeck[T]) Insert(cards ...T) {
	for i, card := range cards {
		u.dict[card.GetID()] = cardDictValue[T]{
			index: len(u.list) + i,
			card:  card,
		}
	}
	u.list = append(u.list, cards...)
}
