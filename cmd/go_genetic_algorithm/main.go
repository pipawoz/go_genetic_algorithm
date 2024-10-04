// main.go
package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pipawoz/go_genetic_algorithm/internal/genetics"
	"github.com/pipawoz/go_genetic_algorithm/internal/utils"
)

// Game represents the game state.
type Game struct {
	geneticAlgorithm  *genetics.GeneticBox
	currentGeneration int
	maxGenerations    int
	counter           int
	moveLimit         int
	walls             []utils.Obstacle
	level             int
	showTrails        bool
	trailImage        *ebiten.Image
	avgFitness        float64
	avgFitnessOld     float64
	fitnessHistory    []float64
}

// NewGame Creates a new game.
func NewGame(populationSize int, maxGenerations int, showTrails bool) *Game {
	game := &Game{
		geneticAlgorithm:  genetics.NewGeneticBox(populationSize),
		currentGeneration: 1,
		maxGenerations:    maxGenerations,
		counter:           0,
		level:             utils.Settings.CurrentLevel,
		showTrails:        showTrails,
		trailImage:        ebiten.NewImage(utils.GameWidth, utils.GameHeight),
	}
	game.moveLimit, game.walls = game.SelectLevel(game.level)
	return game
}

// Update Updates the game state.
func (g *Game) Update() error {
	if g.currentGeneration > g.maxGenerations {
		return fmt.Errorf("Max generations reached")
	}

	allDeadOrWon := true
	for i := range g.geneticAlgorithm.Population {
		individual := &g.geneticAlgorithm.Population[i]
		if individual.IsAlive && !individual.Won {
			allDeadOrWon = false
			individual.Update(g.counter)
			individual.CheckCollision(g.walls)
		}
	}

	g.counter++

	if allDeadOrWon || g.counter > g.moveLimit {
		g.geneticAlgorithm.NextGeneration()

		avgFitnessCurrent := g.geneticAlgorithm.AvgFitness
		avgDistance := g.geneticAlgorithm.AvgDistance

		fmt.Println("")
		fmt.Println("*** Generation ***")
		fmt.Println("Generation: ", g.currentGeneration)
		fmt.Println("Avg Distance: ", avgDistance)
		fmt.Println("Avg Fitness: ", avgFitnessCurrent)

		if g.currentGeneration > 1 {
			percentageChange := ((avgFitnessCurrent - g.avgFitnessOld) / g.avgFitnessOld) * 100
			fmt.Printf("Avg Fitness Change: %.2f%%\n", percentageChange)
		}

		g.avgFitnessOld = g.avgFitness
		g.avgFitness = avgFitnessCurrent
		g.fitnessHistory = append(g.fitnessHistory, avgFitnessCurrent)

		g.counter = 0
		g.currentGeneration++
		g.trailImage.Clear()
	}

	return nil
}

// Draw Draws the game state.
func (g *Game) Draw(screen *ebiten.Image) {

	if !g.showTrails {
		g.trailImage.Clear()
	}

	screen.Fill(color.RGBA{0, 0, 0, 255})

	var goalSize = 40
	var goalX = utils.GameWidth - goalSize - 10
	var goalY = utils.GameHeight/2 - goalSize/2

	// Draw goal
	goalImg := ebiten.NewImage(goalSize, goalSize)
	goalImg.Fill(color.RGBA{0, 255, 0, 255})
	goalOpts := &ebiten.DrawImageOptions{}
	goalOpts.GeoM.Translate(float64(goalX), float64(goalY))
	screen.DrawImage(goalImg, goalOpts)

	// Draw walls
	for _, wall := range g.walls {
		wallImg := ebiten.NewImage(wall.Width, wall.Height)
		wallImg.Fill(color.RGBA{255, 255, 255, 255})
		wallOpts := &ebiten.DrawImageOptions{}
		wallOpts.GeoM.Translate(float64(wall.X), float64(wall.Y))
		screen.DrawImage(wallImg, wallOpts)
	}

	// Draw individuals
	for i := range g.geneticAlgorithm.Population {
		individual := &g.geneticAlgorithm.Population[i]
		if individual.IsAlive {
			individual.Draw(g.trailImage)
		}
	}

	screen.DrawImage(g.trailImage, nil)

	// Draw information
	msg := fmt.Sprintf("Generación: %d/%d", g.currentGeneration, g.maxGenerations)
	ebitenutil.DebugPrintAt(screen, msg, 10, 10)

	frameMsg := fmt.Sprintf("Frame: %d", g.counter)
	ebitenutil.DebugPrintAt(screen, frameMsg, 10, 30)

}

// Layout Layout the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return utils.GameWidth, utils.GameHeight
}

// SelectLevel configura el nivel y retorna el límite de movimientos y las paredes.
func (g *Game) SelectLevel(currentLevel int) (int, []utils.Obstacle) {
	var moveLimit int
	var walls []utils.Obstacle

	switch currentLevel {
	case 1:
		// Level 1
		moveLimit = 350
	case 2:
		// Level 2
		moveLimit = 400
		walls = []utils.Obstacle{
			{X: 500, Y: 150, Width: 20, Height: 420},
		}
	case 3:
		// Level 3
		moveLimit = 500
		walls = []utils.Obstacle{
			{X: 350, Y: 200, Width: 20, Height: 320},
			{X: 750, Y: 200, Width: 20, Height: 320},
			{X: 550, Y: 0, Width: 20, Height: 200},
			{X: 550, Y: 520, Width: 20, Height: 200},
		}
	case 4:
		// Level 4
		moveLimit = 600
		walls = []utils.Obstacle{
			{X: 300, Y: 0, Width: 20, Height: 400},
			{X: 500, Y: 400, Width: 20, Height: 320},
			{X: 730, Y: 0, Width: 20, Height: 310},
			{X: 730, Y: 420, Width: 20, Height: 300},
			{X: 750, Y: 290, Width: 300, Height: 20},
			{X: 750, Y: 420, Width: 300, Height: 20},
		}
	case 5:
		// Level 5
		moveLimit = 700
		walls = []utils.Obstacle{
			{X: 200, Y: 300, Width: 20, Height: 420},
			{X: 500, Y: 0, Width: 20, Height: 350},
			{X: 800, Y: 300, Width: 20, Height: 420},
			{X: 1100, Y: 0, Width: 20, Height: 350},
		}
	default:
		// Default Level
		moveLimit = 350
	}

	g.level = currentLevel
	g.walls = walls
	g.moveLimit = moveLimit

	return moveLimit, walls
}

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

	game := NewGame(populationSize, maxGenerations, showTrails)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
