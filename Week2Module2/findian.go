package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("I promise I can see if your input has i, a and n in it. Write it now: \n")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while scanning: %v \n", err)
	}
	inputNoSpace := strings.Replace(input, " ", "", -1)
	inputLower := strings.ToLower(inputNoSpace)
	anyI := strings.IndexAny(inputLower, "i")
	anyA := strings.IndexAny(inputLower, "a")
	anyN := strings.IndexAny(inputLower, "n")
	if anyI == -1 || anyA == -1 || anyN == -1 {
		fmt.Printf("Not Found\n")
	} else {
		fmt.Printf("Found!\n")
	}
}
