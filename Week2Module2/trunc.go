package main

import "fmt"

func main() {
	var input float64
	var input2 float64
	var output int
	var output2 int

	println("Input your first floatingpoint number: ")
	_, err :=
		fmt.Scan(&input)
	if err != nil {
		println("Scanning gone wrong", err)
	}
	output = int(input)
	println("The output float number is: ", output)

	println("Input another number: ")
	_, err =
		fmt.Scan(&input2)
	if err != nil {
		println("Second scan gone wrong", err)
	}
	output2 = int(input2)
	println("The second int is:", output2)

}
