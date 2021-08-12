package utils_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/ozonva/ova-account-api/internal/utils"
)

func TestChunkIntSlice(t *testing.T) {
	tests := []struct {
		in          []int
		size        int
		want        [][]int
		expectedErr bool
	}{
		{nil, 10, [][]int{}, false},
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}, false},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}, false},
		{[]int{1, 2, 3}, 5, [][]int{{1, 2, 3}}, false},
		{[]int{1, 2, 3}, 0, nil, true},
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := utils.ChunkSliceInt(tt.in, tt.size)

			checkErr(t, "ChunkIntSlice()", err, tt.expectedErr)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSliceString(t *testing.T) {
	even := func(s string) bool {
		i, _ := strconv.Atoi(s)
		return i%2 == 0
	}

	tests := []struct {
		s      []string
		filter func(string) bool
		want   []string
	}{
		{[]string{"1", "3", "4"}, even, []string{"4"}},
		{[]string{}, even, []string{}},
		{nil, even, []string{}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := utils.FilterSliceString(tt.s, tt.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterSliceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func checkErr(t *testing.T, funcName string, err error, expected bool) {
	t.Helper()
	if expected && err == nil {
		t.Errorf("%s got nil, want error", funcName)
	}

	if !expected && err != nil {
		t.Errorf("%s got unexpected error: %v", funcName, err)
	}
}
