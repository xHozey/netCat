package netCat

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	Connections   = make(map[string]net.Conn)
	Mu            sync.Mutex
	LogData       string
	alreadyExists bool
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

		name = cleanStr(name)
		if name == "" {
			conn.Write([]byte("Name cannot be empty. Please enter a valid name.\n[ENTER YOUR NAME]: "))
			continue
		}
		for cnx := range Connections {
			if name == cnx {
				alreadyExists = true
				break
			}
		}
		if alreadyExists {
			conn.Write([]byte("Name is already exist\n[ENTER YOUR NAME]: "))
			alreadyExists = false
			continue
		}
		break
	}
	Mu.Lock()
	Connections[name] = conn
	conn.Write([]byte(LogData))
	Mu.Unlock()
	broadcastMessage("\n"+name+" has joined our chat...\n", nil)
	fmt.Print(name + " has joined our chat...\n")
	return name
}

func broadcastMessage(message string, sender net.Conn) {
	Mu.Lock()
	defer Mu.Unlock()

	for name, conn := range Connections {
		currentTime := time.Now()
		formattedTime := currentTime.Format("2006-01-02 15:04:05")
		if conn != sender {
			user := "[" + formattedTime + "]" + "[" + name + "]:"
			conn.Write([]byte(message))
			conn.Write([]byte(user))
		} else {
			conn.Write([]byte("[" + formattedTime + "]" + "[" + name + "]:"))
		}
	}
	LogData += message
	LogData = strings.TrimSpace(LogData)
	err := os.WriteFile("Server/logs.txt", []byte(LogData), 0o644)
	if err != nil {
		fmt.Print(err)
	}
}
