package engine

import "github.com/pipawoz/go_genetic_algorithm/utils"

// SelectLevel selects the level based on the currentLevel parameter and returns the move limit and walls for that level.
// The move limit determines the maximum number of moves allowed in the level.
// The walls represent the utils.Obstacles in the level that the player needs to navigate through.
func (game *Game) SelectLevel(currentLevel int) (int, []utils.Obstacle) {
	var moveLimit int
	var walls []utils.Obstacle

	switch currentLevel {
	case 1:
		//Load Level 1
		moveLimit = 350
	case 2:
		//Load Level 2
		moveLimit = 400
		walls = []utils.Obstacle{
			{X: 500, Y: 150, Width: 20, Height: 420},
		}
	case 3:
		//Load Level 3
		moveLimit = 500
		walls = []utils.Obstacle{
			{X: 350, Y: 200, Width: 20, Height: 320},
			{X: 750, Y: 200, Width: 20, Height: 320},
			{X: 550, Y: 0, Width: 20, Height: 200},
			{X: 550, Y: 520, Width: 20, Height: 200},
		}
	case 4:
		//Load Level 4
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
		//Load Level 5
		moveLimit = 700
		walls = []utils.Obstacle{
			{X: 200, Y: 300, Width: 20, Height: 420},
			{X: 500, Y: 0, Width: 20, Height: 350},
			{X: 800, Y: 300, Width: 20, Height: 420},
			{X: 1100, Y: 0, Width: 20, Height: 350},
		}
	}

	game.CurrentLevel = currentLevel
	game.Walls = walls
	game.MoveLimit = moveLimit

	return moveLimit, walls
}
