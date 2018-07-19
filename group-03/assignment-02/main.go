package main

import (
    "io"
    "net"
		"log"
		"fmt"
		"bufio"
)

func main () {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatalln(err)
            continue
        }

				s := bufio.NewScanner(c)
				for s.Scan() {}
				fmt.Println("Code got here.")
				io.WriteString(c, "I see you connected.")
				
				c.Close()
    }
}