package main

import (
	"bufio"
	"net"
	"time"
)

func handlerConnecter(connection net.Conn) {
	defer connection.Close()
	connection.Write([]byte(ReadWelcome() + "\n"))
	var name = ChooseName(connection)
	if len(name) == 0 {
		connection.Close()
		return
	} else {
		user++
		if user > 10 {
			user--
			connection.Write([]byte("to many users"))
			connection.Close()
			return
		}
		texte := AfficheHistory()
		connection.Write([]byte(texte))
		mutex.Lock()
		NameT = append(NameT, name)
		Connecte = append(Connecte, connection)
		mutex.Unlock()
		for i, v := range Connecte {
			if v != connection {
				v.Write([]byte("\033[2K\r"))
				v.Write([]byte(name + " has joined our chat...\n"))
				timer4 := time.Now().Format("2006-01-02 15:04:05")
				v.Write([]byte("[" + timer4 + "]" + "[" + NameT[i] + "]:"))
			}
		}
		for {
			var compteur int
			datas := bufio.NewReader(connection)
			timer1 := time.Now().Format("2006-01-02 15:04:05")
			if compteur == 0 {
				connection.Write([]byte("\033[2K\r"))
				connection.Write([]byte("[" + timer1 + "]" + "[" + name + "]:"))
				compteur += 1
			}
			data, err := datas.ReadString('\n')
			if err != nil {
				var tmp int
				mutex.Lock()
				for i, v := range Connecte {
					if v != connection {
						v.Write([]byte("\033[2K\r"))
						v.Write([]byte(name + " has left our chat...\n"))
						timer4 := time.Now().Format("2006-01-02 15:04:05")
						v.Write([]byte("[" + timer4 + "]" + "[" + NameT[i] + "]:"))
					} else {
						tmp = i
					}
				}
				if tmp+1 != len(NameT) || tmp+1 != len(Connecte) {
					Connecte = append(Connecte[:tmp], Connecte[tmp+1:]...)
					NameT = append(NameT[:tmp], NameT[tmp+1:]...)
				} else {
					Connecte = Connecte[:len(Connecte)-1]
					NameT = NameT[:len(NameT)-1]
				}
				mutex.Unlock()
				user -= 1
				break
			}
			n := len(data)
			connection.Write([]byte("\033[1A\033[K"))
			timer2 := time.Now().Format("2006-01-02 15:04:05")
			connection.Write([]byte("[" + timer2 + "]" + "[" + name + "]:" + string(data[:n])))
			SendHistory([]byte("[" + timer2 + "]" + "[" + name + "]:" + string(data[:n])))

			for i, v := range Connecte {
				if v != connection {
					timer3 := time.Now().Format("2006-01-02 15:04:05")
					v.Write([]byte("\033[2K\r"))
					v.Write([]byte("[" + timer3 + "]" + "[" + name + "]:" + string(data[:n])))
					timer4 := time.Now().Format("2006-01-02 15:04:05")
					v.Write([]byte("[" + timer4 + "]" + "[" + NameT[i] + "]:"))
					compteur += 1
				}
			}
			
		}
	}

}
