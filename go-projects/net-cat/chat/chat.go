package chat

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (

	// handle amount of connections
	NbOfConnections int
	//track names being used
	Clients = make(map[string]net.Conn)
)

type MessageStruct struct {
	Username string
	Time     string
	Message  string
}

var History []string

// Handle connection names, input and disconnections
func ClientHandler(conn net.Conn, mutex *sync.Mutex) {

	name := GetName(conn, mutex)

	ServerBroadcast(name)
	for {
		if !WriteMsg(conn, mutex, name) { // stops reading buffer if client disconnects
			break
		}
	}

}

// Retrieves a name from a connecting client, stores it, serves welcome message and broadcasts chat history to client
func GetName(conn net.Conn, mutex *sync.Mutex) string {
	mutex.Lock()
	b, err := os.ReadFile("penguin.txt")
	if err != nil {
		fmt.Println("I'm a funcing penguin")
	}
	mutex.Unlock()
	conn.Write(b)

ReName:
	buffer := make([]byte, 1024)
	bufLen, err := conn.Read(buffer)
	if err != nil {
		NbOfConnections--
		conn.Close()
		fmt.Println(err)
		return "" // Breaks renaming jump statement when client disconnects
	}
	msgBuf := string(buffer)
	name := string(msgBuf[:bufLen-1])
	if name == "" {
		conn.Write([]byte("We unfortunately don't allow empty names. Please give a valid name\n[ENTER YOUR NAME]:"))
		goto ReName
	}
	mutex.Lock()
	for val := range Clients {
		if name == val {
			conn.Write([]byte("This name is already in use, please try another one:"))
			mutex.Unlock()
			goto ReName
		}
	}
	Clients[name] = conn
	PrintHistory(History, conn)
	mutex.Unlock()
	return name
}

// Writes previous chat history to new clients connecting
func PrintHistory(history []string, conn net.Conn) {
	for _, v := range history {
		conn.Write([]byte(v))
	}
}

// Uses a buffer to await client input message
func WriteMsg(conn net.Conn, mutex *sync.Mutex, name string) bool {
	if name == "" {
		return false
	}
	buffer := make([]byte, 1024)
	bufLen, err := conn.Read(buffer)
	if err != nil {
		mutex.Lock()
		fmt.Println(err)
		conn.Close()
		delete(Clients, name)
		for _, c := range Clients {
			c.Write([]byte(name + " has left the chat" + "\n"))
		}

		NbOfConnections--
		mutex.Unlock()
		return false
	}
	msgBuf := string(buffer)
	msgtxt := string(msgBuf[:bufLen-1])
	Clients[name] = conn

	msg := MessageStruct{
		Message:  msgtxt,
		Username: name,
		Time:     time.Now().Format("01-02-2006 15:04:05"),
	}
	msgbroadcast := "[" + msg.Time + "][" + msg.Username + "]:" + msg.Message + "\n"
	for key, c := range Clients {
		if key == name {
			conn.Write([]byte("\033[A"))
		}
		c.Write([]byte(msgbroadcast))
	}
	History = append(History, msgbroadcast)
	return true
}

// Broadcasts new client to current client list
func ServerBroadcast(s string) {
	if s != "" {
		for _, c := range Clients {
			c.Write([]byte(s + " joined the chat" + "\n"))
		}
	}
}
