package netCat

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	Connections = make(map[string]net.Conn)
	Mu          sync.Mutex
	LogData     string
)

func nameClient(conn net.Conn) string {
	name := ""
	for {
		var err error
		name, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected before providing a name.")
				conn.Close()
				return ""
			}
			fmt.Println(err)
			continue
		}
		name = strings.TrimSpace(name)

		if name == "" {
			conn.Write([]byte("Name cannot be empty. Please enter a valid name.\n"))
			continue
		}
		break
	}
	Mu.Lock()
	Connections[name] = conn
	conn.Write([]byte(LogData))
	Mu.Unlock()
	broadcastMessage(name + " has joined our chat...\n")
	fmt.Print(name + " has joined our chat...\n")
	return name
}

func broadcastMessage(message string) {
	Mu.Lock()
	defer Mu.Unlock()

	for _, conn := range Connections {
		conn.Write([]byte(message))
	}
	LogData += message
	err := os.WriteFile("Server/logs.txt", []byte(LogData), 0o644)
	if err != nil {
		fmt.Print(err)
	}
}
