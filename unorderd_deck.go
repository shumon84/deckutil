package deckutil

import "math/rand"

type UnorderedDeck interface {
	Size() int
	RevealAll() []Card
	RandomTrash() (Card, error)
	Trash(card Card) (Card, error)
	TrashN(cards ...Card) ([]Card, error)
	Insert(cards ...Card)
}

type unorderedDeck struct {
	dict   cardDict
	list   []Card
	random rand.Source
}

func NewUnorderedDeck(cards []Card, random rand.Source) UnorderedDeck {
	dict := make(cardDict, len(cards))
	for i, card := range cards {
		dict[card.GetID()] = cardDictValue{
			index: i,
			card:  card,
		}
	}
	return &unorderedDeck{
		dict:   dict,
		list:   cards,
		random: random,
	}
}

func (u *unorderedDeck) Size() int {
	return len(u.dict)
}

func (u *unorderedDeck) RevealAll() []Card {
	return u.list
}

func (u *unorderedDeck) RandomTrash() (Card, error) {
	if len(u.list) == 0 {
		return nil, NewErrNoMoreCards()
	}
	index := int(u.random.Int63()) % len(u.dict)
	card := u.list[index]
	return u.Trash(card)
}

func (u *unorderedDeck) Trash(card Card) (Card, error) {
	cards, err := u.TrashN(card)
	if err != nil {
		return nil, err
	}
	return cards[0], nil
}

func (u *unorderedDeck) TrashN(cards ...Card) ([]Card, error) {
	cardInfos := make([]cardDictValue, 0, len(cards))
	notFoundCards := make([]Card, 0)
	foundCards := make([]Card, 0, len(cards))
	for _, card := range cards {
		cardInfo, ok := u.dict[card.GetID()]
		if ok {
			foundCards = append(foundCards, card)
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
	newList := make([]Card, len(u.dict))
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

func (u *unorderedDeck) Insert(cards ...Card) {
	for i, card := range cards {
		u.dict[card.GetID()] = cardDictValue{
			index: len(u.list) + i,
			card:  card,
		}
	}
	u.list = append(u.list, cards...)
}
