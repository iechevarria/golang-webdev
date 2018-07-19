package main

import (
    "io"
    "net"
    "log"
)

func main () {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
    }
    defer li.Close()

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Fatalln(err)
            continue
        }

        io.WriteString(conn, "I see you connected")
        conn.Close()
    }
}