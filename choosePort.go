package main

import (
	"fmt"
	"os"
)

func ChoosePort() string {
	var port string = "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return ""
	}
	return port
}
