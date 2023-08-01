package main

import (
	"bufio"
	"net"
	"strings"
)

func ChooseName(connection net.Conn) string {
	connection.Write([]byte("[ENTER YOUR NAME]: "))
	var names = bufio.NewReader(connection)
	var name, err = names.ReadString('\n')
	if err != nil {
		return ""
	}
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		connection.Write([]byte("\033[1A\033[K"))
		return ChooseName(connection)
	} else if len(name) > 25 {
		connection.Write([]byte("\033[1A\033[K"))
		connection.Write([]byte("THE SIZE OF THE NAME MUST NOT EXCEED 25 CHARACTERS\n"))
		return ChooseName(connection)
	}
	for _, v := range NameT {
		if v == name {
			connection.Write([]byte("\033[1A\033[K"))
			connection.Write([]byte("[NAME ALREADY USED]\n"))
			return ChooseName(connection)
		}
	}
	return name
}
