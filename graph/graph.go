package graph

//go interepretation of https://networkx.org/documentation/stable/_modules/networkx/classes/graph.html#Graph

type Graph[K comparable, V any] struct {
	nodes     map[K]*Node[K, V]
	edges     map[K]map[K]int
	name      string
	graphType GraphType
}
type Node[K comparable, V any] struct {
	id    K
	value V
	x     int
	y     int
}

type GraphType int

const (
	Grid GraphType = iota
	Directed
	Undirected
)

// For grid type, graph uniqueness is based on x,y coordinates
// For directed and undirected, graph uniqueness is based on the value of the node
var graphTypes = map[string]GraphType{
	"grid":       Grid,
	"directed":   Directed,
	"undirected": Undirected,
}

// GraphType is a string that represents the type of graph
// It can be "grid", "directed", or "undirected"
// Returns a pointer to a new graph object
// Examples
// g := New("MyGraph", "grid"))
func New[K comparable, V any](n string, gt GraphType) *Graph[K, V] {
	return &Graph[K, V]{
		nodes:     make(map[K]*Node[K, V]),
		edges:     make(map[K]map[K]int),
		name:      n,
		graphType: gt,
	}
}

// String identifier of the graph
func (g *Graph[K, V]) Name(s string) {
	g.name = s
}

// Returns True if n is a node, False otherwise
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// fmt.Println(g.Contains("A")) // true
// fmt.Println(g.Contains("B")) // false
func (g *Graph[K, V]) Contains(n K) bool {
	_, exists := g.nodes[n]
	return exists
}

// Returns the number of nodes in the graph
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// g.AddNode("B")
// fmt.Println(g.Length()) // 2
func (g *Graph[K, V]) Length() int {
	return len(g.nodes)
}

// Adds a node to the graph. Must include an x,y coordinate for
// calculating euclidean/manhattan distance in A*
// K is a unique indentifier while V is the value
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", 0, 0)
// g.AddNode("B", 1, 1)
// fmt.Println(g.Length()) // 2
func (g *Graph[K, V]) AddNode(k K, v V, x, y int) {
	// Add a node to the graph
	if _, exists := g.nodes[k]; !exists {
		g.nodes[k] = &Node[K, V]{
			value: v,
			x:     x,
			y:     y,
			id:    k,
		}
	}
}

// Adds multiple nodes to the graph.
// For using later on with converting 2d grid to graph
// Examples
// g := New("MyGraph", false)
// g.AddNodesFrom([]Node{"A", "B", "C"})
// fmt.Println(g.Length()) // 3
func (g *Graph[K, V]) AddNodesFrom(nodes []Node[K, V]) {
	// Add multiple nodes to the graph
	for _, node := range nodes {
		g.AddNode(node.id, node.value, node.x, node.y)
	}
}

// Removes a node from the graph
// If key is simply the value or coordinate pair pass that
// If the node is not in the graph, do nothing
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", "A",0, 0)
// g.AddNode("B", "B",1, 1)
// g.RemoveNode("A")
// fmt.Println(g.Length()) // 1
func (g *Graph[K, V]) RemoveNode(k K) {
	// Remove a node from the graph
	if _, exists := g.nodes[k]; exists {
		delete(g.nodes, k)
	}
}

func (g *Graph[K]) RemoveNodesFrom(nodes []K) {
	// Remove multiple nodes from the graph
	for _, node := range nodes {
		g.RemoveNode(node)
	}
}

func (g *Graph) AddEdge(n1, n2 string) {
	// Add an edge between two nodes
}

func (g *Graph) RemoveNode(n string) {
	// Remove a node from the graph
}

func (g *Graph) RemoveEdge(n1, n2 string) {
	// Remove an edge between two nodes
}
