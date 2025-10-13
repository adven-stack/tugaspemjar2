package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server TCP berjalan di port 8080...")

    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    go func() {
        for {
            message, _ := bufio.NewReader(conn).ReadString('\n')
            fmt.Print("Client: " + message)
        }
    }()

    for {
        fmt.Print("Server: ")
        text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        conn.Write([]byte(text))
    }
}
