package main

import (
	"fmt"
	"os"
    "bufio"
    "strconv"
)

func main(){
	//initialise values
	var numberOfZeros = 0
	var value = 50

	//read the file
	file, err := os.Open("day1-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//fmt.Print(file)

	//loop
	for scanner.Scan(){
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		direction := line[0]

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// 	//check for the operation
		if direction == 'L' {
			value = (value - distance) % 100
			if value < 0 {
				value += 100
			}
		} else {
			value = (value + distance) % 100
		}

		//check if the value is 0
		if value == 0 {
			numberOfZeros++
		}
	}

	fmt.Println(numberOfZeros)
}