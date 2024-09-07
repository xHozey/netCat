package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var ConnectionsMax int

var (
	connections = make(map[string]net.Conn)
	mu          sync.Mutex
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	port := ":8989"
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}

	listner, err := net.Listen("tcp", "localhost"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port " + port)

	go readFromTerminal()

	for {
		conn, err := listner.Accept()
		if ConnectionsMax < 10 {
			ConnectionsMax++
			go linuxMsg(conn)

			if err != nil {
				fmt.Println(err)
				continue
			}
			go func() {
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
		mu.Lock()
		delete(connections, name)
		mu.Unlock()

		conn.Close()
	}()
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				broadcastMessage(name + " has left our chat...\n")
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

func nameClient(conn net.Conn) string {
	name := ""
	for {

		name, _ = bufio.NewReader(conn).ReadString('\n')

		if name == "\n" || strings.Trim(name, " ") == "\n" {
			conn.Write([]byte("Name cannot be empty. Please enter a valid name.\n"))
			continue
		}
		break
	}
	name = name[:len(name)-1]
	mu.Lock()
	connections[name] = conn
	mu.Unlock()
	broadcastMessage(name + " has joined our chat...\n")

	return name
}

func broadcastMessage(message string) {
	mu.Lock()
	defer mu.Unlock()

	for _, conn := range connections {
		conn.Write([]byte(message))
	}
	fmt.Print(message)
}

func linuxMsg(conn net.Conn) {
	message := []string{
		"Welcome to TCP-Chat!\n",
		"         _nnnn_      \n",
		"        dGGGGMMb     \n",
		"       @p~qp~~qMb    \n",
		"       M|@||@) M|    \n",
		"       @,----.JM|    \n",
		"      JS^\\__/  qKL   \n",
		"     dZP        qKRb  \n",
		"    dZP          qKKb \n",
		"   fZP            SMMb\n",
		"   HZM            MMMM\n",
		"   FqM            MMMM\n",
		" __| \".        |\\dS\"qML\n",
		" |    `.       | `' \\Zq\n",
		"_)      \\.___.,|     .'\n",
		"\\____   )MMMMMP|   .'\n",
		"     `-'       `--'\n",
		"[ENTER YOUR NAME]: ",
	}
	for _, line := range message {
		_, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func readFromTerminal() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(text) == "" || strings.Contains(text, "/") || strings.ContainsRune(text, '\\') {
			continue
		}
		broadcastMessage("Server: " + text + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from terminal:", err)
	}
}
