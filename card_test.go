package deckutil

import (
	"math/rand"
	"strconv"
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
