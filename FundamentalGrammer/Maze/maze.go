package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var direcs = []point {
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(q point) point {
	return point{p.i + q.i, p.j + q.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[0]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]

		if (cur == end) {
			break
		}

		Q = Q[1:]

		for _, dir := range direcs {
			next := cur.add(dir)
			val, ok := next.at(maze) 
			if !ok || val == 1 {
				// Out of bound, continue
				continue
			}

			val, ok = next.at(steps) 
			if !ok || val != 0 {
				// We have already visited this point, continue
				continue
			}

			if next == start {
				// If back to original point, continue
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("./maze.in")

	fmt.Println("Go Programming Language Implemeting BFS to traverse a maze")
	fmt.Println("1. Read matrix from file")
	for _, row := range maze {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("2. Output BFS Matrix")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, value := range row {
			fmt.Printf("%3d", value)
		}
		fmt.Println()
	}
	fmt.Println()
}