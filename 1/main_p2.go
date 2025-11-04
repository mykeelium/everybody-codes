// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )
//
// func main() {
// 	file, err := os.Open("./everybody_codes_e2025_q01_p2.txt")
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
// 	current := int(0)
//
// 	for _, instruction := range instructions {
// 		direction := string(instruction[0])
// 		actionRaw := string(instruction[1:])
// 		fmt.Println(actionRaw)
// 		action, _ := strconv.ParseInt(actionRaw, 10, 32)
// 		action32 := int(action)
//
// 		if direction == "L" {
// 			current -= action32
// 		} else {
// 			current += action32
// 		}
// 	}
//
// 	fmt.Printf("current: %v \n", current)
// 	current = current % len(names)
//
// 	fmt.Printf("Final name: %s", names[current])
// }
