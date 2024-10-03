package genetics

import (
	"math/rand"
	"reflect"
	"sort"

	"github.com/pipawoz/go_genetic_algorithm/internal/population"
	"github.com/pipawoz/go_genetic_algorithm/internal/utils"
)

// Genetic GeneticBox represents the genetic algorithm box.
var Genetic GeneticBox

// GeneticBox represents the genetic algorithm box.
type GeneticBox struct {
	Crashed        bool
	Won            bool
	WonTime        int
	Gene           population.DNA
	Acceleration   utils.Vector
	Velocity       utils.Vector
	Fitness        float64
	AvgFitness     float64
	AvgDistance    int
	PopulationSize int
	Population     []population.Box
}

func NewGeneticBox(populationSize int) *GeneticBox {
	g := &GeneticBox{}
	dna := &population.DNA{}
	g.Init(dna, populationSize)
	return g
}

// Init initializes the genetic box with the given DNA and population size.
func (g *GeneticBox) Init(dna *population.DNA, populationSize int) {
	g.Crashed = false
	g.Won = false
	g.WonTime = 0

	// if dna.Chain == nil {
	// 	dna.NewDNA(nil)
	// }

	g.PopulationSize = populationSize
	g.Acceleration.X = 0
	g.Acceleration.Y = 0
	g.Velocity.X = 0
	g.Velocity.Y = 0
	g.Fitness = 0

	for i := 0; i < g.PopulationSize; i++ {

		individualDNA := population.DNA{}
		individualDNA.NewDNA(nil) // Inicializa con genes aleatorios

		// if dna.Chain == nil {
		// }

		// if populationSize == 1 {
		// box := population.NewBox(dna)
		// g.Population = append(g.Population, *box)
		// } else {
		box := population.NewBox(&individualDNA)
		g.Population = append(g.Population, *box)
	}
}

// GetBestBox Returns the best box in the population.
func (g *GeneticBox) GetBestBox() population.Box {
	for i := range g.Population {
		g.Population[i].CalculateFitness()
	}

	sort.Slice(g.Population, func(i, j int) bool {
		return g.Population[i].Fitness < g.Population[j].Fitness
	})

	return g.Population[len(g.Population)-1]
}

// GetAvgFitness Returns the average fitness of the population.
func (g *GeneticBox) GetAvgFitness() float64 {
	var sumFit float64 = 0

	// for i := range g.Population {
	// 	g.Population[i].CalculateFitness()
	// }

	for i := range g.Population {
		sumFit += g.Population[i].Fitness
	}

	return sumFit / float64(len(g.Population))
}

// GetAvgDistance Gets the average distance of the population.
func (g *GeneticBox) GetAvgDistance() int {
	var sumDist int = 0

	// for i := range g.Population {
	// 	g.Population[i].CalculateFitness()
	// }

	for i := range g.Population {
		sumDist += int(g.Population[i].Dist)
	}

	return sumDist / len(g.Population)
}

// GetAvgTraveled Gets the average distance traveled by the population.
func (g *GeneticBox) GetAvgTraveled() float64 {
	var avg float64 = 0

	for i := range g.Population {
		g.Population[i].CalculateFitness()
	}

	for i := range g.Population {
		avg += g.Population[i].Traveled
	}

	return avg / float64(len(g.Population))
}

// Update Updates the genetic box.
// func (g *GeneticBox) Update() {
// 	fmt.Println("Update")
// }

// NextGeneration Updates the population to the next generation.
func (g *GeneticBox) NextGeneration() {

	// Selection: Calculate the probability of each individual to pass.
	// More fit individuals have a higher probability of continuing to the next generation.
	// - Calculate the fitness for all individuals
	// - Calculate the total fitness of all individuals and then calculate the probability
	// of each individual to pass
	// - Randomly choose individuals from the list, creating a new genetic pool, and add them to a list

	for i := range g.Population {
		g.Population[i].CalculateFitness()
	}

	g.AvgFitness = g.GetAvgFitness()
	g.AvgDistance = g.GetAvgDistance()

	var newSelection = []population.Box{}

	totalFitness := 0.0
	for i := range g.Population {
		totalFitness += g.Population[i].Fitness
	}

	probabilityOfSelection := []float64{}
	for i := range g.Population {
		probabilityOfSelection = append(probabilityOfSelection, g.Population[i].Fitness/totalFitness)
	}

	for range g.Population {
		selection := rand.Float64()
		cumulativeProbability := 0.0
		for i, individualProbability := range probabilityOfSelection {
			cumulativeProbability += individualProbability
			if selection <= cumulativeProbability {
				newSelection = append(newSelection, g.Population[i])
				break
			}
		}
	}

	// Crossover
	// - Randomly select parent A and parent B from the list after selection
	// - If they are the same individual, choose another
	// - Perform the crossover and obtain 2 offspring
	// - Save the offspring in a new list, do not save the parents (they die)

	crossoverList := []population.Box{}
	for i := 1; i < g.PopulationSize; i += 2 {
		if len(newSelection) < 2 {
			break
		}

		parentA := newSelection[rand.Intn(len(newSelection))]
		parentB := newSelection[rand.Intn(len(newSelection))]

		if reflect.DeepEqual(parentA, parentB) {
			i--
			continue
		}

		offspringA, offspringB := parentA.Crossover(parentB)
		crossoverList = append(crossoverList, offspringA)
		crossoverList = append(crossoverList, offspringB)
	}

	// Mutation
	// - Perform the mutation for each individual
	// - Each individual will mutate a certain configurable percentage with a random probability

	for i := range crossoverList {
		crossoverList[i].Mutate(i)
	}

	// Replace the population with the new generation
	g.Population = crossoverList

	// Reset all the individuals in the population
	for i := range g.Population {
		g.Population[i].Reset()
	}

}
