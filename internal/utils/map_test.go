package utils_test

import (
	"reflect"
	"testing"

	"github.com/ozonva/ova-account-api/internal/utils"
)

func TestInvertMap(t *testing.T) {
	tests := []struct {
		in   map[string]int
		want map[int]string
	}{
		{in: nil, want: map[int]string{}},
		{in: map[string]int{}, want: map[int]string{}},
		{in: map[string]int{"1": 1, "2": 2}, want: map[int]string{1: "1", 2: "2"}},
		// {in: map[string]int{"1": 1, "2": 1}, want: map[int]string{1: "2"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := utils.InvertMap(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
