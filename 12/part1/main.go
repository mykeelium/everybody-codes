package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type point struct {
	X int
	Y int
}

func handlePoint(x int, y int, xMax int, yMax, barrelValue int, barrels [][]Barrel, newPoints []point) []point {
	if checkPoint(x, y, xMax, yMax, barrelValue, barrels) {
		newPoints = append(newPoints, point{X: x, Y: y})
	}
	return newPoints
}

func checkPoint(x int, y int, xMax int, yMax int, barrelValue int, barrels [][]Barrel) bool {
	return x >= 0 && y >= 0 && x <= xMax && y <= yMax && !barrels[x][y].Popped && barrelValue >= barrels[x][y].Value
}

func containsPoint(x int, y int, points []point) bool {
	for _, point := range points {
		if point.X == x && point.Y == y {
			return true
		}
	}
	return false
}

type Barrel struct {
	X      int
	Y      int
	Value  int
	Popped bool
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q12_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines := []string{}

	barrels := [][]Barrel{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	maxX := 0
	maxY := 0

	for x, line := range lines {
		newLine := []Barrel{}
		for y, barrel := range line {
			val, _ := strconv.ParseInt(string(barrel), 10, 32)
			newLine = append(newLine, Barrel{
				X:      x,
				Y:      y,
				Value:  int(val),
				Popped: false,
			})
		}
		if len(newLine) > maxY {
			maxY = len(newLine) - 1
		}
		barrels = append(barrels, newLine)
	}
	maxX = len(barrels) - 1

	toPop := []point{{X: 0, Y: 0}}
	poppedBarrels := 0

	iteration := 0
	for {
		fmt.Printf("Iteration: %d, toPopLength: %d\n", iteration, len(toPop))
		iteration++
		if len(toPop) == 0 {
			break
		}

		newToPop := []point{}

		for _, loc := range toPop {
			barrel := barrels[loc.X][loc.Y]
			barrel.Popped = true
			barrels[loc.X][loc.Y] = barrel
			poppedBarrels++

			if !containsPoint(loc.X-1, loc.Y, newToPop) {
				newToPop = handlePoint(loc.X-1, loc.Y, maxX, maxY, barrel.Value, barrels, newToPop)
			}
			if !containsPoint(loc.X, loc.Y-1, newToPop) {
				newToPop = handlePoint(loc.X, loc.Y-1, maxX, maxY, barrel.Value, barrels, newToPop)
			}
			if !containsPoint(loc.X+1, loc.Y, newToPop) {
				newToPop = handlePoint(loc.X+1, loc.Y, maxX, maxY, barrel.Value, barrels, newToPop)
			}
			if !containsPoint(loc.X, loc.Y+1, newToPop) {
				newToPop = handlePoint(loc.X, loc.Y+1, maxX, maxY, barrel.Value, barrels, newToPop)
			}
		}

		toPop = newToPop
	}

	fmt.Printf("Output: %d\n", poppedBarrels)
}
