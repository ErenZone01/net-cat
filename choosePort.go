package main

import (
	"fmt"
	"os"
	"strings"
)

func ChoosePort() string {
	var port string = "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
		var texte = strings.Trim(port, "0")
		if len(texte) == 0 {
			return ""
		}
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return ""
	}
	return port
}
