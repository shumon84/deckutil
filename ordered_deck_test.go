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
		name     string
		o        *orderedDeck
		want     []Card
		wantDeck *orderedDeck
	}{
		{
			name:     "simple test",
			o:        NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			want:     makeMockCards(4),
			wantDeck: NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
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
		o := NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck)
		want := makeMockCards(4)
		got := o.RevealAllWithoutShuffle()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealAll() = %v, want %v", got, want)
		}
		got[0] = mockCard(100)
		gotCard := o.list[0]
		wantCard := mockCard(0)
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealAll() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		o    *orderedDeck
		want *orderedDeck
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck(makeMockCards(4), makeMockRand(30, 20, 10, 0)).(*orderedDeck),
			want: &orderedDeck{
				dict: map[int]cardDictValue{
					3: {
						index: 0,
						card:  mockCard(3),
					},
					2: {
						index: 1,
						card:  mockCard(2),
					},
					1: {
						index: 2,
						card:  mockCard(1),
					},
					0: {
						index: 3,
						card:  mockCard(0),
					},
				},
				list: []Card{
					mockCard(3),
					mockCard(2),
					mockCard(1),
					mockCard(0),
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
		o        *orderedDeck
		wantOut  []Card
		wantDeck *orderedDeck
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck(makeMockCards(4), makeMockRand(30, 20, 10, 0)).(*orderedDeck),
			wantOut: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
			},
			wantDeck: &orderedDeck{
				dict: map[int]cardDictValue{
					3: {
						index: 0,
						card:  mockCard(3),
					},
					2: {
						index: 1,
						card:  mockCard(2),
					},
					1: {
						index: 2,
						card:  mockCard(1),
					},
					0: {
						index: 3,
						card:  mockCard(0),
					},
				},
				list: []Card{
					mockCard(3),
					mockCard(2),
					mockCard(1),
					mockCard(0),
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
		o := NewOrderedDeck(makeMockCards(4), makeMockRand(30, 20, 10, 0)).(*orderedDeck)
		want := makeMockCards(4)
		got := o.RevealAllWithShuffle()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealAllWithShuffle() = %v, want %v", got, want)
		}
		got[0] = mockCard(100)
		gotCard := o.list[0]
		wantCard := mockCard(3)
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealAllWithShuffle() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_DrawN(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck
		arg      int
		want     []Card
		wantErr  bool
		wantDeck *orderedDeck
	}{
		{
			name:    "test for draw one",
			o:       NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:     1,
			want:    []Card{mockCard(0)},
			wantErr: false,
			wantDeck: NewOrderedDeck([]Card{
				mockCard(1),
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck),
		},
		{
			name: "test for draw two",
			o:    NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:  2,
			want: []Card{
				mockCard(0),
				mockCard(1),
			},
			wantErr: false,
			wantDeck: NewOrderedDeck([]Card{
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
			arg:      1,
			want:     []Card{},
			wantErr:  true,
			wantDeck: NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
		},
		{
			name:     "test for not enough deck",
			o:        NewOrderedDeck(makeMockCards(2), makeMockRand(0)).(*orderedDeck),
			arg:      3,
			want:     makeMockCards(2),
			wantErr:  true,
			wantDeck: NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
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
		o        *orderedDeck
		want     Card
		wantErr  bool
		wantDeck *orderedDeck
	}{
		{
			name:    "simple test",
			o:       NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			want:    mockCard(0),
			wantErr: false,
			wantDeck: NewOrderedDeck([]Card{
				mockCard(1),
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
			want:     nil,
			wantErr:  true,
			wantDeck: NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
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
		o       *orderedDeck
		arg     int
		want    []Card
		wantErr bool
	}{
		{
			name:    "test for reveal one",
			o:       NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:     1,
			want:    makeMockCards(1),
			wantErr: false,
		},
		{
			name:    "test for reveal two",
			o:       NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:     2,
			want:    makeMockCards(2),
			wantErr: false,
		},
		{
			name:    "test for reveal all",
			o:       NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:     4,
			want:    makeMockCards(4),
			wantErr: false,
		},
		{
			name:    "test for empty deck",
			o:       NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
			arg:     1,
			want:    []Card{},
			wantErr: true,
		},
		{
			name:    "test for not enough deck",
			o:       NewOrderedDeck(makeMockCards(2), makeMockRand(0)).(*orderedDeck),
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
		o := NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck)
		want := makeMockCards(4)
		got, err := o.RevealTop(4)
		if err != nil {
			t.Fatalf("RevealTop() = %v, want %v", err, nil)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealTop() = %v, want %v", got, want)
		}
		got[0] = mockCard(100)
		gotCard := o.list[0]
		wantCard := mockCard(0)
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealTop() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_orderedDeck_Search(t *testing.T) {
	tests := []struct {
		name     string
		o        *orderedDeck
		arg      Card
		want     Card
		wantErr  bool
		wantDeck *orderedDeck
	}{
		{
			name:    "simple test",
			o:       NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:     mockCard(1),
			want:    mockCard(1),
			wantErr: false,
			wantDeck: NewOrderedDeck([]Card{
				mockCard(0),
				mockCard(2),
				mockCard(3),
			}, makeMockRand(0)).(*orderedDeck),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
			arg:      mockCard(0),
			want:     nil,
			wantErr:  true,
			wantDeck: NewOrderedDeck([]Card{}, makeMockRand(0)).(*orderedDeck),
		},
		{
			name:     "test for no exists card in the deck",
			o:        NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
			arg:      mockCard(10),
			want:     nil,
			wantErr:  true,
			wantDeck: NewOrderedDeck(makeMockCards(4), makeMockRand(0)).(*orderedDeck),
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
		o        *orderedDeck
		arg      []Card
		wantDeck *orderedDeck
	}{
		{
			name: "simple test",
			o:    NewOrderedDeck(makeMockCards(4), makeMockRand()).(*orderedDeck),
			arg: []Card{
				mockCard(4),
				mockCard(5),
			},
			wantDeck: NewOrderedDeck([]Card{
				mockCard(4),
				mockCard(5),
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
			}, makeMockRand()).(*orderedDeck),
		},
		{
			name:     "test for empty deck",
			o:        NewOrderedDeck([]Card{}, makeMockRand()).(*orderedDeck),
			arg:      makeMockCards(4),
			wantDeck: NewOrderedDeck(makeMockCards(4), makeMockRand()).(*orderedDeck),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.AddTop(tt.arg...)
			tt.o.random = tt.wantDeck.random
			gotDeck := tt.o
			if !reflect.DeepEqual(gotDeck, tt.wantDeck) {
				t.Errorf("AddTop() = %v, want %v", gotDeck, tt.wantDeck)
			}
		})
	}
}
