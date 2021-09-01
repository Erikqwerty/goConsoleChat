package main

import (
	"fmt"
	"net"
	"os"
)

var mapconn map[string]net.Conn = make(map[string]net.Conn)

func main() {

	listener, err := net.Listen("tcp", ":4321")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		var buf []byte = make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		mapconn[string(buf[:n])] = conn
		fmt.Printf("%#v\n", mapconn)
		go zapros(conn)
	}

}

func readbufer(conn net.Conn) []byte {
	buffer := make([]byte, 128)
	var response []byte
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		response = append(response, buffer[:n]...)

		if n < len(buffer) {
			break
		}
	}

	fmt.Print("Принятые данные:", string(response))
	return response
}

func zapros(conn net.Conn) {
	for {
		bufer := string(readbufer(conn))
		for _, con := range mapconn {
			if con != conn {
				con.Write([]byte(bufer))
			}
		}
	}
}
