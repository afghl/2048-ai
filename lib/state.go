package lib

import "github.com/afghl/2048-ai/utils"

type GameState struct {
	Size int
	Grid [][]int

	Weight float32 // a weight of a generated random tile state
}

func NewState(size int, grid [][]int, weight float32) GameState {
	return GameState{Size: size, Grid: grid, Weight: weight}
}

// GetLegalActions return a list of legal actions
func (s *GameState) GetLegalActions() []Direction {
	return actions
}

// SuccessorState get a successor state of action
func (s *GameState) SuccessorState(direction Direction) GameState {
	grid := move(s.Grid, direction)
	return NewState(s.Size, grid, 1)
}

// GenerateRandomTileState for current state s, generate its successor state by create random tile
func (s *GameState) GenerateRandomTileState() []GameState {
	arr := make([]GameState, 0)
	for i := 0; i < len(s.Grid); i++ {
		for j := 0; j < len(s.Grid[i]); j++ {
			if s.Grid[i][j] == 0 {
				// can add new tile in this grid
				arr = append(arr, generateTileState(*s, i, j, 2, 0.9))
				arr = append(arr, generateTileState(*s, i, j, 4, 0.1))
			}
		}
	}
	return arr
}

func generateTileState(state GameState, x, y, tile int, weight float32) GameState {
	grid := utils.DeepCopy(state.Grid)
	grid[x][y] = tile
	return NewState(len(grid), grid, weight)
}

// move to move grid to
func move(grid [][]int, direction Direction) [][]int {
	data := utils.DeepCopy(grid)
	size := len(grid)
	switch direction {
	case UP:
		for y := 0; y < size; y++ {
			for x := 0; x < size-1; x++ {
				for nx := x + 1; nx < size; nx++ {
					if data[nx][y] == 0 {
						continue
					}
					if data[x][y] <= 0 {
						data[x][y] = data[nx][y]
						data[nx][y] = 0
						x -= 1
					} else if data[x][y] == data[nx][y] {
						data[x][y] += data[nx][y]
						data[nx][y] = 0
					}
					break
				}
			}
		}
	case DOWN:
		for y := 0; y < size; y++ {
			for x := size - 1; x > 0; x-- {
				for nx := x - 1; nx >= 0; nx-- {
					if data[nx][y] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[nx][y]
							data[nx][y] = 0
							x += 1
						} else if data[x][y] == data[nx][y] {
							data[x][y] += data[nx][y]
							data[nx][y] = 0
						}
						break
					}
				}
			}
		}
	case LEFT:
		for x := 0; x < size; x++ {
			for y := 0; y < size-1; y++ {
				for ny := y + 1; ny < size; ny++ {
					if data[x][ny] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[x][ny]
							data[x][ny] = 0
							y -= 1
						} else if data[x][y] == data[x][ny] {
							data[x][y] += data[x][ny]
							data[x][ny] = 0
						}
						break
					}
				}
			}
		}
	case RIGHT:
		for x := 0; x < size; x++ {
			for y := size - 1; y > 0; y-- {
				for ny := y - 1; ny >= 0; ny-- {
					if data[x][ny] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[x][ny]
							data[x][ny] = 0
							y += 1
						} else if data[x][y] == data[x][ny] {
							data[x][y] += data[x][ny]
							data[x][ny] = 0
						}
						break
					}
				}
			}
		}
	}
	return data
}
