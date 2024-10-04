// main.go
package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pipawoz/go_genetic_algorithm/internal/engine"
	"github.com/pipawoz/go_genetic_algorithm/internal/utils"
)

func main() {
	utils.LoadGameSettings()
	utils.LoadGeneticSettings()

	populationSize := utils.DNASettings.PopulationSize
	maxGenerations := utils.DNASettings.MaxGenerations
	showTrails := utils.Settings.PrintTrace

	rand.New(rand.NewSource(rand.Int63()))

	ebiten.SetWindowSize(utils.GameWidth, utils.GameHeight)
	ebiten.SetWindowTitle("Go - Genetic Algorithm Maze")

	ebiten.SetTPS(60)

	game := engine.NewGame(populationSize, maxGenerations, showTrails)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
