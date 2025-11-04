package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q01_p3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	namesRaw := scanner.Text()
	scanner.Scan()
	scanner.Scan()
	instructionsRaw := scanner.Text()

	names := strings.Split(namesRaw, ",")
	instructions := strings.Split(instructionsRaw, ",")

	for _, instruction := range instructions {
		direction := string(instruction[0])
		actionRaw := string(instruction[1:])
		action, _ := strconv.ParseInt(actionRaw, 10, 32)
		action32 := int(action)
		index := 0

		if direction == "L" {
			tempIndex := action32 % len(names)
			if tempIndex != 0 {
				index = int(float64(len(names)) - math.Abs(float64(tempIndex)))
			}
		} else {
			tempIndex := action32 % len(names)
			index += tempIndex
		}

		temp := names[0]
		names[0] = names[index]
		names[index] = temp
	}

	fmt.Printf("Final name: %s", names[0])
}
