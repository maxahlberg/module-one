package main

import (
	"bufio"
        "fmt"
	"strings"
        "os"
)

func main() {

	fmt.Print("Please enter a string:\n")
	scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan() // use `for scanner.Scan()` to keep reading
        line := scanner.Text()

        xl := strings.ToLower(line)
	bPref := strings.HasPrefix(xl, "i")
	bSuf := strings.HasSuffix(xl, "n")
	bContain := strings.Contains(xl, "a")
	if bPref && bSuf && bContain {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}

}

