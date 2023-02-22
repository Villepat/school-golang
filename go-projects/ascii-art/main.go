 package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:] // check if 0 input, check if invalid input

	argstring := string(args[0])

	newLineAtEnd := false

	if argstring[len(argstring)-2:] == "\\n" {
		newLineAtEnd = true
	}

	rows := strings.Split(string(args[0]), "\\n")


	for i, _ := range rows {
		if i > 1 {
			fmt.Printf("%s", "\n")
		}

		r1 := rows[i] // row 1
		var startLines []int

		for _, c := range r1 {
			startLines = append(startLines, (((int(c) - 32) * 9) + 1))
		}

		rawBytes, err := ioutil.ReadFile("standard.txt")
		if err != nil {
			fmt.Println(err)
		}

		lines := strings.Split(string(rawBytes), "\n")

		count := 0

		for j := 0; j < len(startLines); j++ {

			if count == 8 {
				break
			}
			if j == len(startLines)-1 {
				fmt.Printf("%s", lines[startLines[j]+count])
				fmt.Printf("%s", "\n")
				count++
				j = 0
				if count == 8 && newLineAtEnd {
					fmt.Printf("%s", "\n")
				}
			}
			if j < len(startLines) {
				fmt.Printf("%s", lines[startLines[j]+count])
				return
			}
		}
	}
}
