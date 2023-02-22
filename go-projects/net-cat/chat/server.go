package chat

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	Port = "8989"
)

func Server() {
	listen, err := net.Listen("tcp", ""+":"+Port)
	if err != nil {
		log.Fatal(err)
	} else if err == nil {
		fmt.Println("Now listening to port: " + Port)
	}

	var mutex sync.Mutex

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if NbOfConnections == 10 {
			conn.Write([]byte("Chat is full. Get in line you donkey"))
			conn.Close()
		}
		NbOfConnections++
		go ClientHandler(conn, &mutex)
	}

}
