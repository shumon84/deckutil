package deckutil

import (
	"math/rand"
	"strconv"
	"testing"
)

type mockCard int

func makeMockCards(n int) []Card {
	mock := make([]Card, n)
	for i := range mock {
		mock[i] = mockCard(i)
	}

	return mock
}

func (m mockCard) GetID() int {
	return int(m)
}

func (m mockCard) String() string {
	return strconv.Itoa(int(m))
}

type mockRandSource struct {
	index int
	loop  []int64
}

func makeMockRand(n ...int64) rand.Source {
	return &mockRandSource{
		index: 0,
		loop:  n,
	}
}

func (m *mockRandSource) Int63() int64 {
	r := m.loop[m.index]
	m.index = (m.index + 1) % len(m.loop)
	return r
}

func (m *mockRandSource) Seed(seed int64) {
	panic("no implementation")
}

func isSameCards(t *testing.T, arrA, arrB []Card) bool {
	t.Helper()
	type tuple struct {
		A bool
		B bool
	}
	ab := tuple{A: arrA == nil, B: arrB == nil}
	switch ab {
	case tuple{true, true}:
		return true
	case tuple{true, false}:
		return false
	case tuple{false, true}:
		return false
	case tuple{false, false}:
	}

	if len(arrA) != len(arrB) {
		return false
	}
	dict := make(map[int]struct{}, len(arrA))
	for _, v := range arrA {
		dict[v.GetID()] = struct{}{}
	}
	for _, v := range arrB {
		if _, ok := dict[v.GetID()]; !ok {
			return false
		}
	}
	return true
}
