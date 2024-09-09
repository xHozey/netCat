package netCat

import (
	"fmt"
	"os"
	"strings"
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

func cleanStr(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\x1b[A", "")
	s = strings.ReplaceAll(s, "\x1b[B", "")
	s = strings.ReplaceAll(s, "\x1b[C", "")
	s = strings.ReplaceAll(s, "\x1b[D", "")
	return s
}
