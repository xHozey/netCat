package netCat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func HostServer() {
	port := GetAdress()
	ConnectionsMax := 0
	listner, err := net.Listen("tcp", "localhost"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port " + port)

	for {
		conn, err := listner.Accept()
		if ConnectionsMax < 10 {
			ConnectionsMax++
			if err != nil {
				fmt.Println(err)
				continue
			}
			go func() {
				LinuxMsg(conn)
				name := nameClient(conn)
				handleConn(conn, name)
			}()

		} else {
			conn.Write([]byte("chat group is full!"))
			conn.Close()
		}
	}
}

func handleConn(conn net.Conn, name string) {
	defer func() {
		Mu.Lock()
		delete(Connections, name)
		Mu.Unlock()
		conn.Close()
	}()

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				broadcastMessage(name + " has left our chat...\n")
				fmt.Print(name + " has left our chat...\n")
			} else {
				fmt.Println(err)
			}
			break
		}

		currentTime := time.Now()
		formattedTime := currentTime.Format("2006-01-02 15:04:05")
		user := "[" + formattedTime + "]" + "[" + name + "]:"
		if strings.TrimSpace(data) != "" {
			go broadcastMessage(user + data)
		} else {
			conn.Write([]byte(user + ""))
		}

	}
}
