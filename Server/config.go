package netCat

import (
	"fmt"
	"os"
)

func GetAdress() string {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	port := ":8989"
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}
	return port
}
