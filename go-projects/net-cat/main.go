package main

import (
	"fmt"
	"netcat/chat"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if len(args) == 1 {
		p, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		if p < 1024 || p > 65535 {
			fmt.Println("Only ports 1024-65535 allowed")
			return
		}

		chat.Port = args[0]

	}

	chat.Server()

}
