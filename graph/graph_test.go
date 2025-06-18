package graph

import (
	"reflect"
	"testing"

	"main.go/helpers"
)

func TestGraphDirected(t *testing.T) {
	g := New[string, int]("testGraph", true)
	g.AddNode("A", 1)
	g.AddNode("B", 2)
	g.AddEdge("A", "B", 1)

	if len(g.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", g.LengthNodes())
	}

	if len(g.Edges) != 1 {
		t.Errorf("Expected 1 edge, got %d", g.LengthEdges())
	}

	if g.GetEdgeWeight("A", "B") != 1 {
		t.Errorf("Expected edge weight 1, got %f", g.GetEdgeWeight("A", "B"))
	}
	if g.GetEdgeWeight("B", "A") != -1 {
		t.Errorf("Expected edge weight -1 for non-existent edge, got %f", g.GetEdgeWeight("B", "A"))
	}
}

func TestGraphUndirected(t *testing.T) {
	g := New[string, int]("testGraph", false)
	g.AddNode("A", 1)
	g.AddNode("B", 2)
	g.AddEdge("A", "B", 2)

	if len(g.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", g.LengthNodes())
	}

	if len(g.Edges) != 2 {
		t.Errorf("Expected 2 edge, got %d", g.LengthEdges())
	}

	if g.GetEdgeWeight("A", "B") != 2 {
		t.Errorf("Expected edge weight 2, got %f", g.GetEdgeWeight("A", "B"))
	}
	if g.GetEdgeWeight("B", "A") != 2 {
		t.Errorf("Expected edge weight 2 for undirected edge, got %f", g.GetEdgeWeight("B", "A"))
	}
	if g.GetEdgeWeight("A", "C") != -1 {
		t.Errorf("Expected edge weight -1 for non-existent edge, got %f", g.GetEdgeWeight("A", "C"))
	}
}

func TestGraphFromMatrixOrtho(t *testing.T) {
	matrix := [][]float64{
		{0, 1, -1},
		{1, 5, 3},
		{0, 1, 0},
	}
	g := NewGraphFromMatrix("grid", matrix, false)
	if g == nil {
		t.Errorf("Expected a graph, got nil")
	}
	if g.LengthNodes() != 9 {
		t.Errorf("Expected 9 nodes, got %d", g.LengthNodes())
	}

	if g.LengthEdges() != 20 {
		t.Errorf("Expected 20 edges, got %d", g.LengthEdges())
	}
}

func TestGraphFromMatrixDiag(t *testing.T) {
	matrix := [][]float64{
		{0, 1, -1},
		{1, 5, 3},
		{0, 1, 0},
	}
	g := NewGraphFromMatrix("grid", matrix, true)
	if g == nil {
		t.Errorf("Expected a graph, got nil")
	}
	if g.LengthNodes() != 9 {
		t.Errorf("Expected 9 nodes, got %d", g.LengthNodes())
	}

	if g.LengthEdges() != 34 {
		t.Errorf("Expected 34 edges, got %d", g.LengthEdges())
	}
}

func TestSearch(t *testing.T) {
	matrix := [][]float64{
		{0, 1, 2, -1, 0},
		{2, 0, 0, -1, -1},
		{2, -1, 0, 1, 4},
		{-1, 3, 2, 0, 2},
		{-1, -1, 4, 3, 0},
	}

	// Create a graph from the matrix
	g := NewGraphFromMatrix("testGraph", matrix, false)
	if g == nil {
		t.Error("Expected a graph, got nil")
	}
	// Ensure the graph has 25 nodes
	if g.LengthNodes() != 25 {
		t.Errorf("Expected 25 nodes, got %d", g.LengthNodes())
	}

	// Perform a failed search
	start := helpers.Coordinate{X: 0, Y: 4}
	end := helpers.Coordinate{X: 0, Y: 0}
	g.ResetNodes()
	_, _, BFSerr := g.BFS(start, end)
	g.ResetNodes()
	_, _, DFSSerr := g.DFS(start, end)
	g.ResetNodes()
	_, _, Dijkstraerr := g.Dijkstra(start, end)
	heuristic := func(a, b helpers.Coordinate) float64 {
		return 1
	}
	g.ResetNodes()
	_, _, AStarerr := g.AStar(start, end, heuristic)

	if BFSerr == nil || DFSSerr == nil || Dijkstraerr == nil || AStarerr == nil {
		t.Error("Expected search to fail, but it succeeded")
	}

	// Perform a successful search
	start = helpers.Coordinate{X: 0, Y: 0}
	end = helpers.Coordinate{X: 4, Y: 4}

	//Can't test deterministic path for BFS, just verify it finds a path and the path is shortest possible
	g.ResetNodes()
	path, visited, err := g.BFS(start, end)
	if err != nil {
		t.Errorf("BFS failed: %v", BFSerr)
	}
	if len(path) != 9 {
		t.Errorf("BFS found path of length %d, expected 8", len(path))
	}
	if len(visited) != 17 {
		t.Errorf("BFS visited %d nodes, expected 17", len(visited))
	}

	//Can't test path or visited count for DFS, just verify it finds a path
	g.ResetNodes()
	_, _, err = g.DFS(start, end)
	if err != nil {
		t.Errorf("DFS failed: %v", DFSSerr)
	}

	ExpectedPath := []helpers.Coordinate{
		{X: 0, Y: 0},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
		{X: 3, Y: 4},
		{X: 4, Y: 4},
	}
	g.ResetNodes()
	path, Dvisited, err := g.Dijkstra(start, end)
	if err != nil {
		t.Errorf("Dijkstra failed: %v", err)
	}
	if !reflect.DeepEqual(path, ExpectedPath) {
		t.Errorf("Dijkstra found path %v, expected %v", path, ExpectedPath)
	}

	heuristic = helpers.EuclideanDistance
	g.ResetNodes()
	path, Avisited, err := g.AStar(start, end, heuristic)
	if err != nil {
		t.Errorf("A* failed: %v", AStarerr)
	}
	if !reflect.DeepEqual(path, ExpectedPath) {
		t.Errorf("A* found path %v, expected %v", path, ExpectedPath)
	}
	if len(Dvisited) <= len(Avisited) {
		t.Errorf("A* visited %d nodes, expected less than Dijkstra's %d", len(Avisited), len(Dvisited))
	}

}
