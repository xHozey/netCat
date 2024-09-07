package netCat

import (
	"fmt"
	"net"
)

func LinuxMsg(conn net.Conn) {
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
