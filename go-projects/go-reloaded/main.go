package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	
)

func main() {
	_, err := os.Open(os.Args[0])
	if err != nil {
		fmt.Println(err)
	} else if len(os.Args) >= 4 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else { // no errors and a right amount of arguments and we open the file
		data, _ := ioutil.ReadFile(os.Args[1])
		// fmt.Print(string(data))
		txt := string(data)
		txtslice := strings.Split(txt, " ")
		printout := ""
		quotechecker := 0
		

		for index, value := range txtslice { //start ranging over slice to check elements
			// ========= UP =========
			if value == "(up)" { //if an element value equals to a trigger
				txtslice[index-1] = strings.ToUpper(string(txtslice[index-1])) //we change the value at preceding index using a builtin function
				txtslice[index] = ""
			} else if value == "(up," {
				re := regexp.MustCompile("[0-9]+")
				factor := re.FindAllString(string(txtslice[index+1]), -1)
				stringfactor := strings.Join(factor, "")   //convert slice of a string into string
				intfactor, _ := strconv.Atoi(stringfactor) //convert string into int that we can work with
				for i := 1; i <= intfactor; i++ {
					txtslice[index-i] = strings.ToUpper(txtslice[index-i])
				}
				txtslice[index+1] = ""
				txtslice[index] = ""
			}
			// ========= LOW =========
			if value == "(low)" { //if an element value equals to a trigger
				txtslice[index-1] = strings.ToLower(string(txtslice[index-1])) //we change the value at preceding index using a builtin function
				txtslice[index] = ""
			} else if value == "(low," {
				re := regexp.MustCompile("[0-9]+")
				factor := re.FindAllString(string(txtslice[index+1]), -1)
				stringfactor := strings.Join(factor, "")   //convert slice of a string into string
				intfactor, _ := strconv.Atoi(stringfactor) //convert string into int that we can work with
				for i := 1; i <= intfactor; i++ {
					txtslice[index-i] = strings.ToLower(txtslice[index-i])
				}
				txtslice[index+1] = ""
				txtslice[index] = ""
			}
			// ========= CAP =========
			if value == "(cap)" { //if an element value equals to a trigger
				txtslice[index-1] = strings.Title(strings.ToLower(string(txtslice[index-1]))) //we change the value at preceding index using a builtin function
				txtslice[index] = ""
			} else if value == "(cap," {
				re := regexp.MustCompile("[0-9]+")
				factor := re.FindAllString(string(txtslice[index+1]), -1)
				stringfactor := strings.Join(factor, "")   //convert slice of a string into string
				intfactor, _ := strconv.Atoi(stringfactor) //convert string into int that we can work with
				for i := 1; i <= intfactor; i++ {
					txtslice[index-i] = strings.Title(strings.ToLower(txtslice[index-i]))
				}
				txtslice[index+1] = ""
				txtslice[index] = ""
			}
			// ========= HEX =========
			if value == "(hex)" { //if an element value equals to a trigger
				decimal, _ := strconv.ParseInt(txtslice[index-1], 16, 32)
				decimalstring := strconv.FormatInt(decimal, 10)
				txtslice[index-1] = decimalstring //we change the value at preceding index using a builtin function
				txtslice[index] = ""
			}
			// ========= BIN =========
			if value == "(bin)" { //if an element value equals to a trigger
				decimal, _ := strconv.ParseInt(txtslice[index-1], 2, 32)
				decimalstring := strconv.FormatInt(decimal, 10)
				txtslice[index-1] = decimalstring //we change the value at preceding index using a builtin function
				txtslice[index] = ""
			}

			// ========= A AN =========
			if value == "A" || value == "a" && index < len(txtslice) {
				isnextvowel, err := regexp.MatchString(`\A[aeiou]|\A[AEIOU]|\Ah`, txtslice[index+1])
				if err != nil {
					fmt.Println(err)
				}
				if isnextvowel {
				
					if txtslice[index] == "a" {
						txtslice[index] = "an"
					} 
					if txtslice[index] == "A" {
						txtslice[index] = "An"
					}

				}
			}

			// ========= 'PUNCTUATION' =========
			
			if txtslice[index] == "'" && index < len(txtslice)-1 && quotechecker == 0 {
				txtslice[index+1] = "'" + txtslice[index+1]
				txtslice[index] = ""
				quotechecker = 1
				
			}
			if txtslice[index] == "'" && quotechecker == 1   {
			
				txtslice[index-1] = txtslice[index-1] +  "'"
				txtslice[index] = ""
				quotechecker = 0
			}
			if txtslice[index] == "." {
				txtslice[index-1] = txtslice[index-1] + txtslice[index]
				txtslice[index] = ""
			}

			// ========= PUNCTUATION =========
			s := strings.Join(txtslice, " ")
			s2 := strings.Join(strings.Fields(string(s)), " ")

			x := []byte(s2)
			for i := 1; i <= len(x); i++ {
				for i, c := range []byte(x) {
					if c == '.' || c == ',' || c == '!' || c == '?' || c == ';' || c == ':' {
						if x[i-1] == ' ' { //if a punctuaction character has a whitespace before it
							x[i-1] = x[i] //we swap their places
							x[i] = ' '
						}
					}
					if c == '\'' && i == len(x)-1 && x[i-1] == ' '  {
						
						x[i-1] = x[i]
						x[i] = ' '
					}
				}
				//fmt.Println(string(x))
				
			}
			printout = strings.Join(strings.Fields(string(x)), " ") 
			err = ioutil.WriteFile(os.Args[2], []byte(printout), 0666)
			if err != nil {
				fmt.Println(err)

		}
		}	
	}

} 
