package population

import (
	"math"
	"math/rand"

	"github.com/pipawoz/go_genetic_algorithm/internal/utils"
)

// DNA of the individual. It contains the genes of the individual.
type DNA struct {
	Chain []utils.Vector
}

// NewDNA creates a new DNA object with the given genes
// If genes is nil, it will create a new DNA object with random genes
// If genes is not nil, it will create a new DNA object with the given genes
// The genes are a sequence of random numbers that represent the path
func (dna *DNA) NewDNA(genes []utils.Vector) *DNA {
	if genes != nil {
		dna.Chain = genes
	} else {
		for i := 0; i < 1000; i++ {
			angle := float64(rand.Intn(360)) * math.Pi / 180
			dna.Chain = append(dna.Chain, utils.Vector{X: float32(math.Cos(angle)),
				Y: float32(math.Sin(angle))})
		}
	}

	return dna
}
