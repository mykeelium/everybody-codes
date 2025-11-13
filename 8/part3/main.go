package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q08_p3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	connections := make(map[int][]line)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	pointsRaw := strings.Split(inputRaw, ",")
	var points []int
	for _, pointRaw := range pointsRaw {
		point, _ := strconv.ParseInt(pointRaw, 10, 32)
		points = append(points, int(point))
	}

	for i, point := range points {
		if i == 0 {
			continue
		}
		newLine := line{start: points[i-1], end: point}
		startLineList := connections[points[i-1]]
		startLineList = append(startLineList, newLine)
		connections[points[i-1]] = startLineList
		endLineList := connections[point]
		endLineList = append(endLineList, newLine)
		connections[point] = endLineList
	}

	largestCuts := 0
	largestI := 0
	largestJ := 0

	for i := range 257 {
		for j := range 257 {
			if i == j || i == 0 || j == 0 {
				continue
			}
			cuts := 0
			checkFullThreadCutsList := connections[i]
			fullThreadsCut := 0
			for _, item := range checkFullThreadCutsList {
				if (item.start == i && item.end == j) || (item.end == i && item.start == j) {
					fullThreadsCut++
				}
			}
			cuts += (fullThreadsCut / 2)
			if i < j {
				for x := i + 1; x < j; x++ {
					listToCheck := connections[x]
					for _, item := range listToCheck {
						if (item.start > i && item.start < j && (item.end > j || item.end < i)) ||
							(item.end > i && item.end < j && (item.start > j || item.start < i)) {
							cuts++
						}
					}
				}
			} else {
				for x := j + 1; x < i; x++ {
					listToCheck := connections[x]
					for _, item := range listToCheck {
						if (item.start > j && item.start < i && (item.end > i || item.end < j)) ||
							(item.end > j && item.end < i && (item.start > i || item.start < j)) {
							cuts++
						}
					}
				}
			}
			if cuts > largestCuts {
				largestCuts = cuts
				largestI = i
				largestJ = j
			}
		}
	}

	fmt.Printf("Output: %v\n", largestCuts)
	fmt.Printf("Output I: %v\n", largestI)
	fmt.Printf("Output J: %v\n", largestJ)
}
