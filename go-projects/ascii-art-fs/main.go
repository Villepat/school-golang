package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {

	if len(os.Args[1:]) == 2 {

		if string(os.Args[2]) != "standard" && string(os.Args[2]) != "thinkertoy" && string(os.Args[2]) != "shadow" || string(os.Args[1]) == "" {

			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("")

			fmt.Println("EX: go run . something standard")
			return

		}

		if !isASCII(string(os.Args[1])) {
			fmt.Println("Only ASCII characters are allowed")
			return
		}

		args := os.Args[1:]
		template := string(args[1])

		rows := strings.Split(string(args[0]), "\\n")

		for i, _ := range rows {
			if i > 1 {
				fmt.Printf("%s", "\n")
			}

			r1 := rows[i]
			var startLines []int

			for _, c := range r1 {
				startLines = append(startLines, ((int(c) - 32) * 9))
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
					fmt.Printf("%s", lines[startLines[j]+count])
					fmt.Printf("%s", "\n")
					count++
					j = 0
					if count == 8 {
						fmt.Printf("%s", "\n")
					}
				}
				if j < len(startLines) {
					fmt.Printf("%s", lines[startLines[j]+count])
				}
			}
		}

	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("")
		fmt.Print("EX: go run . something standard")
		return

	}
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// credit to Proferssor Stack Overflow for the ASCII check function <3
