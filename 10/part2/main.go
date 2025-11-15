package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func checkPoint(x int, y int, x2 int, y2 int) bool {
	return (x2 <= x && x2 >= 0) && (y2 <= y && y2 >= 0)
}

func checkSheep(lines []string, x int, y int) ([]string, bool) {
	rec := false
	if lines[y][x] == 'S' {
		line := []rune(lines[y])
		line[x] = 'X'
		lines[y] = string(line)
		rec = true
	}
	return lines, rec
}

func handlePoint(x int, y int, lines []string, x2 int, y2 int, sheepTotal int, hideOuts []point, newPoints []point) (int, []point, []string) {
	if checkPoint(x, y, x2, y2) {
		if !containsPoint(x2, y2, hideOuts) {
			hasSheep := false
			if lines, hasSheep = checkSheep(lines, x2, y2); hasSheep {
				sheepTotal++
			}
		}
		newPoint := point{X: x2, Y: y2}
		newPoints = append(newPoints, newPoint)
	}

	return sheepTotal, newPoints, lines
}

func containsPoint(x int, y int, points []point) bool {
	for _, point := range points {
		if point.X == x && point.Y == y {
			return true
		}
	}
	return false
}

func moveSheep(maxX int, maxY int, hideOuts []point, dragonPoints []point, lines []string, sheepTotal int) (int, []string) {
	for y := range maxY + 1 {
		row := maxY - y
		if row == maxY {
			lines[row] = strings.ReplaceAll(lines[row], "S", ".")
		} else {
			line := []rune(lines[row])
			for x := range maxX + 1 {
				if line[x] == 'S' {
					if !containsPoint(x, row+1, hideOuts) && containsPoint(x, row+1, dragonPoints) {
						sheepTotal++
						line[x] = 'X'
					} else {
						nextLine := []rune(lines[row+1])
						nextLine[x] = 'S'
						lines[row+1] = string(nextLine)
					}
					line[x] = '.'
				}
			}
			lines[row] = string(line)
		}
	}

	return sheepTotal, lines
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q10_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	x := len(lines[0]) - 1
	y := len(lines) - 1

	sheepEaten := 0
	pointsToCheck := []point{{X: x / 2, Y: y / 2}}
	hidingSpots := []point{}

	for i := range y + 1 {
		for j := range x + 1 {
			if lines[i][j] == '#' {
				hidingSpots = append(hidingSpots, point{X: j, Y: i})
			}
		}
	}

	for i := range 20 {
		fmt.Printf("Iteration: %d\n", i)
		newPoints := []point{}
		for _, point := range pointsToCheck {
			x_2 := point.X - 2
			x_1 := point.X - 1
			x1 := point.X + 1
			x2 := point.X + 2
			y_2 := point.Y - 2
			y_1 := point.Y - 1
			y1 := point.Y + 1
			y2 := point.Y + 2

			if !containsPoint(x_2, y_1, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_2, y_1, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x_2, y1, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_2, y1, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x_1, y_2, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_1, y_2, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x_1, y2, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_1, y2, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x1, y_2, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x1, y_2, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x1, y2, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x1, y2, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x2, y_1, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x2, y_1, sheepEaten, hidingSpots, newPoints)
			}
			if !containsPoint(x2, y1, newPoints) {
				sheepEaten, newPoints, lines = handlePoint(x, y, lines, x2, y1, sheepEaten, hidingSpots, newPoints)
			}
		}
		sheepEaten, lines = moveSheep(x, y, hidingSpots, newPoints, lines, sheepEaten)

		pointsToCheck = newPoints

	}

	fmt.Printf("Output: %d\n", sheepEaten)
}
