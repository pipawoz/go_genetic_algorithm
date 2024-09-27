package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pipawoz/go_genetic_algorithm/genetics"
	"github.com/pipawoz/go_genetic_algorithm/population"
	"github.com/pipawoz/go_genetic_algorithm/utils"
)

var Ebiten *ebiten.Image
var GameInstance *Game

// var Genetic genetics.GeneticBox
var GameWidth int
var GameHeight int

type World struct {
	area   []bool
	width  int
	heigth int
}

func NewWorld(width, height int) *World {
	w := &World{
		area:   make([]bool, width*height),
		width:  width,
		heigth: height,
	}
	w.init(100)
	return w
}

func (w *World) init(n int) {
	for i := 0; i < n; i++ {
		w.area[i] = true
	}
}

func (w *World) Update() {
	fmt.Println("UpdateWorld")
	genetics.Genetic.Update()
	// Your existing update code...
}

func (w *World) Draw(pix []byte) {
	for i, v := range w.area {
		if v {
			pix[i*4] = 0
			pix[i*4+1] = 0
			pix[i*4+2] = 0
			pix[i*4+3] = 255
		}
	}
}

type Game struct {
	world *World
	// geneticBox *genetics.GeneticBox
	// box *population.Box
	dna *population.DNA
	// CurrentLevel int
	// Walls        []utils.Obstacle
	// MoveLimit    int
	// Goal         utils.Vector
	// Counter      int
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	// screen.DrawImage()
}

const (
	screenWidth  = 1280
	screenHeight = 720
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	// Your existing update code...
	// g.Counter++
	// genetics.Genetic.Update()

	// for _, individual := range genetics.Genetic.Population {
	// 	if individual.IsAlive {
	// individual.CheckCollision(g.Walls)
	// individual.Update(g.Counter)

	// Draw the individual
	// individual.Draw(screen)
	// 	}
	// }

	g.world.Update()
	fmt.Println("Update")

	return nil
}

func main() {
	// settings, _ := utils.LoadGameSettings()
	utils.LoadGeneticSettings()

	// engine.Init()
	// fmt.Println(settings.GameWidth, settings.GameHeight)
	// ebiten.SetWindowSize(settings.GameWidth, settings.GameHeight)
	// // ebiten.SetWindowTitle("Go Genetic Algorithm Maze")

	// // engine.GameInstance.Genetic = population.GeneticBox{PopulationSize: settings.PopulationSize}

	// if err := ebiten.RunGame(engine.GameInstance); err != nil {
	// 	log.Fatal(err)
	// }

	g := &Game{
		world: NewWorld(screenWidth, screenHeight),
		// box:   population.NewBox(),
		dna: &population.DNA{},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Genetic Algorithm in Go")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	// for i := 0; i < settings.Iterations; i++ {
	// 	genetic := population.GeneticBox{PopulationSize: settings.PopulationSize}

	// 	genCount := 0
	// 	exitGame := false
	// 	counter := 0

	// 	for i := 0; i < 1; i++ {
	// 		// for i := 0; i < settings.MaxGenerations; i++ {
	// 		for _, individual := range genetic.Population {
	// 			if individual.IsAlive {
	// 				individual.CheckCollision(engine.GameInstance.Walls)
	// 				individual.Update()
	// 			}

	// 		}
	// 	}
	// }
}
