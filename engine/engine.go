package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pipawoz/go_genetic_algorithm/genetics"
	"github.com/pipawoz/go_genetic_algorithm/utils"
)

var Ebiten *ebiten.Image
var GameInstance *Game

// var Genetic genetics.GeneticBox
var GameWidth int
var GameHeight int

type Game struct {
	world        *World
	pixels       []byte
	CurrentLevel int
	Walls        []utils.Obstacle
	MoveLimit    int
	Goal         utils.Vector
	Counter      int
}

// func (g *Game) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "Hello, World!")
// }

const (
	screenWidth  = 1280
	screenHeight = 720
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	// Your existing update code...
	g.world.Update()
	genetics.Genetic.Update()

	// for _, individual := range genetics.Genetic.Population {
	// 	if individual.IsAlive {
	// individual.CheckCollision(g.Walls)
	// individual.Update(g.Counter)

	// Draw the individual
	// individual.Draw(screen)
	// 	}
	// }

	fmt.Println("Update")

	return nil
}

// func Init() {
// 	GameInstance = &Game{}
// }
