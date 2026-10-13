// Package gamebusiness holds the dart-rules state machine: turn order, round
// progression and win detection. It is kept free of persistence concerns so it
// can be reused by both the game and dart-hit services.
package gamebusiness

import "dart-backend-golang/internal/types"

const (
	throwsPerRound = 3
	totalRounds    = 15
)

type Service struct{}

func New() *Service {
	return &Service{}
}

// ExecuteOneStepInGame moves the game from one state to the next based on
// dart rules: detect a winner, end a round, or advance to the next player.
func (s *Service) ExecuteOneStepInGame(game *types.DartGame) *types.DartGame {
	currentPlayer := &game.Players[game.CurrentPlayerIndex]

	if currentPlayer.Score == 0 {
		winner := *currentPlayer
		game.Winner = &winner
		game.Active = false
		return game
	}

	if isEndOfRoundForPlayer(currentPlayer) {
		if isEndOfCurrentRound(game) {
			if game.CurrentRound == totalRounds {
				game.Active = false
				return game
			}

			game.CurrentRound++
			for i := range game.Players {
				game.Players[i].NbOfDartsThrownThisRound = 0
				game.Players[i].ScoreThisRound = 0
			}
		}

		game.CurrentPlayerIndex = (game.CurrentPlayerIndex + 1) % len(game.Players)
	}

	return game
}

func isEndOfRoundForPlayer(player *types.DartPlayer) bool {
	return player.NbOfDartsThrownThisRound >= throwsPerRound
}

func isEndOfCurrentRound(game *types.DartGame) bool {
	return game.CurrentPlayerIndex == len(game.Players)-1
}
