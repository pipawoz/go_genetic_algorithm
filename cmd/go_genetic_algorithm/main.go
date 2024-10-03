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

const (
	screenWidth  = 1280
	screenHeight = 720
)

type Game struct {
	geneticAlgorithm  *genetics.GeneticBox
	currentGeneration int
	maxGenerations    int
	counter           int
	moveLimit         int
	walls             []utils.Obstacle
	level             int
	showTrails        bool
	trailImage        *ebiten.Image // Imagen para dibujar las trazas
}

// NewGame Creates a new game.
func NewGame(populationSize int, maxGenerations int, showTrails bool) *Game {
	game := &Game{
		geneticAlgorithm:  genetics.NewGeneticBox(populationSize),
		currentGeneration: 1,
		maxGenerations:    maxGenerations,
		counter:           0,
		level:             4,
		showTrails:        showTrails,
		trailImage:        ebiten.NewImage(screenWidth, screenHeight),
	}
	game.moveLimit, game.walls = game.selectLevel(game.level)
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
		// // Calcular el fitness de cada individuo
		// for i := range g.geneticAlgorithm.Population {
		// 	g.geneticAlgorithm.Population[i].CalculateFitness()
		// }

		// Avanzar a la siguiente generación
		g.geneticAlgorithm.NextGeneration()

		fmt.Println("*** Generationn ***")
		fmt.Println("Generation: ", g.currentGeneration)
		fmt.Println("Avg Distance: ", g.geneticAlgorithm.AvgDistance)
		fmt.Println("Avg Fitness: ", g.geneticAlgorithm.AvgFitness)

		// Reiniciar el contador y aumentar la generación actual
		g.counter = 0
		g.currentGeneration++

		g.trailImage.Clear()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if !g.showTrails {
		g.trailImage.Clear()
	}

	// Dibujar fondo
	screen.Fill(color.RGBA{0, 0, 0, 255})

	var goalSize = 40
	var goalX = screenWidth - goalSize - 10
	var goalY = screenHeight/2 - goalSize/2

	utils.Settings.GoalX = goalX
	utils.Settings.GoalY = goalY

	// Dibujar objetivo (cuadrado)
	goalImg := ebiten.NewImage(goalSize, goalSize)
	goalImg.Fill(color.RGBA{0, 255, 0, 255}) // Color verde
	goalOpts := &ebiten.DrawImageOptions{}
	goalOpts.GeoM.Translate(float64(goalX), float64(goalY))
	screen.DrawImage(goalImg, goalOpts)

	// goalRect := ebiten.NewImage(10, screenHeight)
	// goalRect.Fill(color.RGBA{0, 255, 0, 255})
	// opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(float64(screenWidth-10), 0)
	// screen.DrawImage(goalRect, opts)

	// Dibujar obstáculos
	for _, wall := range g.walls {
		wallImg := ebiten.NewImage(wall.Width, wall.Height)
		wallImg.Fill(color.RGBA{255, 255, 255, 255})
		wallOpts := &ebiten.DrawImageOptions{}
		wallOpts.GeoM.Translate(float64(wall.X), float64(wall.Y))
		screen.DrawImage(wallImg, wallOpts)
	}

	// Dibujar individuos
	for i := range g.geneticAlgorithm.Population {
		individual := &g.geneticAlgorithm.Population[i]
		if individual.IsAlive {
			individual.Draw(g.trailImage)
		}
	}

	screen.DrawImage(g.trailImage, nil)

	// Dibujar información
	msg := fmt.Sprintf("Generación: %d/%d", g.currentGeneration, g.maxGenerations)
	ebitenutil.DebugPrintAt(screen, msg, 10, 10)

	// Mostrar el contador de frames
	frameMsg := fmt.Sprintf("Frame: %d", g.counter)
	ebitenutil.DebugPrintAt(screen, frameMsg, 10, 30)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// selectLevel configura el nivel y retorna el límite de movimientos y las paredes.
func (g *Game) selectLevel(currentLevel int) (int, []utils.Obstacle) {
	var moveLimit int
	var walls []utils.Obstacle

	switch currentLevel {
	case 1:
		// Nivel 1
		moveLimit = 350
	case 2:
		// Nivel 2
		moveLimit = 400
		walls = []utils.Obstacle{
			{X: 500, Y: 150, Width: 20, Height: 420},
		}
	case 3:
		// Nivel 3
		moveLimit = 500
		walls = []utils.Obstacle{
			{X: 350, Y: 200, Width: 20, Height: 320},
			{X: 750, Y: 200, Width: 20, Height: 320},
			{X: 550, Y: 0, Width: 20, Height: 200},
			{X: 550, Y: 520, Width: 20, Height: 200},
		}
	case 4:
		// Nivel 4
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
		// Nivel 5
		moveLimit = 700
		walls = []utils.Obstacle{
			{X: 200, Y: 300, Width: 20, Height: 420},
			{X: 500, Y: 0, Width: 20, Height: 350},
			{X: 800, Y: 300, Width: 20, Height: 420},
			{X: 1100, Y: 0, Width: 20, Height: 350},
		}
	default:
		// Nivel por defecto
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

	populationSize := 50
	maxGenerations := 200
	showTrails := true

	rand.New(rand.NewSource(rand.Int63()))

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Genetic Algorithm Maze")

	ebiten.SetTPS(200)

	game := NewGame(populationSize, maxGenerations, showTrails)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
