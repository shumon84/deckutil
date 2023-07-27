package deckutil

import (
	"fmt"
	"strings"
)

type ErrNotFound struct {
	cards []Card
}

func NewErrNotFound(cards ...Card) error {
	return ErrNotFound{
		cards: cards,
	}
}
func (e ErrNotFound) Error() string {
	ids := make([]string, 0, len(e.cards))
	for _, card := range e.cards {
		ids = append(ids, fmt.Sprintf("%03d", card.GetID()))
	}
	return "these cards (ID=" + strings.Join(ids, ", ") + ") are not found."
}

type ErrNoMoreCards struct {
}

func NewErrNoMoreCards() error {
	return ErrNoMoreCards{}
}

func (e ErrNoMoreCards) Error() string {
	return "no more cards"
}

type ErrDuplicateCards struct {
	cards []Card
}

func NewErrDuplicateCards(cards ...Card) error {
	return ErrDuplicateCards{
		cards: cards,
	}
}

func (e ErrDuplicateCards) Error() string {
	ids := make([]string, 0, len(e.cards))
	for _, card := range e.cards {
		ids = append(ids, fmt.Sprintf("%03d", card.GetID()))
	}
	return "these cards (ID=" + strings.Join(ids, ", ") + ") are duplicated"
}
