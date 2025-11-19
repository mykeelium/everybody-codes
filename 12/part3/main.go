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

func runTrial(startX int, startY int, maxX int, maxY int, barrels [][]Barrel) (int, [][]Barrel) {
	poppedBarrels := 0
	toPop := []point{}
	if checkPoint(startX, startY, maxX, maxY, 10, barrels) {
		toPop = append(toPop, point{X: startX, Y: startY})
	}
	for len(toPop) != 0 {
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

	return poppedBarrels, barrels
}

type Barrel struct {
	X      int
	Y      int
	Value  int
	Popped bool
}

func copyBarrels(barrels [][]Barrel) [][]Barrel {
	rec := [][]Barrel{}
	for _, barrelList := range barrels {
		newList := []Barrel{}
		for _, barrel := range barrelList {
			newBarrel := Barrel{
				X:      barrel.X,
				Y:      barrel.Y,
				Value:  barrel.Value,
				Popped: barrel.Popped,
			}
			newList = append(newList, newBarrel)
		}
		rec = append(rec, newList)
	}
	return rec
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q12_p3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines := []string{}

	barrels := [][]Barrel{}
	testBarrels := [][]Barrel{}
	workingBarrels := [][]Barrel{}

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

	totalPoppedBarrels := 0
	maxPoppedBarrels := 0
	currentMaxPoppedX := 0
	currentMaxPoppedY := 0
	maxPoppedX := []int{}
	maxPoppedY := []int{}

	fmt.Println("First barrel")
	for x := range maxX + 1 {
		testBarrels = copyBarrels(barrels)
		for y := range maxY + 1 {
			poppedBarrels, _ := runTrial(x, y, maxX, maxY, testBarrels)
			if poppedBarrels > maxPoppedBarrels {
				maxPoppedBarrels = poppedBarrels
				currentMaxPoppedX = x
				currentMaxPoppedY = y
			}
		}
	}

	fmt.Printf("First barrel found at x: %d, y: %d, popping %d barrels\n", currentMaxPoppedX, currentMaxPoppedY, maxPoppedBarrels)

	workingBarrels = copyBarrels(barrels)
	_, workingBarrels = runTrial(currentMaxPoppedX, currentMaxPoppedY, maxX, maxY, workingBarrels)
	maxPoppedX = append(maxPoppedX, currentMaxPoppedX)
	maxPoppedY = append(maxPoppedY, currentMaxPoppedY)
	totalPoppedBarrels += maxPoppedBarrels
	maxPoppedBarrels = 0
	currentMaxPoppedX = 0
	currentMaxPoppedY = 0

	for x := range maxX + 1 {
		testBarrels = copyBarrels(workingBarrels)
		for y := range maxY + 1 {
			poppedBarrels, _ := runTrial(x, y, maxX, maxY, testBarrels)
			if poppedBarrels > maxPoppedBarrels {
				maxPoppedBarrels = poppedBarrels
				currentMaxPoppedX = x
				currentMaxPoppedY = y
			}
		}
	}

	fmt.Printf("second barrel found at x: %d, y: %d, popping %d barrels\n", currentMaxPoppedX, currentMaxPoppedY, maxPoppedBarrels)

	_, workingBarrels = runTrial(currentMaxPoppedX, currentMaxPoppedY, maxX, maxY, workingBarrels)
	maxPoppedX = append(maxPoppedX, currentMaxPoppedX)
	maxPoppedY = append(maxPoppedY, currentMaxPoppedY)
	totalPoppedBarrels += maxPoppedBarrels
	maxPoppedBarrels = 0
	currentMaxPoppedX = 0
	currentMaxPoppedY = 0
	maxPoppedBarrels = 0

	for x := range maxX + 1 {
		testBarrels = copyBarrels(workingBarrels)
		for y := range maxY + 1 {
			poppedBarrels, _ := runTrial(x, y, maxX, maxY, testBarrels)
			if poppedBarrels > maxPoppedBarrels {
				maxPoppedBarrels = poppedBarrels
				currentMaxPoppedX = x
				currentMaxPoppedY = y
			}
		}
	}

	fmt.Printf("last barrel found at x: %d, y: %d, popping %d barrels\n", currentMaxPoppedX, currentMaxPoppedY, maxPoppedBarrels)

	maxPoppedX = append(maxPoppedX, currentMaxPoppedX)
	maxPoppedY = append(maxPoppedY, currentMaxPoppedY)
	totalPoppedBarrels += maxPoppedBarrels
	maxPoppedBarrels = 0
	currentMaxPoppedX = 0
	currentMaxPoppedY = 0
	maxPoppedBarrels = 0

	poppedBarrels := 0
	toPop := []point{}
	toPop = append(toPop, point{X: maxPoppedX[0], Y: maxPoppedY[0]})
	toPop = append(toPop, point{X: maxPoppedX[1], Y: maxPoppedY[1]})
	toPop = append(toPop, point{X: maxPoppedX[2], Y: maxPoppedY[2]})
	for len(toPop) != 0 {
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

	fmt.Printf("Combined Output: %d\n", totalPoppedBarrels)
	fmt.Printf("X Outputs: %v, Y Outputs: %v\n", maxPoppedX, maxPoppedY)
	fmt.Printf("Single run output: %d\n", poppedBarrels)
}
