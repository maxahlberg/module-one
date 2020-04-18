package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Write a program which prompts the user to enter a floating point number and
prints the integer which is a truncated version of the floating point number that
was entered. Truncation is the process of removing the digits to the right of the
decimal place. */

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a floating point number:\n")
	strInput, _ := reader.ReadString('\n')
	strInput = strings.TrimSuffix(strInput, "\n")

	if floatNum, err := strconv.ParseFloat(strInput, 64); err == nil {
		fmt.Println(int(floatNum))
	}

	return
}
