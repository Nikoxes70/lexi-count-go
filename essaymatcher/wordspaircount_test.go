package essaymatcher

import (
	"reflect"
	"testing"
)

func TestFindTopWords(t *testing.T) {
	tests := []struct {
		name       string
		wordCounts map[string]int
		n          int
		want       []wordCountPair
	}{
		{
			name: "Basic test",
			wordCounts: map[string]int{
				"hello": 10,
				"world": 5,
				"test":  8,
			},
			n: 2,
			want: []wordCountPair{
				{"hello", 10},
				{"test", 8},
			},
		},
		{
			name: "Less words than N",
			wordCounts: map[string]int{
				"hello": 10,
				"world": 5,
			},
			n: 3,
			want: []wordCountPair{
				{"hello", 10},
				{"world", 5},
			},
		},
		{
			name:       "Empty word counts",
			wordCounts: map[string]int{},
			n:          2,
			want:       []wordCountPair{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pf := NewWordCountPair()
			got := pf.FindTopWords(tt.wordCounts, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
