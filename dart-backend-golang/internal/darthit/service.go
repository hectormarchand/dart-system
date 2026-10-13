package darthit

import (
	"context"
	"encoding/json"
	"log"

	"dart-backend-golang/internal/game"
	"dart-backend-golang/internal/gamebusiness"
	"dart-backend-golang/internal/sse"
	"dart-backend-golang/internal/types"
)

type Service struct {
	gameService *game.Service
	business    *gamebusiness.Service
	broker      *sse.Broker
}

func NewService(gameService *game.Service, business *gamebusiness.Service) *Service {
	return &Service{
		gameService: gameService,
		business:    business,
		broker:      sse.NewBroker(),
	}
}

func (s *Service) Broker() *sse.Broker {
	return s.broker
}

func (s *Service) AddDartHit(ctx context.Context, dartHitDto types.DartHitDto) (*types.DartGame, error) {
	// if game active, update game
	currentActiveGame, err := s.gameService.GetActiveGame(ctx)
	if err != nil {
		return nil, err
	}

	if currentActiveGame != nil {
		currentActiveGame = s.updateGameAfterDartHit(dartHitDto, currentActiveGame)
		if _, err := s.gameService.UpdateGame(ctx, currentActiveGame); err != nil {
			return nil, err
		}
		s.gameService.PushGameToClients(currentActiveGame)
	}

	s.pushDartHitToClients(dartHitDto)

	return currentActiveGame, nil
}

func (s *Service) updateGameAfterDartHit(dartHitDto types.DartHitDto, dartGame *types.DartGame) *types.DartGame {
	currentPlayer := &dartGame.Players[dartGame.CurrentPlayerIndex]
	currentPlayer.NbOfDartsThrownThisRound++

	// Update player ppd and player score
	currentPlayer.Score -= dartHitDto.EstimatedPoints
	currentPlayer.ScoreThisRound += dartHitDto.EstimatedPoints
	currentPlayer.ScoreIncludingBust += dartHitDto.EstimatedPoints
	currentPlayer.NbOfDartsThrownTotal++
	currentPlayer.Ppd = float64(currentPlayer.ScoreIncludingBust) / float64(currentPlayer.NbOfDartsThrownTotal)

	// Update the game in general
	return s.business.ExecuteOneStepInGame(dartGame)
}

func (s *Service) pushDartHitToClients(dartHit types.DartHitDto) {
	data, err := json.Marshal(dartHit)
	if err != nil {
		log.Println("Error marshalling dart hit for SSE push:", err)
		return
	}
	s.broker.Publish(string(data))
}
