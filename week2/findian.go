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
	input = strings.TrimSuffix(input, "\n")
	if err != nil {
		fmt.Printf("Error while scanning: %v \n", err)
	}
	inputNoSpace := strings.Replace(input, " ", "", -1)
	inputLower := strings.ToLower(inputNoSpace)
	anyI := strings.HasPrefix(inputLower, "i")
	anyA := strings.IndexAny(inputLower, "a")
	anyN := strings.HasSuffix(inputLower, "n")
	fmt.Printf("%s   %s   %s",anyI, anyA, anyN)
	if !anyI || anyA == -1 || !anyN {
		fmt.Printf("Not Found\n")
	} else {
		fmt.Printf("Found!\n")
	}
	avg := 2% (2 +4 )
	println(avg)
	avg2 := float64(4+2)/2
	println(avg2)
	avg3 := float64(4+2)/2.0
	println(avg3)
}
