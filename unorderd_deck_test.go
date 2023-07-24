package deckutil

import (
	"reflect"
	"testing"
)

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

func Test_unorderedDeck_Insert(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		u    *unorderedDeck
		args args
		want []Card
	}{
		{
			name: "test for single insert",
			u:    NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
			args: args{
				cards: []Card{
					mockCard(0),
				},
			},
			want: append(makeMockCards(4), mockCard(0)),
		},
		{
			name: "test for bulk insert",
			u:    NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
			args: args{
				cards: makeMockCards(2),
			},
			want: append(makeMockCards(4), makeMockCards(2)...),
		},
		{
			name: "test for insert to empty deck",
			u:    NewUnorderedDeck([]Card{}, makeMockRand(0)).(*unorderedDeck),
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
		u        *unorderedDeck
		want     Card
		wantErr  bool
		wantList []Card
	}{
		{
			name:    "test for case when trashed prefix card",
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(3)).(*unorderedDeck),
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
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(2)).(*unorderedDeck),
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
			u:        NewUnorderedDeck([]Card{}, makeMockRand(0)).(*unorderedDeck),
			want:     nil,
			wantErr:  true,
			wantList: []Card{},
		},
		{
			name:    "test for case when rand.Source returned too big number",
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(10)).(*unorderedDeck),
			want:    mockCard(2),
			wantErr: false,
			wantList: []Card{
				mockCard(0),
				mockCard(1),
				mockCard(3),
			},
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
		u    *unorderedDeck
		want []Card
	}{
		{
			name: "simple test",
			u:    NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
}

func Test_unorderedDeck_Size(t *testing.T) {
	tests := []struct {
		name string
		u    *unorderedDeck
		want int
	}{
		{
			name: "simple test",
			u:    NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
		u        *unorderedDeck
		arg      Card
		want     Card
		wantErr  bool
		wantList []Card
	}{
		{
			name:    "test for case when trashed prefix card",
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
			u:        NewUnorderedDeck([]Card{}, makeMockRand(0)).(*unorderedDeck),
			arg:      mockCard(1),
			want:     nil,
			wantErr:  true,
			wantList: []Card{},
		},
		{
			name:    "test for case when specified a no exists card in the deck.",
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
		u        *unorderedDeck
		arg      []Card
		want     []Card
		wantErr  bool
		wantList []Card
	}{
		{
			name:    "simple trash test",
			u:       NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
			u:    NewUnorderedDeck(makeMockCards(4), makeMockRand(0)).(*unorderedDeck),
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
