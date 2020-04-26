package main

import "fmt"
import "os"
import "log"
import "bufio"
import "strings"
import "unicode/utf8"

type Person struct {
	fname string
	lname string
}

func main() {
	var filename string
	var names []Person
	fmt.Print("Enter a file name : ")
	fmt.Scan(&filename)

	f, e := os.Open(filename)
	if e != nil {
		log.Fatalf("Failed to open file: %s", e)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		name := scanner.Text()
		nametokens := strings.Fields(name)
		pers := Person{fname: nametokens[0], lname: nametokens[1]}
		if utf8.RuneCountInString(pers.fname) > 20 {
			pers.fname = pers.fname[:20]
		}
		if utf8.RuneCountInString(pers.lname) > 20 {
			pers.lname = pers.lname[:20]
		}
		names = append(names, pers)
	}

	for _, val := range names {
		fmt.Printf("First Name: %s Last Name: %s\n", val.fname, val.lname)
	}
}
