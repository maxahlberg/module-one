package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Insert a number: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("error while scanning: %v \n ", err)
	}
	fmt.Printf("What is this? %T, saying: %s  \n", input, input)
	s, _ := strconv.Atoi(input)
	fmt.Printf("Then I try to convert to: %T, But the value is:  %v \n", s, s)
	fmt.Printf("the error %v \n", err)
	fmt.Printf("the input without parsing is, a %T with value  %v \n", input, input)
}
