package main

import (
	"bufio"
	"fmt"

	//"io/ioutil"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4321")
	if err != nil {
		fmt.Println("ошбка 1", err)
	}
	fmt.Println("введите имя")
	var m string
	fmt.Scan(&m)
	conn.Write([]byte(m))

	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadBytes('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		conn.Write(message)
		go chat(conn)
	}

}

func chat(conn net.Conn) {
	for {
		var bufer []byte = make([]byte, 1024)
		n, err := conn.Read(bufer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print(string(bufer[:n]))

	}

}
