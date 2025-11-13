package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func all[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

type line struct {
	start int32
	end   int32
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q08_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	connections := make(map[int32][]line)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	pointsRaw := strings.Split(inputRaw, ",")
	var points []int32
	for _, pointRaw := range pointsRaw {
		point, _ := strconv.ParseInt(pointRaw, 10, 32)
		points = append(points, int32(point))
	}

	crosses := 0

	for i, point := range points {
		if i == 0 {
			continue
		}
		newLine := line{start: points[i-1], end: point}
		if points[i-1] < point {
			for x := points[i-1] + 1; x < point; x++ {
				listToCheck := connections[x]
				for _, item := range listToCheck {
					if (item.start > points[i-1] && item.start < point && (item.end > point || item.end < points[i-1])) ||
						(item.end > points[i-1] && item.end < point && (item.start > point || item.start < points[i-1])) {
						crosses++
					}
				}
			}
		} else {
			for x := point + 1; x < points[i-1]; x++ {
				listToCheck := connections[x]
				for _, item := range listToCheck {
					if (item.start > point && item.start < points[i-1] && (item.end > points[i-1] || item.end < point)) ||
						(item.end > point && item.end < points[i-1] && (item.start > points[i-1] || item.start < point)) {
						crosses++
					}
				}
			}
		}

		startLineList := connections[points[i-1]]
		startLineList = append(startLineList, newLine)
		connections[points[i-1]] = startLineList
		endLineList := connections[point]
		endLineList = append(endLineList, newLine)
		connections[point] = endLineList
	}

	fmt.Printf("Output: %v\n", crosses)
}
