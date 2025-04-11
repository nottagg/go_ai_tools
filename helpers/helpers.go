package helpers

// Contains helper functions used across the toolset
// GetNeighbors takes the overall graph and returns neighbors of the node
// T is a generic for the graph type
type Node[graphType any] interface {
	GetParent() Node[graphType]
	SetParent(Node[graphType])
	GetNeighbors(graph graphType) []Node[graphType]
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

// Given a node with a "Parent" pointer, build a path and return as a slice of nodes
func BuildReturnPath[T any](node Node[T]) []Node[T] {
	path := []Node[T]{}
	for node != nil {
		path = append(path, node)
		node = node.GetParent()
	}
	return path
}
