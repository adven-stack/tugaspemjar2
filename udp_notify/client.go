package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8081")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Ketik pesan untuk dikirim ke server UDP:")

	for {
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error mengirim pesan:", err)
			return
		}
	}
}
