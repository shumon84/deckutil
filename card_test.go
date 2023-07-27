package deckutil

import (
	"math/rand"
	"strconv"
	"testing"
)

type mockCard int
type mockCardStruct struct {
	int
}

func (m *mockCardStruct) GetID() int {
	return m.int
}

func (m *mockCardStruct) String() string {
	return strconv.Itoa(m.int)
}

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
	return r << 32
}

func (m *mockRandSource) Seed(seed int64) {
	panic("no implementation")
}

func isSameCards(t *testing.T, arrA, arrB []Card) bool {
	t.Helper()

	switch [2]bool{arrA == nil, arrB == nil} {
	case [2]bool{true, true}:
		return true
	case [2]bool{true, false}:
		return false
	case [2]bool{false, true}:
		return false
	case [2]bool{false, false}:
		// not return
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
