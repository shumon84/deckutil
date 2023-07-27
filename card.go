package deckutil

import "fmt"

type Card interface {
	GetID() int
	fmt.Stringer
}

type cardDict[T Card] map[int]cardDictValue[T]

type cardDictValue[T Card] struct {
	index int
	card  T
}
