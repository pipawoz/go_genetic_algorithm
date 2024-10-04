# Go Genetic Algorithm

An implementation of a genetic algorithm in Go using Ebiten for visualization.

## Overview

This project is a Go implementation of a genetic algorithm that simulates a population of individuals moving towards a goal. The individuals start on the left side of the screen and evolve over generations to reach the goal on the right side. The simulation runs over multiple generations, allowing the individuals to evolve and improve their performance through selection, crossover, and mutation.

## Features

- **Genetic Algorithm**: Implements selection, crossover, and mutation to evolve the population over generations.
- **Visualization with Ebiten**: Uses the Ebiten game library to visualize the simulation in real-time.
- **Configurable Levels**: Includes multiple levels with different obstacles to challenge the individuals.
- **Trail Visualization**: Optionally displays the trails of individuals to visualize their paths.

## Project Structure

The project follows Go's best practices for project layout:

```
GO_GENETIC_ALGORITHM/
├── cmd/
│   └── go_genetic_algorithm/
│       └── main.go
├── internal/
│   ├── engine/
│   │   ├── engine.go
│   │   └── levels.go
│   ├── genetics/
│   │   └── genetic_box.go
│   ├── population/
│   │   ├── box.go
│   │   └── dna.go
│   └── utils/
│       └── utils.go
│       └── config.go
├── configs/
│   ├── genetic_settings.json
│   └── settings.json
├── go.mod
├── go.sum
└── README.md
```

- `cmd/`: Contains the `main.go` file for the executable application.
- `internal/`: Contains the internal packages of the project.
    - `engine/`: Handles the game engine and levels.
    - `genetics/`: Implements the genetic algorithm logic.
    - `population/`: Defines the individual entities and their genetic representation.
    - `utils/`: Provides utility functions and settings.
- `configs/`: Stores configuration files in JSON format.

## Prerequisites

- **Go 1.20 or later**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org).
- **Ebiten v2**: The project uses the Ebiten game library for visualization.

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/go_genetic_algorithm.git
cd go_genetic_algorithm
```

Install dependencies:

```bash
go mod tidy
```

This will download and install all the necessary dependencies specified in `go.mod`.

## Usage

Navigate to the `cmd/go_genetic_algorithm/` directory and run the application:

```bash
cd cmd/go_genetic_algorithm
go run main.go
```

Alternatively, you can build the application and run the executable:

```bash
go build -o genetic_algorithm
./genetic_algorithm
```

## Configuration

The application can be configured using the JSON files located in the `configs/` directory.

- `settings.json`: Contains general settings for the application.
- `genetic_settings.json`: Contains settings specific to the genetic algorithm, such as mutation rate, crossover rate, population size, etc.

### Example `settings.json`

```json
{
    "printTrace": true,
    "currentLevel": 5,
    "simulateOnly": false
}
```

### Example `genetic_settings.json`

```json
{
    "iterations": 5,
    "maxGenerations": 200,
    "populationSize": 100,
    "mutationRate": 0.05,
    "crossoverRate": 1
}
```

You can adjust these settings to change the behavior of the simulation.

## Levels

The application includes multiple levels with different obstacles. You can select the level by modifying the `level` variable in the `main.go` or through the settings.

### Available Levels

1. **Level 1**: Basic level with no obstacles.
2. **Level 2**: Introduces a vertical obstacle in the middle.
3. **Level 3**: Adds multiple obstacles creating a more complex path.
4. **Level 4**: Increases complexity with additional obstacles.
5. **Level 5**: The most challenging level with multiple obstacles.

## Development

### Project Structure Details

- `cmd/go_genetic_algorithm/main.go`: The entry point of the application. Sets up the game window, initializes the game, and starts the main loop.
- `internal/engine/`: Contains the game loop logic and level definitions.
    - `engine.go`: Defines the `Game` struct and the main game loop methods (`Update`, `Draw`, `Layout`).
    - `levels.go`: Contains the `SelectLevel` function that defines the obstacles and move limits for each level.
- `internal/genetics/`: Implements the genetic algorithm.
    - `genetic_box.go`: Defines the `GeneticBox` struct, which manages the population and the genetic operations (`Init`, `NextGeneration`, etc.).
- `internal/population/`: Contains the definitions of individuals and their genetic makeup.
    - `box.go`: Defines the `Box` struct representing an individual, with methods for updating state, drawing, resetting, and genetic operations (`Mutate`, `Crossover`, etc.).
    - `dna.go`: Defines the `DNA` struct, representing the genetic sequence of an individual, and methods for initialization and mutation.
- `internal/utils/`: Provides utility functions and settings management.
    - `utils.go`: Contains common structs and functions used across the application, such as `Vector`, `Obstacle`, and settings loading functions.

## Building and Running Tests

Currently, the project does not include automated tests. Contributions are welcome to add testing to improve code quality and reliability.

## Contributing

Contributions are welcome! If you'd like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes with clear and descriptive messages.
4. Push your branch to your fork.
5. Open a pull request describing your changes.

Please ensure your code adheres to the project's coding standards and includes appropriate documentation.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contact

For any questions or suggestions, feel free to open an issue or contact the maintainer at [pedrowozniak@hotmail.com].

