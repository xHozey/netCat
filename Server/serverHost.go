package netCat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var ConnectionsMax = 0

func HostServer() {
	port := GetAdress()
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port " + port)

	for {
		conn, err := listner.Accept()
		if ConnectionsMax < 2 {
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
		delete(ConnectionsName, conn)
		ConnectionsMax--
		Mu.Unlock()
		conn.Close()
	}()

	for {
		currentTime := time.Now()
		formattedTime := currentTime.Format("2006-01-02 15:04:05")
		user := "\n[" + formattedTime + "]" + "[" + name + "]:"

		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				broadcastMessage("\n"+name+" has left our chat...\n", nil)
				conn.Write([]byte(user))
				fmt.Print(name + " has left our chat...")
			} else {
				fmt.Println(err)
			}
			break
		}

		data = cleanStr(data)
		if data != "" {
			go broadcastMessage(user+data+"\n", conn)
		} else {
			conn.Write([]byte("[" + formattedTime + "]" + "[" + name + "]:"))
		}

	}
}
