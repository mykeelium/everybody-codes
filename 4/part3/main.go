package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q04_p3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gears := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gearRaw := scanner.Text()
		if strings.Contains(gearRaw, "|") {
			split := strings.Split(gearRaw, "|")
			gear1, _ := strconv.ParseFloat(split[0], 64)
			gear2, _ := strconv.ParseFloat(split[1], 64)
			gears = append(gears, gear1)
			gears = append(gears, gear2)
		} else {
			gear, _ := strconv.ParseFloat(gearRaw, 64)
			gears = append(gears, gear)
		}
	}

	ratio := float64(1)
	ratio *= (gears[0] / gears[1])

	for i := 2; i < len(gears)-2; i += 2 {
		ratio *= (gears[i] / gears[i+1])
	}

	ratio *= gears[len(gears)-2] / gears[len(gears)-1]

	fmt.Printf("Output: %f\n", 100*ratio)
}
