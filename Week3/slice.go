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
	fmt.Printf("What is this? %T \n", input)
	if err != nil {
		fmt.Printf("error while scanning: %v \n ", err)
	}
	if s, err := strconv.Atoi(input); err != nil {
		fmt.Printf("%T, %v", s, s)
	}
	fmt.Printf("the error %v \n", err )
	//fmt.Printf("The input is: %v \n ", s)
	fmt.Printf("the input without parsing is %v \n", input)
}
