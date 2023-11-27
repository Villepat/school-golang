package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	_, err := os.Open(os.Args[0])
	if err != nil {
		fmt.Println(err)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else {
		data, _ := ioutil.ReadFile(os.Args[1])
		fmt.Print(string(data))
	}
}
