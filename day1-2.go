package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numberOfZeros := 0
	value := 50

	file, err := os.Open("day1-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// CLICK-BY-CLICK simulation
		for i := 0; i < distance; i++ {
			if direction == 'R' {
				value++
				if value == 100 {
					value = 0
				}
			} else { // L
				value--
				if value < 0 {
					value = 99
				}
			}

			// Count zero ANY time it occurs
			if value == 0 {
				numberOfZeros++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(numberOfZeros)
}
