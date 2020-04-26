package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func removeNewLine(inputString string)(outputString string){

	inputString = strings.TrimSuffix(inputString, "\n")
	inputString = strings.TrimSuffix(inputString, "\r")
	inputString = strings.TrimSuffix(inputString, "\r\n")
	outputString = inputString
	return outputString
}

func main() {

	fmt.Printf("Input your name please:  \n")
	nameInput := bufio.NewReader(os.Stdin)
	name, err := nameInput.ReadString('\n')
	if err != nil {
		fmt.Printf("Reading stwing went wrong,  %v \n", err)
	}
	fmt.Printf("Now enter your adress: \n")
	addressInput := bufio.NewReader(os.Stdin)
	address, err := addressInput.ReadString('\n')
	if err != nil {
		fmt.Printf("Reading adress went wrong, %v \n", err)
	}
	name = removeNewLine(name)
	address = removeNewLine(address)
	// Create a map
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	// Create a json object
	jsonMap, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error marshaling to json, %v \n", err)
	}

	// print the json object
	fmt.Printf("the json bite array %v \n", string(jsonMap))

	// Unmarshaling the json object (just for fun)
	mUnmarshal := make(map[string]string)
	err = json.Unmarshal(jsonMap, &mUnmarshal)
	if err != nil {
		fmt.Printf("Unmarshaling went wrong %v \n", err)
	}
	//fmt.Printf("The unmarshal golang map: %v \n", mUnmarshal)
}
