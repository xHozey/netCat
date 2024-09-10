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
	result := ""
	for _, val := range s {
		if val >= 0 && val < 32 {
			continue
		}
		result += string(val)
	}

	return result
}
