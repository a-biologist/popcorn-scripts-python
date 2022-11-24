package slice_test

import (
	"fmt"
	"github.com/stein-f/popcorn-scripts/lang/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	tests := map[string]struct {
		gotSlice     []int
		gotElement   int
		wantContains bool
	}{
		"contains": {
			gotSlice:     []int{1, 2, 3},
			gotElement:   3,
			wantContains: true,
		},
		"does not contain": {
			gotSlice:     []int{1, 2, 3},
			gotElement:   4,
			wantContains: false,
		},
		"does not contain with empty slice": {
			gotSlice:     []int{},
			gotElement:   4,
			wantContains: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			contains := slice.Contains(test.gotSlice, test.gotElement)
			assert.Equal(t, test.wantContains, contains)
		})
	}
}

func TestUnique(t *testing.T) {
	tests := map[string]struct {
		gotSlice           []int
		wantUniqueElements []int
	}{
		"already unique": {
			gotSlice:           []int{1, 2, 3},
			wantUniqueElements: []int{1, 2, 3},
		},
		"removes dupes": {
			gotSlice:           []int{1, 2, 2},
			wantUniqueElements: []int{1, 2},
		},
		"handles empty slice": {
			gotSlice:           []int{},
			wantUniqueElements: nil,
		},
		"handles nil slice": {
			gotSlice:           nil,
			wantUniqueElements: nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := slice.Unique(test.gotSlice)
			assert.Equal(t, test.wantUniqueElements, result)
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := map[string]struct {
		gotSlice         []int
		gotElementToFind int
		wantIndex        int
	}{
		"finds index": {
			gotSlice:         []int{1, 2, 3},
			gotElementToFind: 2,
			wantIndex:        1,
		},
		"does not find index": {
			gotSlice:         []int{1, 2, 2},
			gotElementToFind: 4,
			wantIndex:        -1,
		},
		"handles empty slice": {
			gotSlice:         []int{},
			gotElementToFind: 1,
			wantIndex:        -1,
		},
		"handles nil slice": {
			gotSlice:         nil,
			gotElementToFind: 1,
			wantIndex:        -1,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := slice.FindIndex(test.gotSlice, func(e int, i int) bool {
				return e == test.gotElementToFind
			})
			assert.Equal(t, test.wantIndex, result)
		})
	}
}

func TestSome(t *testing.T) {
	tests := map[string]struct {
		gotSlice   []int
		gotElement int
		wantSome   bool
	}{
		"contains some": {
			gotSlice:   []int{1, 2, 3},
			gotElement: 3,
			wantSome:   true,
		},
		"does not contain": {
			gotSlice:   []int{1, 2, 3},
			gotElement: 4,
			wantSome:   false,
		},
		"does not contain with empty slice": {
			gotSlice:   []int{},
			gotElement: 4,
			wantSome:   false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			contains := slice.Some(test.gotSlice, func(e int, i int) bool {
				return e == test.gotElement
			})
			assert.Equal(t, test.wantSome, contains)
		})
	}
}

func TestA(t *testing.T) {
	fmt.Println(len("UUDHR5VA3OV75SMH7MNVNWM6WUL4AC2VADLT2JQZNDQVX73B6QQ7U7BFA"))
	fmt.Println(len("UUDHR5VA3OV75SMH7MNVNWM6WUL4AC2VADLT2JQZNDQVX73B6QQ7U7BFA"))
}
