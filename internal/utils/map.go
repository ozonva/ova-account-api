package utils

// InvertMap inverts the key-value pair.
func InvertMap(m map[string]int) map[int]string {
	out := make(map[int]string, len(m))

	for key, value := range m {
		out[value] = key // Unexpected behavior, if the values are repeated
	}

	return out
}
