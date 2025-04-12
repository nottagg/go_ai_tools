package helpers

// Contains helper functions used across the toolset

type Coordinate struct {
	X int
	Y int
}

// MapKeysToSlice takes a generic map and returns the keys as a slice
func MapKeysToSlice[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValuesToSlice takes a generic map and returns the values as a slice
func MapValuesToSlice[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetGridDirections(allowDiagonal bool) [][2]int {
	if allowDiagonal {
		return [][2]int{
			{0, -1},  // Up
			{1, 0},   // Right
			{0, 1},   // Down
			{-1, 0},  // Left
			{-1, -1}, // Up-Left
			{1, -1},  // Up-Right
			{1, 1},   // Down-Right
			{-1, 1},  // Down-Left
		}
	}
	return [][2]int{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}
}
