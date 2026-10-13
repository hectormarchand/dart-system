package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type DartGameType string

const (
	Game301 DartGameType = "301"
	Game501 DartGameType = "501"
	Game701 DartGameType = "701"
)

type DartPlayer struct {
	Name                     string  `json:"name" bson:"name"`
	Score                    int     `json:"score" bson:"score"`
	ScoreIncludingBust       int     `json:"scoreIncludingBust" bson:"scoreIncludingBust"`
	ScoreThisRound           int     `json:"scoreThisRound" bson:"scoreThisRound"`
	Ppd                      float64 `json:"ppd" bson:"ppd"`
	NbOfDartsThrownThisRound int     `json:"nbOfDartsThrownThisRound" bson:"nbOfDartsThrownThisRound"`
	NbOfDartsThrownTotal     int     `json:"nbOfDartsThrownTotal" bson:"nbOfDartsThrownTotal"`
}

type DartGame struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Players            []DartPlayer       `json:"players" bson:"players"`
	Winner             *DartPlayer        `json:"winner,omitempty" bson:"winner,omitempty"`
	GameType           DartGameType       `json:"gameType" bson:"gameType"`
	TotalRound         int                `json:"totalRound" bson:"totalRound"`
	CurrentRound       int                `json:"currentRound" bson:"currentRound"`
	Active             bool               `json:"active" bson:"active"`
	CreatedDate        string             `json:"createdDate" bson:"createdDate"`
	CurrentPlayerIndex int                `json:"currentPlayerIndex" bson:"currentPlayerIndex"`
}

// DTOs

type CreatePlayerDto struct {
	Name string `json:"name"`
}

type CreateDartGameDto struct {
	Players    []CreatePlayerDto `json:"players"`
	GameType   DartGameType      `json:"gameType"`
	TotalRound *int              `json:"totalRound"`
}

type DartHitDto struct {
	R               float64 `json:"r"`
	Theta           float64 `json:"theta"` // Theta in rad
	EstimatedPoints int     `json:"estimatedPoints"`
}

type PatchGameDto struct {
	Type string `json:"type"` // "next-player" | "fix-score"
	Fix  *int   `json:"fix,omitempty"`
}
