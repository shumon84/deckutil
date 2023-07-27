package deckutil

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_orderedDeck_Size(t *testing.T) {
	tests := []struct {
		name string
		o    *orderedDeck[Card]
		want int
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
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
		name     string
		o        *orderedDeck[Card]
		want     []Card
		wantDeck *orderedDeck[Card]
	}{
		{
			name:     "simple test",
			o:        NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			want:     makeMockCards(4),
			wantDeck: NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.RevealAllWithoutShuffle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevealAllWithoutShuffle() = %v, want %v", got, tt.want)
			}
			gotDeck := tt.o
			gotDeck.random = tt.wantDeck.random
			if !reflect.DeepEqual(tt.o, tt.wantDeck) {
				t.Errorf("RevealAllWithoutShuffle() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}

	t.Run("test for side effect", func(t *testing.T) {
		o := NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card])
		want := makeMockCards(4)
		got := o.RevealAllWithoutShuffle()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealAll() = %v, want %v", got, want)
		}
		wantCard := o.list[0]
		got[0] = mockCard(100)
		gotCard := o.list[0]
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealAll() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		o    *orderedDeck[Card]
		want *orderedDeck[Card]
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck[Card](makeMockCards(4), rand.NewSource(0)).(*orderedDeck[Card]),
			want: &orderedDeck[Card]{
				dict: map[int]cardDictValue[Card]{
					2: {
						index: 0,
						card:  mockCard(2),
					},
					1: {
						index: 1,
						card:  mockCard(1),
					},
					0: {
						index: 2,
						card:  mockCard(0),
					},
					3: {
						index: 3,
						card:  mockCard(3),
					},
				},
				list: []Card{
					mockCard(2),
					mockCard(1),
					mockCard(0),
					mockCard(3),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.Shuffle()
			tt.o.random = tt.want.random
			got := tt.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedDeck_RevealAllWithShuffle(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		wantOut  []Card
		wantDeck *orderedDeck[Card]
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck[Card](makeMockCards(4), rand.NewSource(0)).(*orderedDeck[Card]),
			wantOut: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
			},
			wantDeck: &orderedDeck[Card]{
				dict: map[int]cardDictValue[Card]{
					2: {
						index: 0,
						card:  mockCard(2),
					},
					1: {
						index: 1,
						card:  mockCard(1),
					},
					0: {
						index: 2,
						card:  mockCard(0),
					},
					3: {
						index: 3,
						card:  mockCard(3),
					},
				},
				list: []Card{
					mockCard(2),
					mockCard(1),
					mockCard(0),
					mockCard(3),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.o.RevealAllWithShuffle(); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("RevealAllWithShuffle() = %v, want %v", gotOut, tt.wantOut)
			}
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("RevealAllWitShuffle() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}

	t.Run("test for side effect", func(t *testing.T) {
		o := NewOrderedDeck[Card](makeMockCards(4), rand.NewSource(0)).(*orderedDeck[Card])
		want := makeMockCards(4)
		got := o.RevealAllWithShuffle()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealAllWithShuffle() = %v, want %v", got, want)
		}
		wantCard := o.list[0]
		got[0] = mockCard(100)
		gotCard := o.list[0]
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealAllWithShuffle() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_DrawN(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		arg      int
		want     []Card
		wantErr  bool
		wantDeck *orderedDeck[Card]
	}{
		{
			name:    "test for draw one",
			o:       NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:     1,
			want:    []Card{mockCard(0)},
			wantErr: false,
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(1),
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck[Card]),
		},
		{
			name: "test for draw two",
			o:    NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:  2,
			want: []Card{
				mockCard(0),
				mockCard(1),
			},
			wantErr: false,
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck[Card]),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
			arg:      1,
			want:     []Card{},
			wantErr:  true,
			wantDeck: NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
		},
		{
			name:     "test for not enough deck",
			o:        NewOrderedDeck[Card](makeMockCards(2), makeMockRand(0)).(*orderedDeck[Card]),
			arg:      3,
			want:     makeMockCards(2),
			wantErr:  true,
			wantDeck: NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.DrawN(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DrawN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DrawN() got = %v, want %v", got, tt.want)
			}
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("DrawN() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}
}

func Test_orderedDeck_Draw(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		want     Card
		wantErr  bool
		wantDeck *orderedDeck[Card]
	}{
		{
			name:    "simple test",
			o:       NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			want:    mockCard(0),
			wantErr: false,
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(1),
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck[Card]),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
			want:     nil,
			wantErr:  true,
			wantDeck: NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Draw()
			if (err != nil) != tt.wantErr {
				t.Errorf("Draw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Draw() got = %v, want %v", got, tt.want)
			}
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("Draw() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}
}

func Test_orderedDeck_RevealTop(t *testing.T) {
	tests := []struct {
		name    string
		o       *orderedDeck[Card]
		arg     int
		want    []Card
		wantErr bool
	}{
		{
			name:    "test for reveal one",
			o:       NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:     1,
			want:    makeMockCards(1),
			wantErr: false,
		},
		{
			name:    "test for reveal two",
			o:       NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:     2,
			want:    makeMockCards(2),
			wantErr: false,
		},
		{
			name:    "test for reveal all",
			o:       NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:     4,
			want:    makeMockCards(4),
			wantErr: false,
		},
		{
			name:    "test for empty deck",
			o:       NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
			arg:     1,
			want:    []Card{},
			wantErr: true,
		},
		{
			name:    "test for not enough deck",
			o:       NewOrderedDeck[Card](makeMockCards(2), makeMockRand(0)).(*orderedDeck[Card]),
			arg:     3,
			want:    makeMockCards(2),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.RevealTop(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("RevealTop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevealTop() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("test for side effect", func(t *testing.T) {
		o := NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card])
		want := makeMockCards(4)
		got, err := o.RevealTop(4)
		if err != nil {
			t.Fatalf("RevealTop() = %v, want %v", err, nil)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealTop() = %v, want %v", got, want)
		}
		wantCard := o.list[0]
		got[0] = mockCard(100)
		gotCard := o.list[0]
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealTop() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_Search(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		arg      Card
		want     Card
		wantErr  bool
		wantDeck *orderedDeck[Card]
	}{
		{
			name:    "simple test",
			o:       NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:     mockCard(1),
			want:    mockCard(1),
			wantErr: false,
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(0),
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck[Card]),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
			arg:      mockCard(0),
			want:     nil,
			wantErr:  true,
			wantDeck: NewOrderedDeck[Card]([]Card{}, makeMockRand(0)).(*orderedDeck[Card]),
		},
		{
			name:     "test for no exists card in the deck",
			o:        NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
			arg:      mockCard(10),
			want:     nil,
			wantErr:  true,
			wantDeck: NewOrderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Search(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("Search() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}
}

func Test_orderedDeck_AddTop(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		arg      []Card
		wantErr  bool
		wantDeck *orderedDeck[Card]
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck[Card](makeMockCards(4), makeMockRand()).(*orderedDeck[Card]),
			arg: []Card{
				mockCard(4),
				mockCard(5),
			},
			wantErr: false,
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(4),
				mockCard(5),
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
			}, makeMockRand()).(*orderedDeck[Card]),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck[Card]([]Card{}, makeMockRand()).(*orderedDeck[Card]),
			arg:      makeMockCards(4),
			wantErr:  false,
			wantDeck: NewOrderedDeck[Card](makeMockCards(4), makeMockRand()).(*orderedDeck[Card]),
		},
		{
			name:     "test for duplicate card",
			o:        NewOrderedDeck[Card](makeMockCards(4), makeMockRand()).(*orderedDeck[Card]),
			arg:      []Card{mockCard(0)},
			wantErr:  true,
			wantDeck: NewOrderedDeck[Card](makeMockCards(4), makeMockRand()).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.o.AddTop(tt.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("AddTop() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}

	t.Run("test for input side effect", func(t *testing.T) {
		o := NewOrderedDeck[Card]([]Card{mockCard(10)}, makeMockRand(0)).(*orderedDeck[Card])
		testcase := make([]Card, 1, 2)
		testcase[0] = mockCard(0)
		want := []Card{
			mockCard(0),
			mockCard(10),
		}
		o.AddTop(testcase...)
		got := o.list
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddTop() = %v, want %v", got, want)
		}
		wantCard := o.list[0]
		testcase[0] = mockCard(100)
		gotCard := o.list[0]
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("AddTop() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_AddBottom(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		arg      []Card
		wantDeck *orderedDeck[Card]
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck[Card](makeMockCards(4), makeMockRand()).(*orderedDeck[Card]),
			arg: []Card{
				mockCard(4),
				mockCard(5),
			},
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
				mockCard(4),
				mockCard(5),
			}, makeMockRand()).(*orderedDeck[Card]),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck[Card]([]Card{}, makeMockRand()).(*orderedDeck[Card]),
			arg:      makeMockCards(4),
			wantDeck: NewOrderedDeck[Card](makeMockCards(4), makeMockRand()).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.AddBottom(tt.arg...)
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("AddBottom() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}

	t.Run("test for input side effect", func(t *testing.T) {
		o := NewOrderedDeck[*mockCardStruct]([]*mockCardStruct{{0}}, makeMockRand(0)).(*orderedDeck[*mockCardStruct])
		testcase := make([]*mockCardStruct, 1, 2)
		testcase[0] = &mockCardStruct{10}
		want := []*mockCardStruct{
			{0},
			{10},
		}
		o.AddBottom(testcase...)
		got := o.list
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddBottom() = %v, want %v", got, want)
		}
		wantCard := o.list[1]
		testcase[0].int = 100
		gotCard := o.list[1]
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("AddBottom() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_Insert(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck[Card]
		arg      []Card
		wantDeck *orderedDeck[Card]
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck[Card](makeMockCards(4), rand.NewSource(0)).(*orderedDeck[Card]),
			arg: []Card{
				mockCard(4),
				mockCard(5),
			},
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(4),
				mockCard(3),
				mockCard(0),
				mockCard(2),
				mockCard(1),
				mockCard(5),
			}, makeMockRand()).(*orderedDeck[Card]),
		},
		{
			name: "test for empty deck",
			o:    NewOrderedDeck[Card]([]Card{}, rand.NewSource(0)).(*orderedDeck[Card]),
			arg:  makeMockCards(4),
			wantDeck: NewOrderedDeck[Card]([]Card{
				mockCard(2),
				mockCard(1),
				mockCard(0),
				mockCard(3),
			}, makeMockRand()).(*orderedDeck[Card]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.Insert(tt.arg...)
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("Insert() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}

	t.Run("test for input side effect", func(t *testing.T) {
		o := NewOrderedDeck[*mockCardStruct]([]*mockCardStruct{{0}}, rand.NewSource(0)).(*orderedDeck[*mockCardStruct])
		testcase := make([]*mockCardStruct, 1, 2)
		testcase[0] = &mockCardStruct{10}
		want := []*mockCardStruct{
			{0},
			{10},
		}
		o.Insert(testcase...)
		got := o.list
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddBottom() = %v, want %v", got, want)
		}
		wantCard := o.list[1]
		testcase[0].int = 100
		gotCard := o.list[1]
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("AddBottom() = %v, want %v", gotCard, wantCard)
		}
	})
}
