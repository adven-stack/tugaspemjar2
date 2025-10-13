package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    go func() {
        for {
            message, _ := bufio.NewReader(conn).ReadString('\n')
            fmt.Print("Server: " + message)
        }
    }()

    for {
        fmt.Print("Client: ")
        text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        conn.Write([]byte(text))
    }
}
