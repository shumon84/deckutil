package deckutil

import "fmt"

type Card interface {
	GetID() int
	fmt.Stringer
}

type cardDict map[int]cardDictValue

type cardDictValue struct {
	index int
	card  Card
}
