package deckutil

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_unorderedDeck_Insert(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		u    *unorderedDeck[Card]
		args args
		want []Card
	}{
		{
			name: "test for single insert",
			u:    NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			args: args{
				cards: []Card{
					mockCard(0),
				},
			},
			want: append(makeMockCards(4), mockCard(0)),
		},
		{
			name: "test for bulk insert",
			u:    NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			args: args{
				cards: makeMockCards(2),
			},
			want: append(makeMockCards(4), makeMockCards(2)...),
		},
		{
			name: "test for insert to empty deck",
			u:    NewUnorderedDeck[Card]([]Card{}, makeMockRand(0)).(*unorderedDeck[Card]),
			args: args{
				cards: []Card{mockCard(0)},
			},
			want: []Card{mockCard(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.u.Insert(tt.args.cards...)
			if !isSameCards(t, tt.u.list, tt.want) {
				t.Errorf("Insert() got = %v, want %v", tt.u.list, tt.want)
			}
		})
	}
}

func Test_unorderedDeck_RandomTrash(t *testing.T) {
	tests := []struct {
		name     string
		u        *unorderedDeck[Card]
		want     Card
		wantErr  bool
		wantList []Card
	}{
		{
			name:    "test for case when trashed prefix card",
			u:       NewUnorderedDeck[Card](makeMockCards(4), rand.NewSource(3)).(*unorderedDeck[Card]),
			want:    mockCard(0),
			wantErr: false,
			wantList: []Card{
				mockCard(1),
				mockCard(2),
				mockCard(3),
			},
		},
		{
			name:    "test for case when trashed suffix card",
			u:       NewUnorderedDeck[Card](makeMockCards(4), rand.NewSource(15)).(*unorderedDeck[Card]),
			want:    mockCard(3),
			wantErr: false,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
			},
		},
		{
			name:    "test for case when trashed middle card",
			u:       NewUnorderedDeck[Card](makeMockCards(4), rand.NewSource(0)).(*unorderedDeck[Card]),
			want:    mockCard(2),
			wantErr: false,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(3),
			},
		},
		{
			name:     "test for case when trashed from empty deck",
			u:        NewUnorderedDeck[Card]([]Card{}, makeMockRand(0)).(*unorderedDeck[Card]),
			want:     nil,
			wantErr:  true,
			wantList: []Card{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.RandomTrash()
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomTrash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomTrash() got = %v, want %v", got, tt.want)
			}
			if !isSameCards(t, tt.u.list, tt.wantList) {
				t.Errorf("RandomTrash() got = %v, want %v", tt.u.list, tt.wantList)
			}
		})
	}
}

func Test_unorderedDeck_RevealAll(t *testing.T) {
	tests := []struct {
		name string
		u    *unorderedDeck[Card]
		want []Card
	}{
		{
			name: "simple test",
			u:    NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			want: makeMockCards(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.RevealAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevealAll() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("test for side effect", func(t *testing.T) {
		u := NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card])
		want := makeMockCards(4)
		got := u.RevealAll()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("RevealAll() = %v, want %v", got, want)
		}
		got[0] = mockCard(100)
		gotCard := u.list[0]
		wantCard := mockCard(0)
		if !reflect.DeepEqual(gotCard, wantCard) {
			t.Errorf("RevealAll() = %v, want %v", gotCard, wantCard)
		}
	})
}

func Test_unorderedDeck_Size(t *testing.T) {
	tests := []struct {
		name string
		u    *unorderedDeck[Card]
		want int
	}{
		{
			name: "simple test",
			u:    NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unorderedDeck_Trash(t *testing.T) {
	tests := []struct {
		name     string
		u        *unorderedDeck[Card]
		arg      Card
		want     Card
		wantErr  bool
		wantList []Card
	}{
		{
			name:    "test for case when trashed prefix card",
			u:       NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			arg:     mockCard(0),
			want:    mockCard(0),
			wantErr: false,
			wantList: []Card{
				mockCard(1),
				mockCard(2),
				mockCard(3),
			},
		},
		{
			name:    "test for case when trashed suffix card",
			u:       NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			arg:     mockCard(3),
			want:    mockCard(3),
			wantErr: false,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
			},
		},
		{
			name:    "test for case when trashed middle card",
			u:       NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			arg:     mockCard(2),
			want:    mockCard(2),
			wantErr: false,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(3),
			},
		},
		{
			name:     "test for case when trashed from empty deck",
			u:        NewUnorderedDeck[Card]([]Card{}, makeMockRand(0)).(*unorderedDeck[Card]),
			arg:      mockCard(1),
			want:     nil,
			wantErr:  true,
			wantList: []Card{},
		},
		{
			name:    "test for case when specified a no exists card in the deck.",
			u:       NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			arg:     mockCard(10),
			want:    nil,
			wantErr: true,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Trash(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trash() got = %v, want %v", got, tt.want)
			}
			if !isSameCards(t, tt.u.list, tt.wantList) {
				t.Errorf("Trash() got = %v, want %v", tt.u.list, tt.wantList)
			}
		})
	}
}

func Test_unorderedDeck_TrashN(t *testing.T) {
	tests := []struct {
		name     string
		u        *unorderedDeck[Card]
		arg      []Card
		want     []Card
		wantErr  bool
		wantList []Card
	}{
		{
			name:    "simple trash test",
			u:       NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			arg:     makeMockCards(2),
			want:    makeMockCards(2),
			wantErr: false,
			wantList: []Card{
				mockCard(2),
				mockCard(3),
			},
		},
		{
			name: "partial trash test",
			u:    NewUnorderedDeck[Card](makeMockCards(4), makeMockRand(0)).(*unorderedDeck[Card]),
			arg: []Card{
				mockCard(3),
				mockCard(4),
				mockCard(5),
			},
			want: []Card{
				mockCard(3),
			},
			wantErr: true,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(2),
				mockCard(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.TrashN(tt.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrashN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrashN() got = %v, want %v", got, tt.want)
			}
			if !isSameCards(t, tt.u.list, tt.wantList) {
				t.Errorf("Trash() got = %v, want %v", tt.u.list, tt.wantList)
			}
		})
	}
}
