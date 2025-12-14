package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	var range
	var difference

	file, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
}

func isInvalidID(){
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		for ranges := range strings.SpiltSeq(text, ","){
			if ranges == "" {
				continue
			}
		}
	}
}