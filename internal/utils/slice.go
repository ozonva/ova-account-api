package utils

import (
	"errors"
)

// ChunkSliceInt chunks slice on slices of the given size.
func ChunkSliceInt(s []int, size int) ([][]int, error) {
	if size < 1 {
		return nil, errors.New("the slice chunk size less than 1")
	}

	out := make([][]int, 0, (len(s)+size-1)/size)
	l := len(s)
	var i int
	for i = 0; i <= l-size; i += size {
		out = append(out, s[i:i+size])
	}

	if i < l {
		out = append(out, s[i:])
	}

	return out, nil
}

// FilterSliceString filters a slice of strings by the filter predicate.
func FilterSliceString(s []string, filter func(string) bool) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}
