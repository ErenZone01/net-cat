package main

import (
	"fmt"
	"net"
	"sync"
)

var Connecte []net.Conn
var NameT []string
var mutex sync.Mutex
var user int

func main() {
	var port = ChoosePort()
	addresse := ":" + port
	ecouteur, err := net.Listen("tcp", addresse)
	if err != nil {
		fmt.Println("Erreurs lors de la creation de l'ecouteur", err)
		return
	}
	defer ecouteur.Close()
	fmt.Println("Listening on the port :", port)
	Reset()
	for {
		connexion, err := ecouteur.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion", err)
			continue
		}
		go handlerConnecter(connexion)
	}

}
