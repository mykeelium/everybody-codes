package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q04_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gears := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gearRaw := scanner.Text()
		gear, _ := strconv.ParseFloat(gearRaw, 64)
		gears = append(gears, gear)
	}

	ratio := float64(1)

	for i := 0; i < len(gears)-1; i++ {
		ratio *= (gears[i] / gears[i+1])
	}

	total := float64(10000000000000)

	fmt.Printf("Output: %f\n", total/ratio)
}
