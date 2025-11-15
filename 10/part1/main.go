package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x int
	y int
}

func checkPoint(x int, y int, x2 int, y2 int) bool {
	return (x2 <= x && x2 >= 0) && (y2 <= y && y2 >= 0)
}

func checkSheep(lines []string, x int, y int) ([]string, bool) {
	rec := false
	if lines[x][y] == 'S' {
		line := []rune(lines[x])
		line[y] = 'X'
		lines[x] = string(line)
		rec = true
	}
	return lines, rec
}

func handlePoint(x int, y int, lines []string, x2 int, y2 int, sheepTotal int, newPoints []point) (int, []point, []string) {
	if checkPoint(x, y, x2, y2) {
		hasSheep := false
		if lines, hasSheep = checkSheep(lines, x2, y2); hasSheep {
			sheepTotal++
		}
		newPoint := point{x: x2, y: y2}
		newPoints = append(newPoints, newPoint)
	}

	return sheepTotal, newPoints, lines
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q10_p1.txt")
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
	pointsToCheck := []point{{x: x / 2, y: y / 2}}

	for range 4 {
		newPoints := []point{}
		for _, point := range pointsToCheck {
			x_2 := point.x - 2
			x_1 := point.x - 1
			x1 := point.x + 1
			x2 := point.x + 2
			y_2 := point.y - 2
			y_1 := point.y - 1
			y1 := point.y + 1
			y2 := point.y + 2

			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_2, y_1, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_2, y1, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_1, y_2, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x_1, y2, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x1, y_2, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x1, y2, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x2, y_1, sheepEaten, newPoints)
			sheepEaten, newPoints, lines = handlePoint(x, y, lines, x2, y1, sheepEaten, newPoints)

		}

		pointsToCheck = newPoints
	}

	fmt.Printf("Output: %d\n", sheepEaten)
}
