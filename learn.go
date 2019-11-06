package main

import (
	"fmt"
)

func main() {
	// Create a tic-tac-toe board.
	var grid []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var turn = -1
	for !isTicTacToe(grid) {
		turn *= -1
		if turn == 1 {
			printGrid(grid)
			var pos = scanMarkPosition()
			placeMark(grid, pos, turn)
		} else {
			var bestMove int = -1
			var bestMoveScore int
			for i := 0; i < 9; i++ {
				if grid[i] != 0 {
					grid[i] = -1
					score := minimax(grid, 6, false)
					grid[i] = 0
					if bestMove == -1 || bestMoveScore > score {
						bestMove = i
						bestMoveScore = score
					}
				}
			}
			grid[bestMove] = -1
		}
	}
	fmt.Printf("TicTacGo!! winner is: %v\n", toMarkString(turn))
}

var rays [][]int = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{6, 4, 2},
}

func isTicTacToe(grid []int) bool {
	for i := 0; i < len(rays); i++ {
		var x, y, z = rays[i][0], rays[i][1], rays[i][2]
		if (grid[x] == grid[y] && grid[y] == grid[z]) && (grid[x] != 0) {
			return true
		}
	}
	return false
}

func isDraw(grid []int) bool {
	if isTicTacToe(grid) {
		return false
	}
	for i := 0; i < 9; i++ {
		if grid[i] == 0 {
			return false
		}
	}
	return true
}

func placeMark(grid []int, position int, mark int) {
	grid[position] = mark
}

func scanMarkPosition() int {
	var pos int
	fmt.Print("Enter text: ")
	fmt.Scan(&pos)
	return pos - 1
}

func toMarkString(i int) string {
	switch {
	case i == 1:
		return "O"
	case i == -1:
		return "X"
	}
	return " "
}

func printGrid(grid []int) {
	var st string = "%v|%v|%v\n"
	var dv string = "-----"
	fmt.Printf(st, toMarkString(grid[6]), toMarkString(grid[7]), toMarkString(grid[8]))
	fmt.Println(dv)
	fmt.Printf(st, toMarkString(grid[3]), toMarkString(grid[4]), toMarkString(grid[5]))
	fmt.Println(dv)
	fmt.Printf(st, toMarkString(grid[0]), toMarkString(grid[1]), toMarkString(grid[2]))
}

func evaluate(grid []int, depth int, isMaximizing bool) int {
	tttValue := 50
	score := 0
	if isTicTacToe(grid) {
		if isMaximizing {
			score += tttValue
		} else {
			score -= tttValue
		}
	} else if isDraw(grid) {
		score = 0
	}
	if isMaximizing {
		score -= depth
	} else {
		score += depth
	}
	return score
}

func minimize(grid []int, depth int) int {
	if depth == 0 || isTicTacToe(grid) || isDraw(grid) {
		return evaluate(grid, depth, false)
	}

	bestScore := 1000000
	for i := 0; i < 9; i++ {
		if grid[i] == 0 {
			grid[i] = -1
			score := maximize(grid, depth-1)
			grid[i] = 0
			if score < bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func maximize(grid []int, depth int) int {
	if depth == 0 || isTicTacToe(grid) || isDraw(grid) {
		return evaluate(grid, depth, true)
	}

	bestScore := -1000000
	for i := 0; i < 9; i++ {
		if grid[i] == 0 {
			grid[i] = 1
			score := minimize(grid, depth-1)
			grid[i] = 0
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func minimax(grid []int, maxdepth int, isMaximizing bool) int {
	if isMaximizing {
		return maximize(grid, maxdepth)
	}
	return minimize(grid, maxdepth)
}
