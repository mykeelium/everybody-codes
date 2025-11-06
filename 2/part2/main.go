package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cycle(input1 int, input2 int, value1 int, value2 int) (int, int) {
	temp1 := (input1 * input1) - (input2 * input2)
	temp2 := (input1 * input2) + (input1 * input2)
	input1 = temp1
	input2 = temp2
	input1 /= 100000
	input2 /= 100000
	input1 += value1
	input2 += value2
	return input1, input2
}

func iter(start int, end int, size int, iteration int) int {
	return ((end - start) / size) * iteration
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q02_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	inputCenter := strings.Split(inputRaw, "[")[1]
	inputCenter = strings.Split(inputCenter, "]")[0]
	values := strings.Split(inputCenter, ",")
	value1Raw := values[0]
	value2Raw := values[1]
	value164, _ := strconv.ParseInt(value1Raw, 10, 32)
	value264, _ := strconv.ParseInt(value2Raw, 10, 32)
	value1 := int(value164)
	value2 := int(value264)
	value1End := value1 + 1000
	value2End := value2 + 1000

	engraved := 0
	iterations := 0
	xIter := 0
	yIter := 0

	for x := value1; x <= value1End; x += iter(value1, value1End, 101, xIter) {
		xIter++
		yIter = 0
		for y := value2; y <= value2End; y += iter(value2, value2End, 101, yIter) {
			yIter++
			iterations += 1
			toEngrave := true
			test_x := x
			test_y := y
			for i := 0; i < 100; i++ {
				test_x, test_y = cycle(test_x, test_y, x, y)
				if test_x > 1000000 || test_x < -1000000 || test_y > 1000000 || test_y < -1000000 {
					toEngrave = false
					break
				}
			}
			if toEngrave {
				engraved += 1
			}
		}
	}

	fmt.Printf("iterations: %d\n", iterations)
	fmt.Printf("output: %d\n", engraved)
}
