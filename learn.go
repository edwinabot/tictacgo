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
		printGrid(grid)
		var pos = scanMarkPosition()
		placeMark(grid, pos, turn)
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
