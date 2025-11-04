// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )
//
// func main() {
// 	file, err := os.Open("./everybody_codes_e2025_q01_p1.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	scanner := bufio.NewScanner(file)
// 	scanner.Scan()
// 	namesRaw := scanner.Text()
// 	scanner.Scan()
// 	scanner.Scan()
// 	instructionsRaw := scanner.Text()
//
// 	names := strings.Split(namesRaw, ",")
// 	instructions := strings.Split(instructionsRaw, ",")
//
// 	maxName := len(names) - 1
// 	current := int64(0)
//
// 	for _, instruction := range instructions {
// 		direction := string(instruction[0])
// 		actionRaw := string(instruction[len(instruction)-(len(instruction)-1)])
// 		action, _ := strconv.ParseInt(actionRaw, 10, 32)
//
// 		if direction == "L" {
// 			current = int64(math.Max(float64(0), float64(current-action)))
// 			continue
// 		} else {
// 			current = int64(math.Min(float64(maxName), float64(current+action)))
// 		}
// 	}
//
// 	fmt.Printf("Final name: %s", names[current])
// }
