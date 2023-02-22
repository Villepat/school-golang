package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	var newString []string

	argString := string(args[0])
	template := string(args[1])
	outFile := string(args[2])
	if len(outFile) > 8 {
		if outFile[0:9] == "--output=" {
			outFile = outFile[9:]
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Print("EX: go run . something standard --output=<fileName.txt>")
	}

	newLineAtEnd := false

	if argString[len(argString)-2:] == "\\n" {
		newLineAtEnd = true
	}

	rows := strings.Split(string(args[0]), "\\n")

	for i, _ := range rows {
		if i > 1 {
			fmt.Printf("%s", "\n")
		}

		r1 := rows[i]
		var startLines []int

		for _, c := range r1 {
			startLines = append(startLines, (((int(c) - 32) * 9) + 1))
		}

		rawBytes, err := ioutil.ReadFile(template + ".txt")
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
				newString = append(newString, lines[startLines[j]+count])
				newString = append(newString, "\n")
				count++
				j = 0
				if count == 8 && newLineAtEnd {
					newString = append(newString, "\n")
				}
			}
			if j < len(startLines) {
				newString = append(newString, lines[startLines[j]+count])

			}
		}
	}
	newString = append(newString, "\n")
	output := strings.Join(newString, "")

	err := ioutil.WriteFile(outFile, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
