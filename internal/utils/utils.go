package utils

import (
	"encoding/json"
	"os"
	"sync"
)

// Vector represents a 2D vector.
type Vector struct {
	X float32
	Y float32
}

// GameSettings represents the settings for the game.
type GameSettings struct {
	PrintTrace   bool   `json:"printTrace"`
	CurrentLevel int    `json:"currentLevel"`
	OutputFile   string `json:"outputFile"`
	SimulateOnly bool   `json:"simulateOnly"`
}

// GeneticSettings represents the settings for the genetic algorithm.
type GeneticSettings struct {
	Iterations     int     `json:"iterations"`
	MaxGenerations int     `json:"maxGenerations"`
	PopulationSize int     `json:"populationSize"`
	MutationRate   float64 `json:"mutationRate"`
	CrossoverRate  float64 `json:"crossoverRate"`
}

var (
	// Settings represents the game settings.
	Settings GameSettings
	// DNASettings represents the genetic settings.
	DNASettings GeneticSettings
	once        sync.Once
	geneticOnce sync.Once
)

// LoadGameSettings loads the game settings from the settings.json file.
func LoadGameSettings() (GameSettings, error) {
	var err error
	once.Do(func() {
		data, readErr := os.ReadFile("configs/settings.json")
		if readErr != nil {
			err = readErr
			return
		}

		unmarshalErr := json.Unmarshal(data, &Settings)
		if unmarshalErr != nil {
			err = unmarshalErr
			return
		}
	})

	return Settings, err
}

// LoadGeneticSettings loads the genetic settings from the genetic_settings.json file.
func LoadGeneticSettings() (GeneticSettings, error) {
	var err error
	geneticOnce.Do(func() {
		data, readErr := os.ReadFile("configs/genetic_settings.json")
		if readErr != nil {
			err = readErr
			return
		}

		unmarshalErr := json.Unmarshal(data, &DNASettings)
		if unmarshalErr != nil {
			err = unmarshalErr
			return
		}
	})

	return DNASettings, err
}

// Obstacle represents an object that the player must avoid.
type Obstacle struct {
	X      int
	Y      int
	Width  int
	Height int
}
