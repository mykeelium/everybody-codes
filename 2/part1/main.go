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
	input1 /= 10
	input2 /= 10
	input1 += value1
	input2 += value2
	return input1, input2
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q02_p1.txt")
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

	input1 := 0
	input2 := 0
	input1, input2 = cycle(input1, input2, value1, value2)
	input1, input2 = cycle(input1, input2, value1, value2)
	input1, input2 = cycle(input1, input2, value1, value2)

	fmt.Printf("output: [%d,%d]\n", input1, input2)
}
