package game

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"dart-backend-golang/internal/gamebusiness"
	"dart-backend-golang/internal/sse"
	"dart-backend-golang/internal/types"
)

var ErrInvalidGameType = errors.New("invalid game type")

type Service struct {
	collection *mongo.Collection
	business   *gamebusiness.Service
	broker     *sse.Broker
}

func NewService(db *mongo.Database, business *gamebusiness.Service) *Service {
	return &Service{
		collection: db.Collection("games"),
		business:   business,
		broker:     sse.NewBroker(),
	}
}

func (s *Service) Broker() *sse.Broker {
	return s.broker
}

func (s *Service) GetAllGames(ctx context.Context) ([]types.DartGame, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	games := []types.DartGame{}
	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}
	return games, nil
}

func (s *Service) CreateGame(ctx context.Context, dto types.CreateDartGameDto) (*types.DartGame, error) {
	// Unactive all other games
	allGames, err := s.GetAllGames(ctx)
	if err != nil {
		return nil, err
	}
	for i := range allGames {
		allGames[i].Active = false
		if _, err := s.UpdateGame(ctx, &allGames[i]); err != nil {
			return nil, err
		}
	}

	totalScoreToDo, err := gameTypeToScore(dto.GameType)
	if err != nil {
		return nil, err
	}

	players := make([]types.DartPlayer, 0, len(dto.Players))
	for _, playerDto := range dto.Players {
		players = append(players, types.DartPlayer{
			Name:  playerDto.Name,
			Score: totalScoreToDo,
		})
	}

	totalRound := 15
	if dto.TotalRound != nil {
		totalRound = *dto.TotalRound
	}

	now := time.Now()
	dartGame := types.DartGame{
		Players:            players,
		GameType:           dto.GameType,
		TotalRound:         totalRound,
		CurrentRound:       1,
		Active:             true,
		CreatedDate:        fmt.Sprintf("%d-%d-%d %d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute()),
		CurrentPlayerIndex: 0,
	}

	result, err := s.collection.InsertOne(ctx, dartGame)
	if err != nil {
		return nil, err
	}
	dartGame.ID = result.InsertedID.(primitive.ObjectID)

	return &dartGame, nil
}

func (s *Service) GetActiveGame(ctx context.Context) (*types.DartGame, error) {
	cursor, err := s.collection.Find(ctx, bson.M{"active": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	activeGames := []types.DartGame{}
	if err := cursor.All(ctx, &activeGames); err != nil {
		return nil, err
	}

	if len(activeGames) > 1 {
		log.Println("WARNING : More than one game is active")
	}

	if len(activeGames) == 0 {
		return nil, nil
	}
	return &activeGames[0], nil
}

// UpdateGame persists dartGame, replacing the document with the matching _id.
func (s *Service) UpdateGame(ctx context.Context, dartGame *types.DartGame) (*types.DartGame, error) {
	_, err := s.collection.ReplaceOne(ctx, bson.M{"_id": dartGame.ID}, dartGame)
	if err != nil {
		return nil, err
	}
	return dartGame, nil
}

func (s *Service) PatchGame(ctx context.Context, currentActiveGame *types.DartGame, patch types.PatchGameDto) (*types.DartGame, error) {
	switch patch.Type {
	case "next-player":
		currentActiveGame = s.patchGameNextPlayer(currentActiveGame)
	case "fix-score":
		// TODO: patch a specific player's score
	}

	return s.UpdateGame(ctx, currentActiveGame)
}

func (s *Service) patchGameNextPlayer(game *types.DartGame) *types.DartGame {
	currentPlayer := &game.Players[game.CurrentPlayerIndex]

	// Update player ppd and player score
	currentPlayer.NbOfDartsThrownThisRound = 3
	currentPlayer.NbOfDartsThrownTotal = int(math.Ceil(float64(currentPlayer.NbOfDartsThrownTotal+1)/3)) * 3
	currentPlayer.Ppd = float64(currentPlayer.ScoreIncludingBust) / float64(currentPlayer.NbOfDartsThrownTotal)

	return s.business.ExecuteOneStepInGame(game)
}

// PushGameToClients sends the game to every client subscribed to the SSE stream.
func (s *Service) PushGameToClients(dartGame *types.DartGame) {
	data, err := json.Marshal(dartGame)
	if err != nil {
		log.Println("Error marshalling game for SSE push:", err)
		return
	}
	s.broker.Publish(string(data))
}

func gameTypeToScore(gameType types.DartGameType) (int, error) {
	switch gameType {
	case types.Game301:
		return 301, nil
	case types.Game501:
		return 501, nil
	case types.Game701:
		return 701, nil
	default:
		return 0, ErrInvalidGameType
	}
}
