package population

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/pipawoz/go_genetic_algorithm/utils"
)

type Box struct {
	IsAlive      bool
	AliveTime    int
	Position     utils.Vector
	Size         int
	Fitness      float64
	Won          bool
	Acceleration utils.Vector
	Velocity     utils.Vector
	Traveled     int
	Genes        DNA
	Dist         float64
	Frames       int
}

func NewBox() *Box {
	return &Box{
		IsAlive:      true,
		AliveTime:    0,
		Position:     utils.Vector{X: 10, Y: float32(utils.Settings.GameHeight) / 2},
		Size:         5,
		Fitness:      0,
		Won:          false,
		Acceleration: utils.Vector{X: 0, Y: 0},
		Velocity:     utils.Vector{X: 0, Y: 0},
		Traveled:     0,
		Genes:        DNA{},
		Dist:         0,
		Frames:       0,
	}
}

func (box *Box) SetGenes(genes []utils.Vector) {
	box.Genes.NewDNA(genes)
}

// CheckCollision checks if the box collides with any walls or goes out of the game boundaries.
// If a collision is detected, the box's IsAlive flag is set to false.
// Parameters:
// - walls: a slice of engine.Obstacle representing the walls in the game.
// Returns: none.
func (box *Box) CheckCollision(walls []utils.Obstacle) {
	// Check if the box is out of the game boundaries
	if int(box.Position.X)+box.Size > utils.Settings.GameWidth ||
		box.Position.X < 0 || box.Position.Y < 0 || int(box.Position.Y)+
		box.Size > utils.Settings.GameHeight {
		box.IsAlive = false
	}

	// Check if the box collides with any wall
	for _, wall := range walls {
		if int(box.Position.X) < wall.X+wall.Width && int(box.Position.X)+
			box.Size > wall.X && int(box.Position.Y) < wall.Y+wall.Height && int(box.Position.Y)+box.Size > wall.Y {
			box.IsAlive = false
		}
	}
}

// CalculateFitness calculates the fitness of the box based on its position, distance to the goal, and other factors.
// It updates the Fitness field of the box.
func (box *Box) CalculateFitness() {
	// Calculate the fitness of the box
	box.Dist = math.Sqrt(math.Pow(float64(box.Position.X-float32(utils.Settings.GoalX)), 2) +
		math.Pow(float64(box.Position.X-float32(utils.Settings.GoalY)), 2))

	box.Fitness = 1 - (box.Dist / float64(utils.Settings.GameWidth))

	if box.IsAlive {
		box.Fitness *= 1.5
	}

	if box.Won {

		// TODO: Cambiar por movelimit
		efficiency := 700 / box.Frames
		box.Fitness *= (2 * float64(efficiency))
	}
}

// Reset resets the state of the Box.
func (box *Box) Reset() {
	box.IsAlive = false
	box.Position.X = 10
	box.Position.Y = float32(utils.Settings.GameHeight) / 2
	box.Velocity.X = 0
	box.Velocity.Y = 0
	box.Acceleration.X = 0
	box.Acceleration.Y = 0
	box.Won = false
	box.Traveled = 0
	box.Frames = 0
}

// Update updates the state of the Box.
// If the Box has reached the goal, it sets the Frames and Won properties,
// stops the Box's movement, and marks it as not alive.
// It also replaces the remaining genes with no-ops.
// The Box's acceleration is determined by the current frame of the game.
// The Box's velocity is updated based on its acceleration.
// The Box's position is updated based on its velocity.
// The Traveled property is updated based on the distance traveled by the Box.
func (box *Box) Update(counter int) {
	if !box.IsAlive {
		box.Frames = counter
	}

	// Check if the box has reached the goal
	boxRect := image.Rect(int(box.Position.X), int(box.Position.Y),
		int(box.Position.X)+box.Size, int(box.Position.Y)+box.Size)

	winRect := image.Rect(80, 80, utils.Settings.GoalX+40, utils.Settings.GoalY-40)

	if boxRect.Overlaps(winRect) && !box.Won {
		box.Frames = counter
		box.Won = true
		box.Velocity = utils.Vector{X: 0, Y: 0}
		box.Acceleration = utils.Vector{X: 0, Y: 0}
		box.IsAlive = false
		// box.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}

		// Replace the remaining genes with no-ops
		for i := range box.Genes.Chain[box.Frames+1:] {
			box.Genes.Chain[box.Frames+1+i] = utils.Vector{X: 0, Y: 0} // Replace with nop's
		}
	}

	if counter > 0 {
		box.Acceleration = box.Genes.Chain[counter-1]
	} else {
		box.Acceleration = box.Genes.Chain[counter]
	}

	box.Velocity.X += box.Acceleration.X
	box.Velocity.Y += box.Acceleration.Y

	// # Hago unos ajustes por temas visuales
	// self.velLimit = 6

	// if self.vel.x > self.velLimit and self.acc.x > 0:
	//     self.vel.x = self.velLimit
	// if self.vel.x < -self.velLimit and self.acc.x < 0:
	//     self.vel.x = -self.velLimit
	// if self.vel.y > self.velLimit and self.acc.y > 0:
	//     self.vel.y = self.velLimit
	// if self.vel.y < -self.velLimit and self.acc.y < 0:
	//     self.vel.y = -self.velLimit

	box.Position.X += box.Velocity.X
	box.Position.Y += box.Velocity.Y

	box.Traveled += int(math.Sqrt(math.Pow(float64(box.Velocity.X), 2) +
		math.Pow(float64(box.Velocity.Y), 2)))
}

// Mutate applies mutation to the Box's genes based on the mutation rate specified in the settings.
// If the mutation rate is met, a random gene in the Box's gene chain is selected and its X or Y value is multiplied by 1.01.
// The mutation quantity is set to 1 by default.
func (box *Box) Mutate() {
	if rand.Float64() < utils.DNASettings.MutationRate {
		mutationQuantity := 1
		for i := 0; i < mutationQuantity; i++ {
			index := rand.Int() % (len(box.Genes.Chain) - 1)
			if rand.Int()%10 > 5 {
				box.Genes.Chain[index].Y *= 1.01
			} else {
				box.Genes.Chain[index].X *= 1.01
			}
		}
	}
}

func (box *Box) Crossover(partner Box) (Box, Box) {
	if rand.Float64() < utils.DNASettings.CrossoverRate {
		newGenes1 := make([]utils.Vector, len(box.Genes.Chain))
		newGenes2 := make([]utils.Vector, len(box.Genes.Chain))

		middlePoint := rand.Intn(len(box.Genes.Chain)-1) + 1

		for i := range box.Genes.Chain {
			if i < middlePoint {
				newGenes1[i] = partner.Genes.Chain[i]
				newGenes2[i] = box.Genes.Chain[i]
			} else {
				newGenes1[i] = box.Genes.Chain[i]
				newGenes2[i] = partner.Genes.Chain[i]
			}
		}

		return Box{Genes: DNA{Chain: newGenes1}}, Box{Genes: DNA{Chain: newGenes2}}
	}

	return Box{Genes: DNA{Chain: box.Genes.Chain}}, Box{Genes: DNA{Chain: partner.Genes.Chain}}
}

func (box *Box) Draw(screen *ebiten.Image) {
	// Create an empty image of the size of the individual
	img, _ := ebiten.NewImage(5, 5, ebiten.FilterDefault)

	// Fill the image with a color
	img.Fill(color.RGBA{255, 0, 0, 255})

	// Draw the image on the screen at the individual's position
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(box.Position.X), float64(box.Position.Y))
	screen.DrawImage(img, opts)
}

// func (individual *Individual) Draw(screen *ebiten.Image) {
// 	// Create an ebiten.Image for the individual
// 	individualImage, _ := ebiten.NewImage(individual.Width, individual.Height, ebiten.FilterDefault)

// 	// Fill the image with a color
// 	individualImage.Fill(individual.Color)

// 	// Create an ebiten.DrawImageOptions
// 	opts := &ebiten.DrawImageOptions{}

// 	// Set the position where the image should be drawn
// 	opts.GeoM.Translate(individual.X, individual.Y)

// 	// Draw the individual image onto the screen
// 	screen.DrawImage(individualImage, opts)
// }
