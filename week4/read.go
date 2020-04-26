package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Define a name struct
	type name struct {
		fname string
		lname string
	}

	var fileName string
	// Recieving a file
	fmt.Printf("Input name of file: \n")
	fmt.Scan(&fileName)

	// Reading parts of the file
	f, err := os.Open(fileName)
	if err != nil{
		fmt.Printf("Error opening file: %v \n", err)
	}
	rd := bufio.NewReader(f)
	var nameSlice []name
	for {
		var restOfLine []byte
		firstPart, err := rd.ReadString(' ')
		restOfLine, _, err = rd.ReadLine()
		if err != nil{
			break
		}
		lineName := name{
			fname: firstPart,
			lname: string(restOfLine),
		}
		nameSlice = append(nameSlice, lineName)
	}
	// Iterate through the slice and print each struct
	for i, v := range nameSlice{
		fmt.Printf("The name on the %v row was %s \n", i+1, v)
	}

	// close the file when finished
	f.Close()
}
