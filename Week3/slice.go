package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	i := 0
	// Creating a slice of size 3
	var sortMe = make([]int, 3)
	// Starting the loop to insert Integers
	for i < 1 {
		// Asking for input
		fmt.Println("Insert a number: ")
		// Reading it
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if err != nil {
			fmt.Printf("error while scanning: %v \n ", err)
		}
		// Check if the input was x and end loop
		if input == "X" || input == "x" {
			i = i + 1
			continue
		}
		// Convert the string input to Integers
		s, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Cant convert to int: \n  %v \n", err)
		}
		// Adding the new inputed number to the slice
		sortMe = append(sortMe, s)
		// Sort the input slice "sortMe"
		lowestThisK := 0
		lowestIteK := 0
		var sortMeFinal = make([]int, len(sortMe))
		var sortMeLeftOver = make([]int, 0)
		for k := 0; len(sortMe) > 0; k++ {
			lowest := math.Inf(+1)
			for ite, val := range sortMe {
				if float64(val) < lowest {
					lowestThisK = val
					lowestIteK = ite
					lowest = float64(val)
				}
			}
			// creating new shorted and left over slice
			sortMeFinal[k] = lowestThisK
			for iteration, value := range sortMe {
				if iteration != lowestIteK {
					sortMeLeftOver = append(sortMeLeftOver, value)
				}
			}
			sortMe = sortMeLeftOver
			sortMeLeftOver = make([]int, 0)
		}
		fmt.Printf("The sorted slice is:  %v \n", sortMeFinal)
		//resetting sortMe
		sortMe = sortMeFinal
	}

}
