package main

import (
	"fmt"
	"os"
	"time"
)

var Hhistory string


func Reset() {
	var timer = time.Now().Format("2006-01-02 15:04:05")
	r, err := os.Create("history" + timer + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	Hhistory = r.Name()
}

func ReadWelcome() string {
	var texte, err = os.ReadFile("welcome.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture de Welcome.txt")
	}
	textes := string(texte)
	return textes
}

func AfficheHistory() string {
	file, err := os.ReadFile(Hhistory)
	if err != nil {
		files, err := os.Create(Hhistory)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		file, err := os.ReadFile(files.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		texte := string(file)
		return texte
	}
	texte := string(file)
	return texte
}

func SendHistory(data []byte) {
	file, err := os.OpenFile(Hhistory, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		files, err := os.Create(Hhistory)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		files.WriteString(string(data))
		return
	}
	file.WriteString(string(data))
}
