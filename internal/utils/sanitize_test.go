package utils_test

import (
	"reflect"
	"testing"

	"github.com/ozonva/ova-account-api/internal/utils"
)

// Stop list: "bad", "word", "shift".
func TestFilterWords(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"Hello", "bad", "word"}, []string{"Hello"}},
		{nil, []string{}},
		{[]string{}, []string{}},
		{[]string{"Hello", "Bad", "word"}, []string{"Hello", "Bad"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := utils.FilterWords(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastFilterWords(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"Accept credit cards", "All new", "As seen on", "Bargain", "Ad"}, []string{}},
		{nil, []string{}},
		{[]string{}, []string{}},
		{[]string{"Hello", "Bad", "zzz", "Trial"}, []string{"Hello", "Bad", "zzz"}},
		{[]string{"", ""}, []string{"", ""}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := utils.FastFilterWords(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FastFilterWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
