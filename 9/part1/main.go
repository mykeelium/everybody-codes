package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q09_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw1 := scanner.Text()
	scanner.Scan()
	inputRaw2 := scanner.Text()
	scanner.Scan()
	inputRaw3 := scanner.Text()

	scale1 := strings.Split(inputRaw1, ":")[1]
	scale2 := strings.Split(inputRaw2, ":")[1]
	scale3 := strings.Split(inputRaw3, ":")[1]

	var child string
	parents := []string{}
	scale1HasUnique := false
	scale2HasUnique := false
	scale3HasUnique := false
	dna := len(scale1)
	for x := range dna {
		if scale1HasUnique && scale2HasUnique {
			child = scale3
			break
		} else if scale1HasUnique && scale3HasUnique {
			child = scale2
			break
		} else if scale2HasUnique && scale3HasUnique {
			child = scale1
			break
		}

		if !scale3HasUnique && scale1[x] == scale2[x] && scale1[x] != scale3[x] {
			scale3HasUnique = true
			parents = append(parents, scale3)
			continue
		} else if !scale2HasUnique && scale1[x] == scale3[x] && scale1[x] != scale2[x] {
			scale2HasUnique = true
			parents = append(parents, scale2)
			continue
		} else if !scale1HasUnique && scale2[x] == scale3[x] && scale1[x] != scale2[x] {
			scale1HasUnique = true
			parents = append(parents, scale1)
			continue
		}
	}

	diff1 := 0
	diff2 := 0
	diffString1 := ""
	diffString2 := ""

	for x := range dna {
		if parents[0][x] == child[x] {
			diff1++
			diffString1 += "+"
		} else {
			diffString1 += " "
		}
		if parents[1][x] == child[x] {
			diff2++
			diffString2 += "+"
		} else {
			diffString2 += " "
		}
	}

	fmt.Printf("Output: %d\n", diff1*diff2)
	fmt.Printf("scale1HasUnique: %v, scale2HasUnique: %v, scale3HasUnique: %v\n", scale1HasUnique, scale2HasUnique, scale3HasUnique)
	fmt.Printf("diff1: %d, diff2: %d\n\n", diff1, diff2)

	fmt.Printf("prnt1: %s\nchild: %s\ndiffs: %s\n\n", parents[0], child, diffString1)
	fmt.Printf("prnt2: %s\nchild: %s\ndiffs: %s\n\n", parents[1], child, diffString2)
	fmt.Printf("parents: %v\n", parents)
}
