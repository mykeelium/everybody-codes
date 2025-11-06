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
	file, err := os.Open("./everybody_codes_e2025_q03_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	numsRaw := strings.Split(inputRaw, ",")
	chain := []int64{}

	for range numsRaw {
		for _, valueRaw := range numsRaw {
			box, _ := strconv.ParseInt(valueRaw, 10, 32)
			if len(chain) == 0 || box > chain[0] {
				chain = append([]int64{box}, chain...)
			} else {
				added := false
				for i := 0; i < len(chain)-1; i++ {
					if chain[i] > box && chain[i+1] < box {
						chain = append(chain[0:i], append([]int64{box}, chain[i+1:]...)...)
						added = true
					}
				}
				if !added && box < chain[len(chain)-1] {
					chain = append(chain, box)
				}
			}

		}
	}

	sum := int64(0)
	for _, value := range chain {
		sum += value
	}
	fmt.Printf("Output: %d\n", sum)
	fmt.Printf("Chain Length: %d\n", len(chain))
	fmt.Printf("Chain: %v\n", chain)
}
