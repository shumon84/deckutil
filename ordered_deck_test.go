package deckutil

import (
	"reflect"
	"testing"
)

func Test_orderedDeck_Size(t *testing.T) {
	tests := []struct {
		name string
		o    *orderedDeck
		want int
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedDeck_RevealAllWithoutShuffle(t *testing.T) {
	tests := []struct {
		name string
		o    *orderedDeck
		want []Card
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			want: makeMockCards(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.RevealAllWithoutShuffle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevealAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
