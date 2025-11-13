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

func main() {
	file, err := os.Open("./everybody_codes_e2025_q08_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
		diff := points[i-1] - point

		if diff == 16 || diff == -16 {
			crosses++
		}
	}

	fmt.Printf("Output: %v\n", crosses)
}
