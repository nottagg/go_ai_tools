package minimax

import (
	"context"
	"testing"
)

// Mock implementation of GameState for testing
type MockGameState struct {
	moves []interface{}
	value int
}

func (m *MockGameState) Evaluate() int {
	return m.value
}

func (m *MockGameState) GetPossibleMoves() []interface{} {
	return m.moves
}

func (m *MockGameState) MakeMove(move interface{}) GameState {
	newMoves := make([]interface{}, 0, len(m.moves)-1)
	for _, mv := range m.moves {
		if mv != move {
			newMoves = append(newMoves, mv)
		}
	}
	return &MockGameState{moves: newMoves, value: m.value + 1}
}

func (m *MockGameState) UndoMove(move interface{}) {}

func (m *MockGameState) IsTerminal() bool {
	return len(m.moves) == 0
}

func TestMinimax(t *testing.T) {
	state := &MockGameState{
		moves: []interface{}{"move1", "move2"},
		value: 0,
	}

	ctx := context.Background()
	score, timedOut := Minimax(ctx, state, 3, true, -1000, 1000)

	if timedOut {
		t.Error("Minimax timed out unexpectedly")
	}
	if score != 2 { // Expected score based on mock implementation
		t.Errorf("Expected score 2, got %d", score)
	}
}
