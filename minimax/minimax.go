package minimax

import (
	"context"
)

// GameState is the interface users must implement for their game.
type GameState interface {
	Evaluate() int
	GetPossibleMoves() []interface{}
	MakeMove(move interface{}) GameState
	UndoMove(move interface{})
	IsTerminal() bool
}

// Minimax performs the minimax algorithm with alpha-beta pruning.
// It returns the best score for the maximizing player and a boolean indicating if the operation was interrupted.
// The context can be used to implement a timeout or cancellation.
// The depth parameter controls how deep the algorithm will search in the game tree.
// The alpha and beta parameters are used for alpha-beta pruning.
// The isMaximizing parameter indicates whether the current player is the maximizing player.
// If the context is done, it returns the evaluation of the current state and a true boolean indicating a timeout.
// If the depth is zero or the state is terminal, it returns the evaluation of the current state and false indicating no timeout.
func Minimax(
	ctx context.Context,
	state GameState,
	depth int,
	isMaximizing bool,
	alpha int,
	beta int,
) (int, bool) {
	select {
	case <-ctx.Done():
		return state.Evaluate(), true // Timer expired or cancelled
	default:
	}

	if depth == 0 || state.IsTerminal() {
		return state.Evaluate(), false
	}

	moves := state.GetPossibleMoves()
	if isMaximizing {
		bestScore := -1 << 31
		for _, move := range moves {
			newState := state.MakeMove(move)
			score, timedOut := Minimax(ctx, newState, depth-1, false, alpha, beta)
			state.UndoMove(move)
			if timedOut {
				return 0, true
			}
			bestScore = max(bestScore, score)
			alpha = max(alpha, bestScore)
			if beta <= alpha {
				break // Beta cut-off
			}
		}
		return bestScore, false
	} else {
		bestScore := 1 << 31
		for _, move := range moves {
			newState := state.MakeMove(move)
			score, timedOut := Minimax(ctx, newState, depth-1, true, alpha, beta)
			state.UndoMove(move)
			if timedOut {
				return 0, true
			}
			bestScore = min(bestScore, score)
			beta = min(beta, bestScore)
			if beta <= alpha {
				break // Alpha cut-off
			}
		}
		return bestScore, false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
