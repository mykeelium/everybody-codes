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
	file, err := os.Open("./everybody_codes_e2025_q03_p3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	numsRaw := strings.Split(inputRaw, ",")
	nums := []int64{}
	chains := [][]int64{}

	for _, valueRaw := range numsRaw {
		box, _ := strconv.ParseInt(valueRaw, 10, 32)
		nums = append(nums, box)
	}

	for len(nums) > 0 {
		chain := []int64{}
		for range nums {
			for _, num := range nums {
				if len(chain) == 0 || num > chain[0] {
					chain = append([]int64{num}, chain...)
				} else {
					added := false
					for i := 0; i < len(chain)-1; i++ {
						if chain[i] > num && chain[i+1] < num {
							chain = append(chain[0:i], append([]int64{num}, chain[i+1:]...)...)
							added = true
						}
					}
					if !added && num < chain[len(chain)-1] {
						chain = append(chain, num)
					}
				}
			}
		}

		fmt.Printf("Finished chain: %d\n", len(chains))

		for _, item := range chain {
			for i, original := range nums {
				if item == original {
					nums[i] = nums[len(nums)-1]
					nums = nums[:len(nums)-1]
					break
				}
			}
		}

		chains = append(chains, chain)
		fmt.Printf("Remaining nums: %d\n", len(nums))
	}

	fmt.Printf("Chains: %d\n", len(chains))
}
